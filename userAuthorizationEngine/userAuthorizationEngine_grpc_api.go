package userAuthorizationEngine

import (
	"context"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
)

/***********************************************************************/
// Do a user have the correct rights to execute a specific API
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) IsUserAuthorized(ctx context.Context, userAuthorizationRequest *userAuthorizationEngine_grpc_api.UserAuthorizationRequest) (*userAuthorizationEngine_grpc_api.UserAuthorizationResponse, error) {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'UserAuthorization'")

	//
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizationResponse

	// Create return message
	returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
		UserIsAllowedToExecuteCallingApi: true,
		Acknack:                          true,
		Comments:                         "",
	}

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'UserAuthorization'")

	return returnMessage, nil

}

/***********************************************************************/
// Saves Aggregated Signature for User in secretMessageGenerator Memory cache
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) ShutDownUserAuthorizationServer(ctx context.Context, emptyParameter *userAuthorizationEngine_grpc_api.EmptyParameter) (*userAuthorizationEngine_grpc_api.AckNackResponse, error) {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "b67c80c8-d3b8-465d-af4a-19e4a0a7148f",
	}).Debug("Incoming 'ShutDownUserAuthorizationServer'")

	//
	var returnMessage *userAuthorizationEngine_grpc_api.AckNackResponse

	// Create return message
	returnMessage = &userAuthorizationEngine_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "userAuthorizationEngine Server will shutdown",
	}

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "045c72a1-d248-47ff-9ee6-d92b055a4582",
		"returnMessage": returnMessage,
	}).Debug("userAuthorizationEngine Server will soon shutdown")

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "9fe67ea7-c903-42de-8029-7811aa8a0a12",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'ShutDownUserAuthorizationServer'")

	// Start shut shutdown after leaving this method
	defer func() {
		doControlledExitOfProgramChannel <- true
	}()

	return returnMessage, nil
}
