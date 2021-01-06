package userRequestsEngine

import (
	"github.com/campoy/unique"
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

		// Call authorization server for list of authorized accounts
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
				Comments:  "Problem to register client AuthorizationEngineServer: " + err.Error(),
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

/*******************************************************************/
// Get users authorized account types, via gRPC-call to userAuthorizationEngine
//
func (userRequestsServerObject *userRequestsServerObjectStruct) getUserAuthorizedAccountTypes(userAuthorizedAccountTypesRequest userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesRequest) (userAuthorizedAccountTypesResponse *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.UserAuthorizationServer_address + common_config.UserAuthorizationServer_port

	// Set up connection to AuthorizationEngineServer
	remoteGrpcAuthorizationEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "7805d802-f341-425d-80f5-1f7a559492a6",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to AuthorizationEngineServer")

		userAuthorizedAccountTypesResponse = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse{
			UserId:       userAuthorizedAccountTypesRequest.UserId,
			CompanyId:    userAuthorizedAccountTypesRequest.CompanyId,
			Acknack:      false,
			Comments:     "Couldn't connect to AuthorizationEngineServer: " + err.Error(),
			AccountTypes: nil,
		}

		return userAuthorizedAccountTypesResponse

	} else {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "79419aa9-31fb-4eb2-8101-b47e46e3b58b",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to AuthorizationEngineServer")

		// Creates a new AuthorizationEngineServer-Client
		authorizationGrpcClient := userAuthorizationEngine_grpc_api.NewUserAuthorizationGrpcServiceClient(remoteGrpcAuthorizationEngineServerConnection)

		// Call authorization server for list of authorized account types
		ctx := context.Background()
		userAuthorizedAccountTypesResponse, err := authorizationGrpcClient.ListUsersAuthorizedAccountTypes(ctx, &userAuthorizedAccountTypesRequest)

		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "512ae342-c61f-43d8-9838-0c2b9fa05b1a",
				"returnMessage": userAuthorizedAccountTypesResponse,
				"err.Error()":   err.Error(),
			}).Error("Problem to register client AuthorizationEngineServer")

			return &userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse{
				UserId:       userAuthorizedAccountTypesRequest.UserId,
				CompanyId:    userAuthorizedAccountTypesRequest.CompanyId,
				Acknack:      false,
				Comments:     "Problem to register client AuthorizationEngineServer: " + err.Error(),
				AccountTypes: nil,
			}

		} else {
			// Retrun

			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "7746d490-a17c-49d4-876b-e9b6bcf74824",
				"returnMessage": userAuthorizedAccountTypesResponse,
				"error":         err,
			}).Debug("Success in connecting AuthorizationEngineServer for authorized accounts")

			return userAuthorizedAccountTypesResponse
		}
	}
}

/*******************************************************************/
// Get users authorized companies via gRPC-call to userAuthorizationEngine
//
func (userRequestsServerObject *userRequestsServerObjectStruct) getUserAuthorizedCompanies(userAuthorizedCompaniesRequest userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesRequest) (userAuthorizedCompaniesResponse *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.UserAuthorizationServer_address + common_config.UserAuthorizationServer_port

	// Set up connection to AuthorizationEngineServer
	remoteGrpcAuthorizationEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "b4b8a5c3-e2ae-4504-8263-9fc858a33ab1",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to AuthorizationEngineServer")

		userAuthorizedCompaniesResponse = &userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse{
			UserId:    userAuthorizedCompaniesRequest.UserId,
			Acknack:   false,
			Comments:  "Couldn't connect to AuthorizationEngineServer: " + err.Error(),
			Companies: nil,
		}

		return userAuthorizedCompaniesResponse

	} else {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "195db8d4-eb9b-4dcc-b5d9-ebd823a0e896",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to AuthorizationEngineServer")

		// Creates a new AuthorizationEngineServer-Client
		authorizationGrpcClient := userAuthorizationEngine_grpc_api.NewUserAuthorizationGrpcServiceClient(remoteGrpcAuthorizationEngineServerConnection)

		// Call authorization server for list of authorized companies
		ctx := context.Background()
		userAuthorizedCompaniesResponse, err := authorizationGrpcClient.ListUsersAuthorizedCompanies(ctx, &userAuthorizedCompaniesRequest)

		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "cd6bfa74-50ba-41cb-9524-db994e8aa46e",
				"returnMessage": userAuthorizedCompaniesResponse,
				"err.Error()":   err.Error(),
			}).Error("Problem to register client AuthorizationEngineServer")

			return &userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse{
				UserId:    userAuthorizedCompaniesRequest.UserId,
				Acknack:   false,
				Comments:  "Problem to register client AuthorizationEngineServer: " + err.Error(),
				Companies: nil,
			}

		} else {
			// Retrun

			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "f2d3db39-20c2-4646-9987-7029a0535d4e",
				"returnMessage": userAuthorizedCompaniesResponse,
				"error":         err,
			}).Debug("Success in connecting AuthorizationEngineServer for authorized accounts")

			return userAuthorizedCompaniesResponse
		}
	}
}

