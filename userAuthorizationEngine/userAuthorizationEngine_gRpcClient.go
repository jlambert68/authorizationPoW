package userAuthorizationEngine

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
	"jlambert/authorizationPoW/grpc_api/a3s_grpc_api"
	"jlambert/authorizationPoW/grpc_api/secretMessageGenerator_grpc_api"
)

/*******************************************************************/
// Generate a secret message from user data via gRPC-call to secretMessageGeneratorEngine
//
type getSecretFromUserDataResponseStruct struct {
	userId   string
	company  string
	ackNack  bool
	comments string
	secret   string
}

func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) getSecretFromUserData(generateSecretFromInputRequest *secretMessageGenerator_grpc_api.GenerateSecretFromInputRequest) (getSecretFromUserDataResponse getSecretFromUserDataResponseStruct) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.SecretMessageGeneratorServer_address + common_config.SecretMessageGeneratorServer_port

	// Set up connection to SecretMessageGeneratorEngineServer
	remoteGrpcSA3sSimulatorEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "fbd24ed4-638a-43ac-a07b-c622f0ab325c",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to SecretMessageGeneratorEngineServer")

		getSecretFromUserDataResponse = getSecretFromUserDataResponseStruct{
			userId:   generateSecretFromInputRequest.GetUserId(),
			company:  generateSecretFromInputRequest.GetCompany(),
			ackNack:  false,
			comments: "Couldn't connect to SecretMessageGeneratorEngineServe",
			secret:   "",
		}

		return getSecretFromUserDataResponse

	} else {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "14d029db-0031-4837-b139-7b04b707fabf",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to SecretMessageGeneratorEngineServer")

		// Creates a new SecretMessageGeneratorEngineServer-Client
		secretMessageGrpcClient := secretMessageGenerator_grpc_api.NewSecretMessageGeneratorGrpcServiceClient(remoteGrpcSA3sSimulatorEngineServerConnection)

		// Call secretMessage server for generating the secret message
		ctx := context.Background()
		generateSecretFromInputResponse, err := secretMessageGrpcClient.GenerateSecretFromInput(ctx, generateSecretFromInputRequest)

		if err != nil {
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":            "364e2a90-1c8b-47df-be64-b73457317911",
				"returnMessage": generateSecretFromInputResponse,
				"err.Error()":   err.Error(),
			}).Error("Problem to register client at SecretMessageGeneratorEngineServer")

			getSecretFromUserDataResponse = getSecretFromUserDataResponseStruct{
				userId:   generateSecretFromInputRequest.GetUserId(),
				company:  generateSecretFromInputRequest.GetCompany(),
				ackNack:  false,
				comments: "Problem to register client at SecretMessageGeneratorEngineServer: " + err.Error(),
				secret:   "",
			}

			return getSecretFromUserDataResponse

		} else {
			// Success in calling gRPC at secretMessage-server

			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":                              "116024c5-268b-4688-97ca-272ab3db385f",
				"generateSecretFromInputResponse": generateSecretFromInputResponse,
				"error":                           err.Error(),
			}).Debug("Success in connecting SecretMessageGeneratorEngineServer to generate secret")

			getSecretFromUserDataResponse = getSecretFromUserDataResponseStruct{
				userId:   generateSecretFromInputRequest.GetUserId(),
				company:  generateSecretFromInputRequest.GetCompany(),
				ackNack:  generateSecretFromInputResponse.GetAcknack(),
				comments: generateSecretFromInputResponse.GetComments(),
				secret:   generateSecretFromInputResponse.GetSecret(),
			}

			return getSecretFromUserDataResponse
		}
	}
}

