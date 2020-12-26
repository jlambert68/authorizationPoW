package userRequestsEngine

import (
	"context"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/userRequests_grpc_api"
)

/***********************************************************************/
// List accounts that belongs to Company and that the user is authorized to see
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) ListAccounts(ctx context.Context, listAccountsRequest *userRequests_grpc_api.ListAccountsRequest) (*userRequests_grpc_api.ListAccountsResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'ListAccounts'")

	//
	var returnMessage *userRequests_grpc_api.ListAccountsResponse

	// TODO Check if user should have access to this function has access to
	hasUserAccesToThisFunction := true

	if hasUserAccesToThisFunction == false {
		// User hasn't got access to function
		returnMessage = &userRequests_grpc_api.ListAccountsResponse{
			UserId:    listAccountsRequest.UserId,
			CompanyId: listAccountsRequest.CompanyId,
			Acknack:   false,
			Comments:  "User hasn't got access to this function",
			Accounts:  nil,
		}

		return returnMessage, nil

	} else {
		// User has access to function
		// TODO generate list of accounts that user
		usersAccounts := []string{"283592388-31", "12412412-31"}

		// Create return message based on SQL question
		returnMessage = userRequestsServerObject.sqlListAccounts(listAccountsRequest, usersAccounts)

		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"id":            "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
			"returnMessage": returnMessage,
		}).Debug("Leaveing 'ListAccounts'")

		return returnMessage, nil
	}
}

/***********************************************************************/
// List accounts that is of the type the user provided
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) ListAccountsBaseOnProvidedType(ctx context.Context, listAccountsBasedOnProvidedTypeRequest *userRequests_grpc_api.ListAccountsBasedOnProvidedTypeRequest) (*userRequests_grpc_api.ListAccountsBasedOnProvidedTypeResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "8c30e65e-a1b9-47a8-8c11-f502fe92b51a",
	}).Debug("Incoming 'ListAccountsBasedOnProvidedTypeRequest'")

	//
	var returnMessage *userRequests_grpc_api.ListAccountsBasedOnProvidedTypeResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.ListAccountsBasedOnProvidedTypeResponse{
		UserId:       listAccountsBasedOnProvidedTypeRequest.UserId,
		CompanyId:    listAccountsBasedOnProvidedTypeRequest.CompanyId,
		AccounType:   listAccountsBasedOnProvidedTypeRequest.AccounType,
		Acknack:      true,
		Comments:     "",
		JsonResponse: "{}",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "842a4db2-864d-4154-a714-a0f52f41f56f",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'ListAccountsBasedOnProvidedTypeRequest'")

	return returnMessage, nil

}

/***********************************************************************/
// Add an account of a certain type
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) AddAccount(ctx context.Context, addAccountRequest *userRequests_grpc_api.AddAccountRequest) (*userRequests_grpc_api.AddAccountResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "1525cf10-2bb8-4513-b601-267ffb64d865",
	}).Debug("Incoming 'AddAccount'")

	//
	var returnMessage *userRequests_grpc_api.AddAccountResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.AddAccountResponse{
		UserId:    addAccountRequest.UserId,
		CompanyId: addAccountRequest.CompanyId,
		Acknack:   true,
		Comments:  "",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "4f6f0064-a4fe-45b6-961d-e275baf5a097",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'AddAccount'")

	return returnMessage, nil

}

/***********************************************************************/
// Delete an account
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) DeleteAccount(ctx context.Context, deleteAccountRequest *userRequests_grpc_api.DeleteAccountRequest) (*userRequests_grpc_api.DeleteAccountResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "6212dc39-3ac9-4834-ae35-735ba60830d3",
	}).Debug("Incoming 'DeleteAccount'")

	//
	var returnMessage *userRequests_grpc_api.DeleteAccountResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.DeleteAccountResponse{
		UserId:    deleteAccountRequest.UserId,
		CompanyId: deleteAccountRequest.CompanyId,
		Acknack:   true,
		Comments:  "",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "2f99bb9e-384d-4010-a826-bd67ab838351",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'DeleteAccount'")

	return returnMessage, nil

}

