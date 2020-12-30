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
// Combine users input Accounts, AccountTypes & Companies with Users Authorized Accounts, AccountTypes & Companies
// Used for validating that user has the right to use that input data

type combinationInputAndAuthorizationStruct struct {
	userId                    string
	company                   string
	CallingAPI                int
	userInputAccouts          []userAuthorizationEngine_grpc_api.Account
	userInputAccoutTypes      []userAuthorizationEngine_grpc_api.AccountType
	userInputCompanies        []userAuthorizationEngine_grpc_api.Company
	userAuthorizedAccouts     []userAuthorizationEngine_grpc_api.Account
	userAuthorizedAccoutTypes []userAuthorizationEngine_grpc_api.AccountType
	userAuthorizedCompanies   []userAuthorizationEngine_grpc_api.Company
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

func (userRequestsServerObject *userRequestsServerObjectStruct) combineUserInputWithAuthorizedData(combinationInputAndAuthorization combinationInputAndAuthorizationStruct) (bool, error) {

	//var err error
	var concatenatedAccounts []string
	var concatenatedAccountTypes []string
	var concatenateCompanies []string

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
	for _, tempAccount := range combinationInputAndAuthorization.userAuthorizedAccouts {
		concatenatedAccountTypes = append(concatenatedAccountTypes, tempAccount.Account)
	}
	for _, tempAccount := range combinationInputAndAuthorization.userInputAccouts {
		concatenatedAccountTypes = append(concatenatedAccountTypes, tempAccount.Account)
	}
	// Sort accounts and remove duplicates
	unique.Slice(concatenatedAccountTypes, lessString(concatenatedAccountTypes))

	// extract companies and concatenate into []string
	for _, tempAccount := range combinationInputAndAuthorization.userAuthorizedAccouts {
		concatenateCompanies = append(concatenateCompanies, tempAccount.Account)
	}
	for _, tempAccount := range combinationInputAndAuthorization.userInputAccouts {
		concatenateCompanies = append(concatenateCompanies, tempAccount.Account)
	}
	// Sort accounts and remove duplicates
	unique.Slice(concatenateCompanies, lessString(concatenateCompanies))

	return false, nil

}

/*
func strings(vs ...string) *[]string { return &vs }
func lessString(v interface{}) func(i, j int) bool {
	s := *v.(*[]string)
	return func(i, j int) bool { return s[i] < s[j] }


}

*/