/*******************************************************************/
// Store users aggregated signature via gRPC-call to A3S simulatorEngine
//
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) updateUserAggregatedSignature(updateUserAggregatedSignatureRequest *a3s_grpc_api.UpdateUserAggregatedSignatureRequest) (retunMessage a3s_grpc_api.AckNackResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.A3SServer_address + common_config.A3SServer_port

	// Set up connection to A3SSImulatorEngine
	remoteGrpcSA3sSimulatorEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "2d01433a-878b-44b9-afa5-c0cc4eb85355",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to A3SSimulatorEngineServer")

		retunMessage = a3s_grpc_api.AckNackResponse{
			Acknack:  false,
			Comments: "Couldn't connect to A3SSimulatorEngineServer",
		}

		return retunMessage

	} else {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "c287b145-74a9-446d-803d-d5f4219f2f0a",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to A3SSimulatorEngineServer")

		// Creates a new A3SSimulatorEngineServer-Client
		a3sSimulatorEngineGrpcClient := a3s_grpc_api.NewA3SGrpcServiceClient(remoteGrpcSA3sSimulatorEngineServerConnection)

		// Call A3SSimulatorEngine server for generating for storing aggregated signature
		ctx := context.Background()
		a3SResponse, err := a3sSimulatorEngineGrpcClient.UpdateUsersAuthorizationSignature(ctx, updateUserAggregatedSignatureRequest)

		if err != nil {
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":          "724cc259-4951-4d86-8e75-21b1db4bba58",
				"a3SResponse": a3SResponse,
				"err.Error()": err.Error(),
			}).Error("Problem to register client at A3sSimulatorEngineServer")

			retunMessage = a3s_grpc_api.AckNackResponse{
				Acknack:  false,
				Comments: "Problem to register client at A3sSimulatorEngineServer," + err.Error(),
			}

			return retunMessage

		} else {
			// Success in calling gRPC at A3sSumulator-server

			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":          "1fb6e4dd-01ce-4211-a55b-992d5ebf4474",
				"a3SResponse": a3SResponse,
			}).Debug("Success in connecting A3sSimulatorEngineServer to store aggregated signature")

			retunMessage = a3s_grpc_api.AckNackResponse{
				Acknack:  true,
				Comments: "",
			}

			return retunMessage
		}
	}
}

/*******************************************************************/
// Get users aggregated signature via gRPC-call to A3S simulatorEngine
//
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) getUserAggregatedSignature(getUserAggregatedSignatureRequest *a3s_grpc_api.GetUserAggregatedSignatureRequest) (getUserAggregatedSignatureResponse a3s_grpc_api.GetUserAggregatedSignatureResponse) {

	var err error
	var addressToDial string

	// Find parents address and port to call
	addressToDial = common_config.A3SServer_address + common_config.A3SServer_port

	// Set up connection to A3SSImulatorEngine
	remoteGrpcSA3sSimulatorEngineServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "25cbe1a7-669c-434e-aab6-b684a8f0d7b0",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to A3SSimulatorEngineServer")

		getUserAggregatedSignatureResponse = a3s_grpc_api.GetUserAggregatedSignatureResponse{
			UserId:                  getUserAggregatedSignatureRequest.UserId,
			UserAggregatedSignature: "",
			Acknack:                 false,
			Comments:                "Couldn't connect to A3SSimulatorEngineServer",
		}

		return getUserAggregatedSignatureResponse

	} else {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "d3539caf-dd28-4b96-8dde-b02e1201add0",
			"addressToDial": addressToDial,
		}).Debug("gRPC connection OK to A3SSimulatorEngineServer")

		// Creates a new A3SSimulatorEngineServer-Client
		a3sSimulatorEngineGrpcClient := a3s_grpc_api.NewA3SGrpcServiceClient(remoteGrpcSA3sSimulatorEngineServerConnection)

		// Call A3SSimulatorEngine server for generating for storing aggregated signature
		ctx := context.Background()
		getUserAggregatedSignatureResponse, err := a3sSimulatorEngineGrpcClient.GetUserAggregatedSignature(ctx, getUserAggregatedSignatureRequest)

		if err != nil {
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":                                 "6a4582e4-bd28-4b3f-9b9e-c8e593aad382",
				"getUserAggregatedSignatureResponse": getUserAggregatedSignatureResponse,
				"err.Error()":                        err.Error(),
			}).Error("Problem to register client at A3sSimulatorEngineServer")

			return *getUserAggregatedSignatureResponse

		} else {
			// Success in calling gRPC at A3sSumulator-server

			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"ID":                                 "3933fd30-14c1-441e-94aa-ed54f6cc291b",
				"getUserAggregatedSignatureResponse": getUserAggregatedSignatureResponse,
			}).Debug("Success in connecting A3sSimulatorEngineServer to get aggregated signature")

			return *getUserAggregatedSignatureResponse
		}
	}
}
