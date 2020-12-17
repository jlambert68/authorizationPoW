package userAuthorizationEngine

import (
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/****************************************************/
// userAuthorizationEngine Server object hodling "some" information
type userAuthorizationEngineServerObjectStruct struct {
	logger *logrus.Logger
}

var userAuthorizationEngineServerObject *userAuthorizationEngineServerObjectStruct

/****************************************************/
//  userAuthorizationEngine gRPC Server objects
var (
	userAuthorizationEngineGrpcServer *grpc.Server
	lis                               net.Listener
)

// Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type userAuthorizationEngine_GrpcServerStruct struct{}

/****************************************************/
// Standard gRPC Client connect towards other gRPC server
//--- Not needed for userAuthorizationEngine Server ---
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
// channel for doing a controlled exit from the program
var doControlledExitOfProgramChannel chan bool

/****************************************************/
// Used for only process cleanup once
var cleanupProcessed bool = false

func cleanup() {

	if cleanupProcessed == false {

		cleanupProcessed = true

		// Cleanup before close down application
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{}).Info("Clean up and shut down servers")

		// Stop Backend gRPC Server
		userAuthorizationEngineServerObject.StopGrpcServer()

	}
}

/****************************************************/
func userAuthorizationEngineServerMain() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// Saved objects having no expiration time will never be deleted
	databaseMemoryCache = cache.New(5*time.Minute, 10*time.Minute)

	// Set up userAuthorizationEngine-Object
	userAuthorizationEngineServerObject = &userAuthorizationEngineServerObjectStruct{}
	userAuthorizationEngineServerObject.InitLogger("")

	// Clean up when leaving. Is placed after logger because shutdown logs information
	defer cleanup()

	// Start userAuthorizationEngine gRPC-server
	userAuthorizationEngineServerObject.InitGrpcServer()

	// Exits when user press "ctrl-c" in terminal
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	// Channel receives data from gRPC-api when to end program in a controlled way
	doControlledExitOfProgramChannel = make(chan bool, 1)
	// Write message in terminal to show that process is alive
	<-doControlledExitOfProgramChannel

}

// Set up and start Backend gRPC-server
func (userAuthorizationEngineServerObject *userAuthorizationEngineServerObjectStruct) InitGrpcServer() {

	var err error

	// Find first non allocated port from defined start port
	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Info("userAuthorizationEngine gRPC Server tries to start")

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"Id": "ca3593b1-466b-4536-be91-5e038de178f4",
		"common_config.userAuthorizationEngineServer_port: ": common_config.UserAuthorizationServer_port,
	}).Debug("Start listening on:")
	lis, err = net.Listen("tcp", common_config.UserAuthorizationServer_port)

	if err != nil {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"Id": "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
			"common_config.userAuthorizationEngineServer_port: ": common_config.UserAuthorizationServer_port,
		}).Info("Success in listening on port:")

	}

	// Creates a new and start userAuthorizationEngineGrpcServer
	go func() {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
		}).Info("Starting userAuthorizationEngine gRPC Server")
		userAuthorizationEngineGrpcServer = grpc.NewServer()
		userAuthorizationEngine_grpc_api.RegisterUserAuthorizationGrpcServiceServer(userAuthorizationEngineGrpcServer, &userAuthorizationEngine_GrpcServerStruct{})

		err = userAuthorizationEngineGrpcServer.Serve(lis)
		if err != nil {
			userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
				"Id":    "2a5bf98b-e4ab-434c-9079-c1656b86bbbd",
				"err: ": err,
			}).Panic("Couldn't start 'userAuthorizationEngineGrpcServer'")

		} else {
			userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
				"Id":                           "e843ece9-b707-4c60-b1d8-14464305e68f",
				"localServerEngineLocalPort: ": common_config.UserAuthorizationServer_port,
			}).Info("registerTestInstructionBackendServer for TestInstruction Backend Server started")

		}
	}()

}

/****************************************************/
// Stop Backend gRPC-server in a controlled way
func (userAuthorizationEngineServerObject *userAuthorizationEngineServerObjectStruct) StopGrpcServer() {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{}).Info("Gracefull stop for: registerTaxiHardwareStreamServer")
	userAuthorizationEngineGrpcServer.GracefulStop()

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"localServerEngineLocalPort: ": common_config.UserAuthorizationServer_port,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