/*******************************************************************/
// Check if users is authorized to execute request, via gRPC-call to userAuthorizationEngine
//
func (userRequestsServerObject *userRequestsServerObjectStruct) isUserAuthorizedToExecute(userAuthorizationRequest *userAuthorizationEngine_grpc_api.UserAuthorizationRequest) (userAuthorizationResponse *userAuthorizationEngine_grpc_api.UserAuthorizationResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.UserAuthorizationServer_address + common_config.UserAuthorizationServer_port

	// Set up connection to AuthorizationEngineServer
	remoteGrpcAuthorizationEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "0a4653bb-4467-45cc-9cd4-8ae570d81f0d",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to AuthorizationEngineServer")

		userAuthorizationResponse = &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
			UserIsAllowedToExecuteCallingApi: false,
			Acknack:                          false,
			Comments:                         "Couldn't connect to AuthorizationEngineServer: " + err.Error(),
		}

		return userAuthorizationResponse

	} else {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"ID":            "57772d09-2d42-4479-b3aa-7904f11fbbca",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to AuthorizationEngineServer")

		// Creates a new AuthorizationEngineServer-Client
		authorizationGrpcClient := userAuthorizationEngine_grpc_api.NewUserAuthorizationGrpcServiceClient(remoteGrpcAuthorizationEngineServerConnection)

		// Call authorization server to check if user is authorized
		ctx := context.Background()
		userAuthorizationResponse, err := authorizationGrpcClient.IsUserAuthorized(ctx, userAuthorizationRequest)

		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "34d016d0-8840-43e6-acd8-95f9ce3f0285",
				"returnMessage": userAuthorizationResponse,
				"err.Error()":   err.Error(),
			}).Error("Problem to register client AuthorizationEngineServer")

			return &userAuthorizationEngine_grpc_api.UserAuthorizationResponse{
				UserIsAllowedToExecuteCallingApi: false,
				Acknack:                          false,
				Comments:                         "Problem to register client AuthorizationEngineServer: " + err.Error(),
			}

		} else {
			// Retrun

			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"ID":            "f2d3db39-20c2-4646-9987-7029a0535d4e",
				"returnMessage": userAuthorizationResponse,
				"error":         err,
			}).Debug("Success in connecting AuthorizationEngineServer for authorized accounts")

			return userAuthorizationResponse
		}
	}
}

/*******************************************************************/
// Combine users input Accounts, AccountTypes & Companies with Users Authorized Accounts, AccountTypes & Companies
// Used for validating that user has the right to use that input data

type combinationInputAndAuthorizationStruct struct {
	userId                    string
	company                   string
	CallingAPI                int32
	userInputAccouts          []userAuthorizationEngine_grpc_api.Account
	userInputAccoutTypes      []userAuthorizationEngine_grpc_api.AccountType
	userInputCompanies        []userAuthorizationEngine_grpc_api.Company
	userAuthorizedAccouts     []userAuthorizationEngine_grpc_api.Account
	userAuthorizedAccoutTypes []userAuthorizationEngine_grpc_api.AccountType
	userAuthorizedCompanies   []userAuthorizationEngine_grpc_api.Company
}

type combineUserInputWithAuthorizedDataResponseStruct struct {
	userId          string
	company         string
	CallingAPI      int32
	userAccouts     []userAuthorizationEngine_grpc_api.Account
	userAccoutTypes []userAuthorizationEngine_grpc_api.AccountType
	userCompanies   []userAuthorizationEngine_grpc_api.Company
}

func strings(vs ...string) *[]string { return &vs }
func lessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }
}

