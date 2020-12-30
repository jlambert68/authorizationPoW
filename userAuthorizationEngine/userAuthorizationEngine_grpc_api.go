package userAuthorizationEngine

import (
	"context"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/secretMessageGenerator_grpc_api"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
)

/***********************************************************************/
// Do a user have the correct rights to execute a specific API
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) IsUserAuthorized(ctx context.Context, userAuthorizationRequest *userAuthorizationEngine_grpc_api.UserAuthorizationRequest) (*userAuthorizationEngine_grpc_api.UserAuthorizationResponse, error) {

	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizationResponse

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'UserAuthorization'")

	// Get users authorized accounts
	userAuthorizedAccountsRequest := &userAuthorizationEngine_grpc_api.UserAuthorizedAccountsRequest{
		UserId:    userAuthorizationRequest.UserId,
		CompanyId: userAuthorizationRequest.CompanyId,
	}
	userAuthorizedAccountsResponse := userAuthorizationEngineServerObject.sqlListUsersAuthorizedAccounts(userAuthorizedAccountsRequest)

	// Create return message, and exit, telling that there were no success in reading database
	if userAuthorizedAccountsResponse.Acknack == false {
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
			UserIsAllowedToExecuteCallingApi: false,
			Acknack:                          false,
			Comments:                         userAuthorizedAccountsResponse.Comments,
		}
	} else {

		// Get users authorized account types
		userAuthorizedAccountTypesRequest := &userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesRequest{
			UserId:    userAuthorizationRequest.UserId,
			CompanyId: userAuthorizationRequest.CompanyId,
		}
		userAuthorizedAccountTypesResponse := userAuthorizationEngineServerObject.sqlListUsersAuthorizedAccountTypes(userAuthorizedAccountTypesRequest)

		// Create return message, and exit, telling that there were no success in reading database
		if userAuthorizedAccountTypesResponse.Acknack == false {
			returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
				UserIsAllowedToExecuteCallingApi: false,
				Acknack:                          false,
				Comments:                         userAuthorizedAccountTypesResponse.Comments,
			}
		} else {

			// Get users authorized companies
			userAuthorizedCompaniesRequest := &userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesRequest{
				UserId: userAuthorizationRequest.UserId,
			}
			userAuthorizedCompaniesResponse := userAuthorizationEngineServerObject.sqlListUsersAuthorizedCompanies(userAuthorizedCompaniesRequest)

			// Create return message, and exit, telling that there were no success in reading database
			if userAuthorizedCompaniesResponse.Acknack == false {
				returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
					UserIsAllowedToExecuteCallingApi: false,
					Acknack:                          false,
					Comments:                         userAuthorizedCompaniesResponse.Comments,
				}

			} else {
				// OK and all authorized data was retrieved

				// Concatenate and remove duplicate data
				userAuthorizedAccountsResponse := userAuthorizedAccountsResponse.GetAccounts()
				userAuthorizedAccountTypesResponse := userAuthorizedAccountTypesResponse.GetAccountTypes()
				userAuthorizedCompaniesResponse := userAuthorizedCompaniesResponse.GetCompanies()
				combinationInputAndAuthorizationData := &combinationInputAndAuthorizationStruct{
					userId:                    userAuthorizationRequest.UserId,
					company:                   userAuthorizationRequest.CompanyId,
					CallingAPI:                userAuthorizationRequest.CallingApi,
					userInputAccouts:          userAuthorizationRequest.Accounts,
					userInputAccoutTypes:      userAuthorizationRequest.AccountTypes,
					userInputCompanies:        userAuthorizationRequest.Companies,
					userAuthorizedAccouts:     userAuthorizedAccountsResponse,
					userAuthorizedAccoutTypes: userAuthorizedAccountTypesResponse,
					userAuthorizedCompanies:   userAuthorizedCompaniesResponse,
				}

				userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
					"id":                                   "8c39c5bb-3917-4198-b7af-ccae6f54bfee",
					"combinationInputAndAuthorizationData": combinationInputAndAuthorizationData,
				}).Debug("Concatenate and remove duplicates for accounts, accountTypes and companies")

				concatenatedAccounts, concatenatedAccountTypes, concatenateCompanies := userAuthorizationEngineServerObject.combineUserInputWithAuthorizedData(*combinationInputAndAuthorizationData)

				userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
					"id":                       "d26204c5-b284-4b00-b705-bdcd05041d6a",
					"concatenatedAccounts":     concatenatedAccounts,
					"concatenatedAccountTypes": concatenatedAccountTypes,
					"concatenateCompanies":     concatenateCompanies,
				}).Debug("Done with Concatenate and remove duplicates for accounts, accountTypes and companies")

				// Generate Secret from concatenated data
				generateSecretFromInputRequest := &secretMessageGenerator_grpc_api.GenerateSecretFromInputRequest{
					UserId:       userAuthorizationRequest.UserId,
					Company:      userAuthorizationRequest.CompanyId,
					Accounts:     concatenatedAccounts,
					AccountTypes: concatenatedAccountTypes,
					Companies:    concatenateCompanies,
				}
				getSecretFromUserDataResponse := userAuthorizationEngineServerObject.getSecretFromUserData(generateSecretFromInputRequest)

				//Do a cryptographic validation of generated secret
				// TODO XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

			}
		}
	}

	// Create return message
	returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
		UserIsAllowedToExecuteCallingApi: true,
		Acknack:                          true,
		Comments:                         "",
	}

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
		"returnMessage": returnMessage,
	}).Debug("Leaving 'UserAuthorization'")

	return returnMessage, nil

}

