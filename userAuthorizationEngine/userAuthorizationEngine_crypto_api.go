package userAuthorizationEngine

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/sirupsen/logrus"
	"jlambert/authorizationPoW/grpc_api/a3s_grpc_api"
	"math"
	"reflect"
	"strconv"
	//"golang.org/x/crypto/ed25519"
	//"golang.org/x/crypto/ed25519/cosi"
	"github.com/bford/golang-x-crypto/ed25519"
	"github.com/bford/golang-x-crypto/ed25519/cosi"
)

type userType struct {
	userId                         string
	userSecret                     [32]byte
	allowedFunctions               []string
	publicKeys                     []ed25519.PublicKey
	privateKeys                    []ed25519.PrivateKey
	aggregatedSignatureStoredInA3S []byte
}

type ruleStruct struct {
	ruleName   string
	publicKey  ed25519.PublicKey
	privateKey ed25519.PrivateKey
}

/***************************************************************/
// Generate all public- and private keys to be used for one user specific user
type usersRulesStruct struct {
	userName      string
	allowedFRules []string
	allRules      []string
	usersSecret   [32]byte
}

type usersKeysStruct struct {
	userId                         string
	userSecret                     [32]byte
	allowedRules                   []string
	publicKeys                     []ed25519.PublicKey
	privateKeys                    []ed25519.PrivateKey
	allPublicKeys                  []ed25519.PublicKey
	aggregatedSignatureStoredInA3S []byte
}

type returnMessageStruct struct {
	userId        string
	publicKeys    []ed25519.PublicKey
	allPublicKeys []ed25519.PublicKey
}

func generatePublicAndPrivateKeysForUser(usersRules usersRulesStruct) (returnMessage returnMessageStruct) {

	type ruleKeys struct {
		functionName string
		publicKey    ed25519.PublicKey
		privateKey   ed25519.PrivateKey
	}

	var allRulesWithKeys []ruleKeys
	var userKeys []ruleKeys
	var usersKeysReturnMessage usersKeysStruct

	usersKeysReturnMessage = usersKeysStruct{
		userId:     usersRules.userName,
		userSecret: usersRules.usersSecret,
	}

	var allRulesWithKeysAsMap map[string]ruleKeys
	allRulesWithKeysAsMap = make(map[string]ruleKeys)

	// Loop all rules
	for _, ruleName := range usersRules.allRules {
		// Create keypairs for the rule
		pubKey, priKey, _ := ed25519.GenerateKey(nil)

		// Create keys for rule and append to list of rules
		rulesKeys := ruleKeys{
			functionName: ruleName,
			publicKey:    pubKey,
			privateKey:   priKey,
		}
		allRulesWithKeys = append(allRulesWithKeys, rulesKeys)
		allRulesWithKeysAsMap[ruleName] = rulesKeys

		usersKeysReturnMessage.allPublicKeys = append(usersKeysReturnMessage.allPublicKeys, pubKey)

	}

	// Loop users rules and copy values from total list
	for _, ruleName := range usersRules.allowedFRules {

		// Copy keys that for rules that users has access too
		userKeys = append(userKeys, allRulesWithKeysAsMap[ruleName])

		usersKeysReturnMessage.allowedRules = append(usersKeysReturnMessage.allowedRules, ruleName)
		usersKeysReturnMessage.privateKeys = append(usersKeysReturnMessage.privateKeys, allRulesWithKeysAsMap[ruleName].privateKey)
		usersKeysReturnMessage.publicKeys = append(usersKeysReturnMessage.publicKeys, allRulesWithKeysAsMap[ruleName].publicKey)
	}

	// Create an aggregated signature for all rules that users can use.[Change to only use users own keys as all possible keys]
	// usersKeysReturnMessage.aggregatedSignatureStoredInA3S = SignMulti(usersKeysReturnMessage.userSecret, usersKeysReturnMessage.publicKeys, usersKeysReturnMessage.privateKeys, usersKeysReturnMessage.allPublicKeys)
	usersKeysReturnMessage.aggregatedSignatureStoredInA3S = SignMulti(usersKeysReturnMessage.userSecret, usersKeysReturnMessage.publicKeys, usersKeysReturnMessage.privateKeys, usersKeysReturnMessage.allPublicKeys)

	// Store aggregated signature in A3S simulator Engine
	updateUserAggregatedSignatureRequest := a3s_grpc_api.UpdateUserAggregatedSignatureRequest{
		UserId:                  usersRules.userName,
		UserAggregatedSignature: string(usersKeysReturnMessage.aggregatedSignatureStoredInA3S),
	}

	updateUserAggregatedSignatureResponse := userAuthorizationEngineServerObject.updateUserAggregatedSignature(&updateUserAggregatedSignatureRequest)

	if updateUserAggregatedSignatureResponse.Acknack == false {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"id":                                    "b26e0c90-d1d0-4a50-b8eb-347ed9e5174a",
			"updateUserAggregatedSignatureResponse": updateUserAggregatedSignatureResponse,
		}).Debug("Couldn't update A3S simulator Engine with 'Aggregated signature'")
	}

	// Create return message
	returnMessage = returnMessageStruct{
		userId:        usersRules.userName,
		publicKeys:    usersKeysReturnMessage.publicKeys,
		allPublicKeys: usersKeysReturnMessage.allPublicKeys,
	}

	return returnMessage
}

