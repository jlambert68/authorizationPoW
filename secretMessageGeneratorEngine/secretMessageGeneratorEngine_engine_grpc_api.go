package secretMessageGeneratorEngine

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/secretMessageGenerator_grpc_api"
)

/***********************************************************************/
// Saves Aggregated Signature for User in secretMessageGenerator Memory cache
func (secretMessageGenerator_GrpcServer *secretMessageGenerator_GrpcServerStruct) UpdateUsersAuthorizationSignature(ctx context.Context, updateUserAggregatedSignatureRequest *secretMessageGenerator_grpc_api.UpdateUserAggregatedSignatureRequest) (*secretMessageGenerator_grpc_api.AckNackResponse, error) {

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'UpdateUsersAuthorizationKey'")

	//
	var returnMessage *secretMessageGenerator_grpc_api.AckNackResponse

	// Save Users Aggregated Key in Memory Cache
	databaseMemoryCache.Set(
		updateUserAggregatedSignatureRequest.UserId,
		updateUserAggregatedSignatureRequest.UserAggregatedSignature,
		cache.NoExpiration)

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id": "7a9bbadd-ceb6-48ff-85a6-26f0787fb902",
		"updateUserAggregatedSignatureRequest.UserId":                  updateUserAggregatedSignatureRequest.UserId,
		"updateUserAggregatedSignatureRequest.UserAggregatedSignature": updateUserAggregatedSignatureRequest.UserAggregatedSignature,
	}).Debug("Saved users aggregates signature in memory cache")

	// Create return message
	returnMessage = &secretMessageGenerator_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "Users Aggregated Signature was saved in database",
	}

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id":            "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'UpdateUsersAuthorizationKey'")

	return returnMessage, nil

}

/***********************************************************************/
// Gets Aggregated Signature for User from secretMessageGenerator Memory cache
func (secretMessageGenerator_GrpcServer *secretMessageGenerator_GrpcServerStruct) GetUserAggregatedSignature(ctx context.Context, getUserAggregatedSignatureRequest *secretMessageGenerator_grpc_api.GetUserAggregatedSignatureRequest) (*secretMessageGenerator_grpc_api.GetUserAggregatedSignatureResponse, error) {

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id": "f240cd38-2e0c-45f8-a6ec-75723073cc7e",
	}).Debug("Incoming 'GetUserAggregatedSignature'")

	//
	var returnMessage *secretMessageGenerator_grpc_api.GetUserAggregatedSignatureResponse

	// Get Users Aggregated Key from Memory Cache
	agggregatedSignature, agggregatedSignatureWasFound := databaseMemoryCache.Get(getUserAggregatedSignatureRequest.UserId)
	if agggregatedSignatureWasFound {
		agggregatedSignatureAsString := agggregatedSignature.(string)

		secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
			"id": "c9fa971d-a944-4f1f-9904-a1c1998393fd",
			"getUserAggregatedSignatureRequest.UserId": getUserAggregatedSignatureRequest.UserId,
			"agggregatedSignature":                     agggregatedSignatureAsString,
		}).Debug("Users Aggregated Signature was found in Memory Cache")

		// Create return message when user was found
		returnMessage = &secretMessageGenerator_grpc_api.GetUserAggregatedSignatureResponse{
			UserId:                  getUserAggregatedSignatureRequest.UserId,
			UserAggregatedSignature: agggregatedSignatureAsString,
			Acknack:                 false,
			Comments:                "Users Aggregated Signature found in Memory Cache",
		}

	} else {

		secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
			"id": "3297d4ee-c10d-4390-94d2-8581de9ed15b",
			"getUserAggregatedSignatureRequest.UserId": getUserAggregatedSignatureRequest.UserId,
			"agggregatedSignature":                     agggregatedSignature,
		}).Debug("Users Aggregated Signature was not found in Memory Cache")

		// Create return message when user was not found
		returnMessage = &secretMessageGenerator_grpc_api.GetUserAggregatedSignatureResponse{
			UserId:                  getUserAggregatedSignatureRequest.UserId,
			UserAggregatedSignature: "",
			Acknack:                 false,
			Comments:                "Users Aggregated Signature couldn't be found in Memory Cache",
		}

	}

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id":            "6bdc7a0a-88cc-402c-9ddb-af59d7a300b1",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'GetUserAggregatedSignature'")

	return returnMessage, nil

}

/***********************************************************************/
// Saves Aggregated Signature for User in secretMessageGenerator Memory cache
func (secretMessageGenerator_GrpcServer *secretMessageGenerator_GrpcServerStruct) ShutDownsecretMessageGeneratorServer(ctx context.Context, emptyParameter *secretMessageGenerator_grpc_api.EmptyParameter) (*secretMessageGenerator_grpc_api.AckNackResponse, error) {

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id": "b67c80c8-d3b8-465d-af4a-19e4a0a7148f",
	}).Debug("Incoming 'ShutDownsecretMessageGeneratorServer'")

	//
	var returnMessage *secretMessageGenerator_grpc_api.AckNackResponse

	// Create return message
	returnMessage = &secretMessageGenerator_grpc_api.AckNackResponse{
		Acknack:  true,
		Comments: "secretMessageGenerator Server will shutdown",
	}

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id":            "045c72a1-d248-47ff-9ee6-d92b055a4582",
		"returnMessage": returnMessage,
	}).Debug("secretMessageGenerator Server will soon shutdown")

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id":            "9fe67ea7-c903-42de-8029-7811aa8a0a12",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'ShutDownsecretMessageGeneratorServer'")

	// Start shut shutdown after leaving this method
	defer func() {
		doControlledExitOfProgramChannel <- true
	}()

	return returnMessage, nil
}