/*
func lessAccounts(v interface{}) func(i, j int) bool {
	s := *v.(*[]userAuthorizationEngine_grpc_api.Account)
	return func(i, j int) bool {
		return s[i] < s[j]
	}
}
*/

func (userRequestsServerObject *userRequestsServerObjectStruct) combineUserInputWithAuthorizedData(combinationInputAndAuthorization combinationInputAndAuthorizationStruct) combineUserInputWithAuthorizedDataResponseStruct {

	//var err error
	var concatenatedAccounts []string
	var concatenatedAccountTypes []string
	var concatenateCompanies []string
	var userAccouts []userAuthorizationEngine_grpc_api.Account
	var userAccoutTypes []userAuthorizationEngine_grpc_api.AccountType
	var userCompanies []userAuthorizationEngine_grpc_api.Company

	//s := []int{3, 5, 1, 7, 2, 3, 7, 5, 2}
	//less := func(i, j int) bool { return s[i] < s[j] }
	//unique.Slice(&s, less)

	//s := combinationInputAndAuthorization.userInputAccouts
	s := strings("one", "two", "three", "two", "one")

	//unique.Slice(&s, lessString)
	unique.Slice(s, lessString(s))

	// extract accounts and concatenate into []string
	for _, tempAccount := range combinationInputAndAuthorization.userAuthorizedAccouts {
		concatenatedAccounts = append(concatenatedAccounts, tempAccount.Account)
	}
	for _, tempAccount := range combinationInputAndAuthorization.userInputAccouts {
		concatenatedAccounts = append(concatenatedAccounts, tempAccount.Account)
	}
	// Sort accounts and remove duplicates
	unique.Slice(concatenatedAccounts, lessString(concatenatedAccounts))

	// extract account types and concatenate into []string
	for _, tempAccountType := range combinationInputAndAuthorization.userAuthorizedAccoutTypes {
		concatenatedAccountTypes = append(concatenatedAccountTypes, tempAccountType.AccountType)
	}
	for _, tempAccountType := range combinationInputAndAuthorization.userInputAccoutTypes {
		concatenatedAccountTypes = append(concatenatedAccountTypes, tempAccountType.AccountType)
	}
	// Sort accounts and remove duplicates
	unique.Slice(concatenatedAccountTypes, lessString(concatenatedAccountTypes))

	// extract companies and concatenate into []string
	for _, tempCompany := range combinationInputAndAuthorization.userAuthorizedCompanies {
		concatenateCompanies = append(concatenateCompanies, tempCompany.Company)
	}
	for _, tempCompany := range combinationInputAndAuthorization.userInputCompanies {
		concatenateCompanies = append(concatenateCompanies, tempCompany.Company)
	}
	// Sort accounts and remove duplicates
	unique.Slice(concatenateCompanies, lessString(concatenateCompanies))

	// Convert Accounts back into type used in gRPC
	for _, tempAccount := range concatenatedAccounts {
		tempAccountConverted := userAuthorizationEngine_grpc_api.Account{
			Account: tempAccount,
		}
		userAccouts = append(userAccouts, tempAccountConverted)
	}

	// Convert AccountTypeds back into type used in gRPC
	for _, tempAccountType := range concatenatedAccountTypes {
		tempAccountTypeConverted := userAuthorizationEngine_grpc_api.AccountType{
			AccountType: tempAccountType,
		}
		userAccoutTypes = append(userAccoutTypes, tempAccountTypeConverted)
	}

	// Convert Company back into type used in gRPC
	for _, tempCompany := range concatenatedAccountTypes {
		tempCompanyConverted := userAuthorizationEngine_grpc_api.Company{
			Company: tempCompany,
		}
		userCompanies = append(userCompanies, tempCompanyConverted)
	}

	// Create response message
	combineUserInputWithAuthorizedDataResponse := combineUserInputWithAuthorizedDataResponseStruct{
		userId:          combinationInputAndAuthorization.userId,
		company:         combinationInputAndAuthorization.company,
		CallingAPI:      combinationInputAndAuthorization.CallingAPI,
		userAccouts:     userAccouts,
		userAccoutTypes: userAccoutTypes,
		userCompanies:   userCompanies,
	}

	return combineUserInputWithAuthorizedDataResponse

}

/*
func strings(vs ...string) *[]string { return &vs }
func lessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }


}

*/
