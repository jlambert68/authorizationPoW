package userRequestsEngine

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
)

/*******************************************************************/
// Get users authorized accounts, via gRPC-call to userAuthorizationEngine
//
func (userRequestsServerObject *userRequestsServerObjectStruct) getUserAuthorizedAccounts(userAuthorizedAccountsRequest userAuthorizationEngine_grpc_api.UserAuthorizedAccountsRequest) (userAuthorizedAccountsResponse *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.UserAuthorizationServer_address + common_config.UserAuthorizationServer_port

	// Set up connection to AuthorizationEngineServer
	remoteGrpcAuthorizationEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "fbd24ed4-638a-43ac-a07b-c622f0ab325c",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to AuthorizationEngineServer")

		userAuthorizedAccountsResponse = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse{
			UserId:    userAuthorizedAccountsRequest.UserId,
			CompanyId: userAuthorizedAccountsRequest.CompanyId,
			Acknack:   false,
			Comments:  "Couldn't connect to AuthorizationEngineServer: " + err.Error(),
			Accounts:  nil,
		}

		return userAuthorizedAccountsResponse

	} else {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "14d029db-0031-4837-b139-7b04b707fabf",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to AuthorizationEngineServer")

		// Creates a new AuthorizationEngineServer-Client
		authorizationGrpcClient := userAuthorizationEngine_grpc_api.NewUserAuthorizationGrpcServiceClient(remoteGrpcAuthorizationEngineServerConnection)

		// Call authrization server for list of authorized accounts
		ctx := context.Background()
		userAuthorizedAccountsResponse, err := authorizationGrpcClient.ListUsersAuthorizedAccounts(ctx, &userAuthorizedAccountsRequest)

		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "364e2a90-1c8b-47df-be64-b73457317911",
				"returnMessage": userAuthorizedAccountsResponse,
				"err.Error()":   err.Error(),
			}).Error("Problem to register client AuthorizationEngineServer")

			return &userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse{
				UserId:    userAuthorizedAccountsRequest.UserId,
				CompanyId: userAuthorizedAccountsRequest.CompanyId,
				Acknack:   false,
				Comments:  "Problem to register client AuthorizationEngineServerv: " + err.Error(),
				Accounts:  nil,
			}

		} else {
			// Retrun

			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "116024c5-268b-4688-97ca-272ab3db385f",
				"returnMessage": userAuthorizedAccountsResponse,
				"error":         err,
			}).Debug("Success in connecting AuthorizationEngineServer for authorized accounts")

			return userAuthorizedAccountsResponse
		}
	}
}
