package secretMessageGeneratorEngine

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/secretMessageGenerator_grpc_api"
)

/***********************************************************************/
// Generate secret the message that is used when signing
func (secretMessageGenerator_GrpcServer *secretMessageGenerator_GrpcServerStruct) GenerateSecretFromInput(ctx context.Context, generateSecretFromInputRequest *secretMessageGenerator_grpc_api.GenerateSecretFromInputRequest) (*secretMessageGenerator_grpc_api.GenerateSecretFromInputResponse, error) {

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id": "85799f31-71b1-4c0e-9693-81fedd56bd41",
	}).Debug("Incoming 'GenerateSecretFromInput'")

	//
	var returnMessage *secretMessageGenerator_grpc_api.GenerateSecretFromInputResponse

	// Internal Secret data used for creating the public secret
	secretPart := "00c164e0-526e-4fe7-b20a-527cd0f961a1"
	secretToHash := secretPart + generateSecretFromInputRequest.UserId + generateSecretFromInputRequest.MessageUsedForSecret
	producedSecret := sha256.Sum256([]byte(secretToHash))
	a := producedSecret[:]
	producedSecretString := hex.EncodeToString(a)

	// Create return message
	returnMessage = &secretMessageGenerator_grpc_api.GenerateSecretFromInputResponse{
		Secret:   producedSecretString,
		Acknack:  true,
		Comments: "",
	}

	secretMessageGeneratorServerObject.logger.WithFields(logrus.Fields{
		"id":            "8ba74bad-a3c9-4018-b0c3-d26593d30f9f",
		"returnMessage": returnMessage,
	}).Debug("Leaveing 'GenerateSecretFromInput'")

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