/***************************************************************/
// Verify that signature match userSecret
type usersKeysAndUserSceretUsedInValidatingStruct struct {
	userId        string
	userSecret    [32]byte
	allowedRules  []string
	publicKeys    []ed25519.PublicKey
	allPublicKeys []ed25519.PublicKey
}

func verifuIfUserSignatureMatchSecret(usersKeysAndUserSceretUsedInValidating usersKeysAndUserSceretUsedInValidatingStruct) (userSignatureMatchSecret bool) {

	// Get aggregated Signature from A3S simulator Engine
	getUserAggregatedSignatureRequest := a3s_grpc_api.GetUserAggregatedSignatureRequest{
		UserId: usersKeysAndUserSceretUsedInValidating.userId,
	}
	getUserAggregatedSignatureResponse := userAuthorizationEngineServerObject.getUserAggregatedSignature(&getUserAggregatedSignatureRequest)
	aggregatedSignatureStoredInA3S := []byte(getUserAggregatedSignatureResponse.UserAggregatedSignature)

	// Verify if users aggregated signatur math user secret
	userSignatureMatchSecret = cosi.Verify(usersKeysAndUserSceretUsedInValidating.allPublicKeys, cosi.ThresholdPolicy(len(usersKeysAndUserSceretUsedInValidating.publicKeys)), usersKeysAndUserSceretUsedInValidating.userSecret[:], aggregatedSignatureStoredInA3S)

	if userSignatureMatchSecret == true {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"id":                       "dd88f7c4-40bd-4335-887f-bad9cfeea45d",
			"usersKeys":                usersKeysAndUserSceretUsedInValidating.publicKeys,
			"userSignatureMatchSecret": userSignatureMatchSecret,
		}).Debug("Users aggregated signature matches the secret")

	} else {
		userAuthorizationEngineServerObject.logger.WithFields(logrus.Fields{
			"id":                       "5495e5aa-488a-4a6e-9c14-a517ece9dbf8",
			"usersKeys":                usersKeysAndUserSceretUsedInValidating.publicKeys,
			"userSignatureMatchSecret": userSignatureMatchSecret,
		}).Debug("Users aggregated signature does NOT match the secret")

	}

	return userSignatureMatchSecret
}