/***********************************************************************/
// Add an account type
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) AddAccountType(ctx context.Context, addAccountTypeRequest *userRequests_grpc_api.AddAccountTypeRequest) (*userRequests_grpc_api.AddAccountTypeResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "9d52c8e4-8d6c-46d2-9133-a190f2342999",
	}).Debug("Incoming 'AddAccountType'")

	//
	var returnMessage *userRequests_grpc_api.AddAccountTypeResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.AddAccountTypeResponse{
		UserId:    addAccountTypeRequest.UserId,
		CompanyId: addAccountTypeRequest.CompanyId,
		Acknack:   true,
		Comments:  "",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "a62b062b-8c05-4c0a-a828-4da9024b1b2d",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'AddAccountType'")

	return returnMessage, nil

}

/***********************************************************************/
// Delete an account type
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) DeleteAccountType(ctx context.Context, deleteAccountTypeTypeRequest *userRequests_grpc_api.DeleteAccountTypeTypeRequest) (*userRequests_grpc_api.DeleteAccountTypeResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "6a00a0b5-bd62-4e17-8943-522668ba573b",
	}).Debug("Incoming 'DeleteAccountType'")

	//
	var returnMessage *userRequests_grpc_api.DeleteAccountTypeResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.DeleteAccountTypeResponse{
		UserId:    deleteAccountTypeTypeRequest.UserId,
		CompanyId: deleteAccountTypeTypeRequest.CompanyId,
		Acknack:   true,
		Comments:  "",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "32e78fca-d92e-4cf9-9b49-a6f54dbff1dd",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'DeleteAccountType'")

	return returnMessage, nil

}

/***********************************************************************/
// Update company information
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) UpdateCompanyInformation(ctx context.Context, updateCompanyInformationRequest *userRequests_grpc_api.UpdateCompanyInformationRequest) (*userRequests_grpc_api.UpdateCompanyInformationResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "8d322f01-e21a-49fa-81eb-79a127bc2b5c",
	}).Debug("Incoming 'UpdateCompanyInformation'")

	//
	var returnMessage *userRequests_grpc_api.UpdateCompanyInformationResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.UpdateCompanyInformationResponse{
		UserId:    updateCompanyInformationRequest.UserId,
		CompanyId: updateCompanyInformationRequest.CompanyId,
		Acknack:   true,
		Comments:  "",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "ac412e0c-1586-4477-8a43-b43a5ad226a2",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'UpdateCompanyInformation'")

	return returnMessage, nil

}

/***********************************************************************/
/***********************************************************************/
// Shut down engine in a controlled way
func (userRequests_GrpcServer *userRequests_GrpcServerStruct) ShutDownUserRequestsServer(ctx context.Context, emptyParameter *userRequests_grpc_api.EmptyParameter) (*userRequests_grpc_api.AckNackResponse, error) {

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id": "b67c80c8-d3b8-465d-af4a-19e4a0a7148f",
	}).Debug("Incoming 'ShutDownUserRequestsServer'")

	//
	var returnMessage *userRequests_grpc_api.AckNackResponse

	// Create return message
	returnMessage = &userRequests_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "ShutDownUserRequestsServer Server will shutdown",
	}

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "045c72a1-d248-47ff-9ee6-d92b055a4582",
		"returnMessage": returnMessage,
	}).Debug("ShutDownUserRequestsServer Server will soon shutdown")

	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"id":            "9fe67ea7-c903-42de-8029-7811aa8a0a12",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'ShutDownUserRequestsServer'")

	// Start shut shutdown after leaving this method
	defer func() {
		doControlledExitOfProgramChannel <- true
	}()

	return returnMessage, nil
}
