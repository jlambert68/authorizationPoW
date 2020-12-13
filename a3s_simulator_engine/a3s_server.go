package a3s_simulator_engine

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/****************************************************/
// A3S Server object hodling "some" information
type A3SServerObject_struct struct {
	logger *logrus.Logger
}

var a3SServerObject *A3SServerObject_struct

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

		//log.Println("Close DB_session: %v", DB_session)
		//DB_session.Close()
	}
}

/****************************************************/
func BackendServer_main() {

	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	// Saved objects having no expiration time will never be deleted
	databaseMemoryCache = cache.New(5*time.Minute, 10*time.Minute)

	// Set up A3S-Object
	a3SServerObject = &A3SServerObject_struct{}

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