// This example demonstrates how to generate a
// collective signature involving two cosigners,
// and how to check the resulting collective signature.
func main() {

	type userType struct {
		userId                         string
		userSecret                     [32]byte
		allowedFunctions               []string
		publicKeys                     []ed25519.PublicKey
		privateKeys                    []ed25519.PrivateKey
		aggregatedSignatureStoredInA3S []byte
	}

	type functionType struct {
		functionName string
		publicKey    ed25519.PublicKey
		privateKey   ed25519.PrivateKey
	}

	var myFunctionList []functionType

	// Create keypairs for the "functions" where system should validate authorization  .
	pubKey1, priKey1, _ := ed25519.GenerateKey(nil)
	pubKey2, priKey2, _ := ed25519.GenerateKey(nil)
	pubKey3, priKey3, _ := ed25519.GenerateKey(nil)
	pubKey4, priKey4, _ := ed25519.GenerateKey(nil)
	pubKey5, priKey5, _ := ed25519.GenerateKey(nil)
	pubKey6, priKey6, _ := ed25519.GenerateKey(nil)
	pubKey7, priKey7, _ := ed25519.GenerateKey(nil)
	pubKey8, priKey8, _ := ed25519.GenerateKey(nil)
	pubKey9, priKey9, _ := ed25519.GenerateKey(nil)

	// All functions with their public- and private keys
	function1 := functionType{
		functionName: "Function 1",
		publicKey:    pubKey1,
		privateKey:   priKey1,
	}
	function2 := functionType{
		functionName: "Function 2",
		publicKey:    pubKey2,
		privateKey:   priKey2,
	}
	function3 := functionType{
		functionName: "Function 3",
		publicKey:    pubKey3,
		privateKey:   priKey3,
	}
	function4 := functionType{
		functionName: "Function 4",
		publicKey:    pubKey4,
		privateKey:   priKey4,
	}
	function5 := functionType{
		functionName: "Function 5",
		publicKey:    pubKey5,
		privateKey:   priKey5,
	}
	function6 := functionType{
		functionName: "Function 6",
		publicKey:    pubKey6,
		privateKey:   priKey6,
	}
	function7 := functionType{
		functionName: "Function 7",
		publicKey:    pubKey7,
		privateKey:   priKey7,
	}
	function8 := functionType{
		functionName: "Function 8",
		publicKey:    pubKey8,
		privateKey:   priKey8,
	}
	function9 := functionType{
		functionName: "Function 9",
		publicKey:    pubKey9,
		privateKey:   priKey9,
	}
	myFunctionList = append(myFunctionList, function1, function2, function3, function4,
		function5, function6, function7, function8, function9)

	// Create users and their list of functions where the userType has correct authorization to access
	user1 := &userType{
		userId:           "User 1",
		userSecret:       sha256.Sum256([]byte("User 1")),
		allowedFunctions: []string{"Function 1", "Function 2", "Function 3"},
		publicKeys:       []ed25519.PublicKey{pubKey1, pubKey2, pubKey3},
		privateKeys:      []ed25519.PrivateKey{priKey1, priKey2, priKey3},
	}

	user2 := &userType{
		userId:           "User 2",
		userSecret:       sha256.Sum256([]byte("User 2")),
		allowedFunctions: []string{"Function 1", "Function 2", "Function 9"},
		publicKeys:       []ed25519.PublicKey{pubKey1, pubKey2, pubKey3, pubKey9},
		privateKeys:      []ed25519.PrivateKey{priKey1, priKey2, priKey3, priKey9},
	}

	user3 := &userType{
		userId:           "User 3",
		userSecret:       sha256.Sum256([]byte("User 3")),
		allowedFunctions: []string{"Function 5"},
		publicKeys:       []ed25519.PublicKey{pubKey5},
		privateKeys:      []ed25519.PrivateKey{priKey5},
	}

	allPublicKeys := []ed25519.PublicKey{pubKey1, pubKey2, pubKey3, pubKey4, pubKey5, pubKey6, pubKey7, pubKey8, pubKey9}

	/*	pubKeys_user_2 :=
		privKey_user_1 := []ed25519.PrivateKey{priKey1, priKey2, priKey9}

		pubKeys_user_3 := []ed25519.PublicKey{pubKey5}
		privKey_user_1 := []ed25519.PrivateKey{priKey1, priKey2, priKey9}
	*/
	// Create an aggregated signature for all functions that users can access.
	user1.aggregatedSignatureStoredInA3S = SignMulti(user1.userSecret, user1.publicKeys, user1.privateKeys, allPublicKeys)
	user2.aggregatedSignatureStoredInA3S = SignMulti(user2.userSecret, user2.publicKeys, user2.privateKeys, allPublicKeys)
	user3.aggregatedSignatureStoredInA3S = SignMulti(user3.userSecret, user3.publicKeys, user3.privateKeys, allPublicKeys)

	fmt.Print(cosi.Verify(allPublicKeys, cosi.ThresholdPolicy(len(user1.publicKeys)), user1.userSecret[:], user1.aggregatedSignatureStoredInA3S))
	fmt.Print(cosi.Verify(allPublicKeys, cosi.ThresholdPolicy(len(user2.publicKeys)), user2.userSecret[:], user2.aggregatedSignatureStoredInA3S))
	fmt.Print(cosi.Verify(allPublicKeys, cosi.ThresholdPolicy(len(user3.publicKeys)), user3.userSecret[:], user3.aggregatedSignatureStoredInA3S))

	//fmt.Println(IsUserAuthorizedToUseFunction(ruleStruct(function1), user1.publicKeys, user1.userSecret, user1.aggregatedSignatureStoredInA3S))
	/*
		id_customer_1 := []byte("Costomer 1")
		id_customer_2 := []byte("Costomer 2")
		id_customer_3 := []byte("Costomer 3")



		sig_customer_1 := Sign(id_customer_1, pubKeys_user_1, priKey1, priKey2)
		//sig_customer_2_Save := Sign3(customerId_2, pubKeys_user_2, priKey1, priKey2, priKey3)
		//sig_customer_2b_Save := Sign(customerId_2, pubKeys_customer2b, priKey1, priKey2)

		// Now verify the resulting collective signature.
		// This can be done by anyone any time, not just the leader.
		valid := cosi.Verify(pubKeys_user_1, nil, id_customer_1, sig_customer_1)
		fmt.Println("signature valid 1: %v", valid)

		valid2 := cosi.Verify(pubKeys_user_2, cosi.ThresholdPolicy(2), id_customer_1, sig_customer_1)
		fmt.Println("signature valid 2: %v", valid2)

		valid1b := cosi.Verify(pubKeys_customer1b, cosi.ThresholdPolicy(2), id_customer_1, sig_customer_1)
		fmt.Println("signature valid 1: %v", valid1b)
	*/
}

