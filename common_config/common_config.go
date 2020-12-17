package common_config

import "github.com/sirupsen/logrus"

/*************************************************************/
// gRPC-addresses and ports

// A3S Server
const A3SServer_address = "127.0.0.1"
const A3SServer_port = ":6000"

// SecretMessageGenerator Server
const SecretMessageGeneratorServer_address = "127.0.0.1"
const SecretMessageGeneratorServer_port = ":6001"

// SetUpUserAuthorizationRules Server
const SetUpUserAuthorizationRulesServer_address = "127.0.0.1"
const SetUpUserAuthorizationRulesServer_port = ":6002"

// UserRequests Server
const UserRequestsServer_address = "127.0.0.1"
const UserRequestsServer_port = ":6003"

// UserRequests Server
const UserAuthorizationServer_address = "127.0.0.1"
const UserAuthorizationServer_port = ":6004"

/*************************************************************/
// Logrus debug level
//const LoggingLevel = logrus.DebugLevel
//const LoggingLevel = logrus.InfoLevel
const LoggingLevel = logrus.DebugLevel // InfoLevel