/***********************************************************************/
// List users authorized accounts
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) ListUsersAuthorizedAccounts(ctx context.Context, userAuthorizedAccountsRequest *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsRequest) (*userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse, error) {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "b5499021-238f-4880-b53c-e91a9c119837",
	}).Debug("Incoming 'ListUsersAuthorizedAccounts'")

	//
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse

	// Create return message
	returnMessage = userAuthorizationEngineServerObject.sqlListUsersAuthorizedAccounts(userAuthorizedAccountsRequest)

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "29693194-cd45-4f8e-8194-e14ce5a730f6",
		"returnMessage": returnMessage,
	}).Debug("Leaving 'ListUsersAuthorizedAccounts'")

	return returnMessage, nil

}

/***********************************************************************/
// List users authorized account types
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) ListUsersAuthorizedAccountTypes(ctx context.Context, userAuthorizedAccountTypesRequest *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesRequest) (*userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse, error) {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "b5499021-238f-4880-b53c-e91a9c119837",
	}).Debug("Incoming 'ListUsersAuthorizedAccountTypes'")

	//
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse

	// Create return message
	returnMessage = userAuthorizationEngineServerObject.sqlListUsersAuthorizedAccountTypes(userAuthorizedAccountTypesRequest)

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "29693194-cd45-4f8e-8194-e14ce5a730f6",
		"returnMessage": returnMessage,
	}).Debug("Leaving 'ListUsersAuthorizedAccountTypes'")

	return returnMessage, nil

}

/***********************************************************************/
// List users authorized companies
func (userAuthorizationEngine_GrpcServer *userAuthorizationEngine_GrpcServerStruct) ListUsersAuthorizedCompanies(ctx context.Context, userAuthorizedCompaniesRequest *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesRequest) (*userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse, error) {

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id": "6759e53e-61b1-4825-89c8-e59647a942a2",
	}).Debug("Incoming 'ListUsersAuthorizedCompanies'")

	//
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse

	// Create return message
	returnMessage = userAuthorizationEngineServerObject.sqlListUsersAuthorizedCompanies(userAuthorizedCompaniesRequest)

	userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
		"id":            "90670e25-fd43-4f0a-8c92-3428a1c9298c",
		"returnMessage": returnMessage,
	}).Debug("Leaving 'ListUsersAuthorizedCompanies'")

	return returnMessage, nil

}

/***********************************************************************/
// Shut down Authorization server
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
	}).Debug("Leaving 'ShutDownUserAuthorizationServer'")

	// Start shut shutdown after leaving this method
	defer func() {
		doControlledExitOfProgramChannel <- true
	}()

	return returnMessage, nil
}