// Verfy if a userType should have access to a certain function
func isUserAuthorizedToUseFunction(functionToAuthorize ruleStruct, usersPublicKeys []ed25519.PublicKey, userSecret [32]byte, aggregatedSignatureStoredInA3S []byte) bool {

	for pubKeyPosistion, pubKey := range usersPublicKeys {
		// If public key exist in list then remove it from the list
		if string(pubKey) == string(functionToAuthorize.publicKey) {
			copy(usersPublicKeys[pubKeyPosistion:], usersPublicKeys[pubKeyPosistion+1:])
			usersPublicKeys = usersPublicKeys[:len(usersPublicKeys)-1]
			break
		}
	}

	valid := cosi.Verify(usersPublicKeys, nil, userSecret[:], aggregatedSignatureStoredInA3S)

	return valid
}

// Helper function to implement a bare-bones cosigning process.
// In practice the two cosigners would be on different machines
// ideally managed by independent administrators or key-holders.
func Sign(message []byte, pubKeys []ed25519.PublicKey,
	priKey1, priKey2 ed25519.PrivateKey) []byte {

	// Each cosigner first needs to produce a per-message commit.
	commit1, secret1, _ := cosi.Commit(nil)
	commit2, secret2, _ := cosi.Commit(nil)
	commits := []cosi.Commitment{commit1, commit2}

	// The leader then combines these into an aggregate commit.
	cosigners := cosi.NewCosigners(pubKeys, nil)
	aggregatePublicKey := cosigners.AggregatePublicKey()
	aggregateCommit := cosigners.AggregateCommit(commits)

	// The cosigners now produce their parts of the collective signature.
	sigPart1 := cosi.Cosign(priKey1, secret1, message, aggregatePublicKey, aggregateCommit)
	sigPart2 := cosi.Cosign(priKey2, secret2, message, aggregatePublicKey, aggregateCommit)
	sigParts := []cosi.SignaturePart{sigPart1, sigPart2}

	// Finally, the leader combines the two signature parts
	// into a final collective signature.
	sig := cosigners.AggregateSignature(aggregateCommit, sigParts)

	return sig
}

