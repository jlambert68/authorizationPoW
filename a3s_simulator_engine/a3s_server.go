package a3s_simulator_engine

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
	"jlambert/authorizationPoW/grpc_api/a3s_grpc_api"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/****************************************************/
// A3S Server object hodling "some" information
type A3SServerObjectStruct struct {
	logger *logrus.Logger
}

var a3SServerObject *A3SServerObjectStruct

/****************************************************/
//  A3S gRPC Server objects
var (
	a3sGrpcServer *grpc.Server
	lis           net.Listener
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type A3S_GrpcServerStruct struct{}

/****************************************************/
// Standard gRPC Client connect towards other gRPC server
//--- Not needed for A3S Server ---
/*
var (
	remoteXxxGrpcServerConnection *grpc.ClientConn
	XxxGrpcServerClient           qml_server_grpc_api.QmlGrpcServicesClient

	XxxGrpcServer_address_to_dial string = common_config.XxxxServer_address + common_config.XxxxServer_port
)
*/

/****************************************************/
// Database cache object
var databaseMemoryCache *cache.Cache

/****************************************************/
// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		a3SServerObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend gRPC Server
		a3SServerObject.StopGrpcServer()

	}
}

/****************************************************/
func BackendServerMain() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// Saved objects having no expiration time will never be deleted
	databaseMemoryCache = cache.New(5*time.Minute, 10*time.Minute)

	// Set up A3S-Object
	a3SServerObject = &A3SServerObjectStruct{}

	// Init logger
	a3SServerObject.InitLogger("")

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start A3S gRPC-server
	a3SServerObject.InitGrpcServer()

	// Exits when user press "ctrl-c" in terminal
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	// Write message in terminal to show that process is alive
	for {
		fmt.Println("sleeping...for another 5 minutes")
		time.Sleep(300 * time.Second) // or runtime.Gosched() or similar per @misterbee
	}
}

// Set up and start Backend gRPC-server
func (a3SServerObject *A3SServerObjectStruct) InitGrpcServer() {

	var err error

	// Find first non allocated port from defined start port
	a3SServerObject.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("A3S gRPC Server tries to start")

	a3SServerObject.logger.WithFields(logrus.Fields{
		"Id":                             "ca3593b1-466b-4536-be91-5e038de178f4",
		"common_config.A3SServer_port: ": common_config.A3SServer_port,
	}).Debug("Start listening on:")
	lis, err = net.Listen("tcp", ":"+common_config.A3SServer_port)

	if err != nil {
		a3SServerObject.logger.WithFields(logrus.Fields{
			"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		a3SServerObject.logger.WithFields(logrus.Fields{
			"Id":                             "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
			"common_config.A3SServer_port: ": common_config.A3SServer_port,
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterWorkerServer gRPC server
	go func() {
		a3SServerObject.logger.WithFields(logrus.Fields{
			"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
		}).Info("Starting A3S gRPC Server")
		a3sGrpcServer = grpc.NewServer()
		a3s_grpc_api.RegisterA3SGrpcServiceServer(a3sGrpcServer, &A3S_GrpcServerStruct{})

		a3SServerObject.logger.WithFields(logrus.Fields{
			"Id":                           "e843ece9-b707-4c60-b1d8-14464305e68f",
			"localServerEngineLocalPort: ": common_config.A3SServer_port,
		}).Info("registerTestInstructionBackendServer for TestInstruction Backend Server started")
		a3sGrpcServer.Serve(lis)
	}()

}

/****************************************************/
// Stop Backend gRPC-server in a controlled way
func (a3SServerObject *A3SServerObjectStruct) StopGrpcServer() {

	a3SServerObject.logger.WithFields(logrus.Fields{}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")
	a3sGrpcServer.GracefulStop()

	a3SServerObject.logger.WithFields(logrus.Fields{
		"localServerEngineLocalPort: ": common_config.A3SServer_port,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
