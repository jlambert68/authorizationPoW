package userAuthorizationEngine

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"jlambert/authorizationPoW/common_config"
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
	remoteGrpcSecretMessageGeneratorServerConnection, err = grpc.Dial(addressToDial, grpc.WithInsecure())
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"ID":            "fbd24ed4-638a-43ac-a07b-c622f0ab325c",
			"addressToDial": addressToDial,
			"err.Error()":   err.Error(),
		}).Warning("Couldn't connect to SecretMessageGeneratorEngineServer")

		// TODO XXXXXXXXXXXXXXXXXXXXX
		//TODO skapa en meddelandtyp som gör att detta kan returneras
		// TODO XXXXXXXXXXXXXXXXXXXXX = &get

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
		secretMessageGrpcClient := secretMessageGenerator_grpc_api.NewSecretMessageGeneratorGrpcServiceClient(remoteGrpcSecretMessageGeneratorServerConnection)

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