/**************************************/
// Aggregate signatures from each Rule connected to each user
func SignMulti(userMessage [32]byte, userPubKeys []ed25519.PublicKey, userPriKeys []ed25519.PrivateKey, allPublicKeys []ed25519.PublicKey) []byte {

	var sigPart cosi.SignaturePart
	var sigParts []cosi.SignaturePart
	var keyFound bool
	var keyPositionInUserKeys int
	var keyPositionInAllKeys int

	// Generate Bitmask, and decimal representation, for Rules that did sign the Secret Message
	bitArray, commits, secrets := generateBitArrayCommitsAndSecrets(userPubKeys, allPublicKeys)
	bitMaskAsDecimal := binaryAsStringToDecimal(bitArray)
	bitMaskAsByteArray := integerToByteArray(bitMaskAsDecimal, uint64(len(bitArray)))

	//Combine Rule-commits into an aggregate Rule-commit.
	cosigners := cosi.NewCosigners(allPublicKeys, bitMaskAsByteArray)
	aggregatePublicKey := cosigners.AggregatePublicKey()
	aggregateCommit := cosigners.AggregateCommit(commits)

	// Set the length of signature- array
	sigParts = make([]cosi.SignaturePart, len(allPublicKeys))

	// Per cosigner produce its part of the collective signature.
	// Loop all public keys to find correct position for userKey among allKeys
	for allKeyPosition, pubKey := range allPublicKeys {
		keyFound = false

		// For each public key do check it exists in among users public keys
		for userKeyPosition, userPubKey := range userPubKeys {
			if reflect.DeepEqual(userPubKey, pubKey) {
				keyFound = true
				keyPositionInAllKeys = allKeyPosition
				keyPositionInUserKeys = userKeyPosition

				break
			}
		}

		// For found keys generate each signature
		if keyFound == true {
			sigPart = cosi.Cosign(userPriKeys[keyPositionInUserKeys], secrets[keyPositionInUserKeys], userMessage[:], aggregatePublicKey, aggregateCommit)
			sigParts[keyPositionInAllKeys] = sigPart
		}
	}

	// Combines the signature parts into a final collective signature.
	aggregatedSignature := cosigners.AggregateSignature(aggregateCommit, sigParts)

	return aggregatedSignature
}

/**************************************/
// Create an binary array
// First Key is placed to the right and last key is placed to the left in the binary array
func generateBitArrayCommitsAndSecrets(userPubKeys []ed25519.PublicKey, allPublicKeys []ed25519.PublicKey) (string, []cosi.Commitment, []*cosi.Secret) {
	var binaryStringArray string = ""
	var keyFound bool
	var commit cosi.Commitment
	var secret *cosi.Secret
	var commits []cosi.Commitment
	var secrets []*cosi.Secret
	var keyPositionInAllKeys int
	var keyPositionInUserKeys int

	// Set the length of commit- and secret- arrays
	commits = make([]cosi.Commitment, len(allPublicKeys))
	secrets = make([]*cosi.Secret, len(userPubKeys))

	// Loop all public keys to generate bitmask of used Rules together with commits and secrets
	for allKeyPosition, pubKey := range allPublicKeys {
		keyFound = false

		// For each public key do check it exists in among users public keys
		for userKeyPosition, userPubKey := range userPubKeys {
			if reflect.DeepEqual(userPubKey, pubKey) {
				keyFound = true
				keyPositionInAllKeys = allKeyPosition
				keyPositionInUserKeys = userKeyPosition

				break
			}
		}

		// Keys that are found gets an "0" and not found keys gets a "1"
		// For found keys generate each rules specific commit and secret
		if keyFound == true {
			binaryStringArray = "0" + binaryStringArray

			commit, secret, _ = cosi.Commit(nil)
			commits[keyPositionInAllKeys] = commit
			secrets[keyPositionInUserKeys] = secret
			// For
		} else {
			binaryStringArray = "1" + binaryStringArray
		}

	}

	return binaryStringArray, commits, secrets
}

/**************************************/
// BinaryToDecimal converts a binary number string to a decimal number.
// It does not support negative binary numbers.
func binaryAsStringToDecimal(binaryArray string) uint64 {
	if convertedInteger, err := strconv.ParseInt(binaryArray, 2, 64); err != nil {
		println(err)
		return 0
	} else {
		return uint64(convertedInteger)
	}
}

func integerToByteArray(integerToConvert uint64, numberofBitsInBitmask uint64) []byte {

	var numberOfBytesNeededForBitmask uint64

	// Calculate number of bytes need for return value
	numberOfBytesNeededForBitmask = uint64(math.Ceil((float64(numberofBitsInBitmask)) / 8))

	bitmaskByteArray64 := make([]byte, 8)
	binary.LittleEndian.PutUint64(bitmaskByteArray64, integerToConvert)

	// Determine how many bytes that is needed for Schnorr signature Bitmask
	bitmaskByteArraayToRetrun := bitmaskByteArray64[0:numberOfBytesNeededForBitmask]

	// How to convert back into int64: x1 := binary.LittleEndian.Uint64(cn[0:8])
	//x1 := binary.LittleEndian.Uint64(bitmaskByteArray64[0:8])
	//fmt.Println(x1)

	return bitmaskByteArraayToRetrun

}
