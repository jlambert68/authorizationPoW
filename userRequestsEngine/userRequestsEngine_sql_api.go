package userRequestsEngine

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
	"jlambert/authorizationPoW/grpc_api/userRequests_grpc_api"
	"os"
)

/****************************************************/
// Execute user request: 'ListAccounts (ListAccountsRequest) returns (ListAccountsResponse)'
func (userRequestsServerObject *userRequestsServerObjectStruct) sqlListAccounts(listAccountsRequest *userRequests_grpc_api.ListAccountsRequest, allowedUsersAccounts userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse) *userRequests_grpc_api.ListAccountsResponse {
	var err error
	var accountsList string = ""
	var returnMessage *userRequests_grpc_api.ListAccountsResponse

	// If user doesn't has access to any accounts then exit with empty result set
	if len(allowedUsersAccounts.Accounts) == 0 {
		returnMessage = &userRequests_grpc_api.ListAccountsResponse{
			UserId:    listAccountsRequest.UserId,
			CompanyId: listAccountsRequest.CompanyId,
			Acknack:   false,
			Comments:  "User doesn't have access to any accounts.",
			Accounts:  nil,
		}

		// return message to user
		return returnMessage

	} else {
		// Convert User Allowed Accounts into a string
		for accountPosition, account := range allowedUsersAccounts.GetAccounts() {
			if accountPosition == 0 {
				accountsList = "'" + account.Account + "'"
			} else {
				accountsList = accountsList + "," + "'" + account.Account + "'"
			}
		}
	}

	// SQl for 'ListAccounts'
	sqlText := "SELECT AccountNumber "
	sqlText += "FROM Accounts "
	sqlText += "WHERE "
	sqlText += "Company = " + "'" + listAccountsRequest.CompanyId + "' AND "
	sqlText += "AccountNumber In (" + accountsList + ") "
	sqlText += "ORDER BY AccountNumber "

	// Execute a sql quesry
	sqlResponseRows, err := userRequestsServerObject.sqlDbObject.Query(sqlText)
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":          "236e8545-920e-4ef9-a815-0d39587cc114",
			"err.Error()": err.Error(),
			"sqlText":     sqlText,
		}).Warning("Couldn't execute sql-query")

		// Create return message
		returnMessage = &userRequests_grpc_api.ListAccountsResponse{
			UserId:    listAccountsRequest.UserId,
			CompanyId: listAccountsRequest.CompanyId,
			Acknack:   false,
			Comments:  "Error While executing SQL",
			Accounts:  nil,
		}
		return returnMessage

	} else {

		// Success in executing sqlStatement
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":              "146e1294-141a-4584-9cc3-cd0e48421b5b",
			"sqlResponseRows": sqlResponseRows,
		}).Debug("Success in executing sql for 'ListAccounts'")

		// Extract data from SQL results and create response object
		var accountsList []*userRequests_grpc_api.Account
		var AccountNumber string

		// Iterate and fetch the records from result cursor
		for sqlResponseRows.Next() {
			sqlResponseRows.Scan(&AccountNumber)
			convertedAccount := &userRequests_grpc_api.Account{Account: AccountNumber}
			accountsList = append(accountsList, convertedAccount)
		}

		// Create return message
		returnMessage = &userRequests_grpc_api.ListAccountsResponse{
			UserId:    listAccountsRequest.UserId,
			CompanyId: listAccountsRequest.CompanyId,
			Acknack:   true,
			Comments:  "",
			Accounts:  accountsList,
		}
	}

	return returnMessage
}

/****************************************************/
// Execute user request: 'ListAccountsBaseOnProvidedType (ListAccountsBasedOnProvidedTypeRequest) returns (ListAccountsBasedOnProvidedTypeResponse)'

/****************************************************/
// Execute user request: 'AddAccount (AddAccountRequest) returns (AddAccountResponse)'

/****************************************************/
// Execute user request: 'DeleteAccount (DeleteAccountRequest) returns (DeleteAccountResponse)'

/****************************************************/
// Execute user request: 'AddAccountType (AddAccountTypeRequest) returns (AddAccountTypeResponse)'

/****************************************************/
// Execute user request: 'DeleteAccountType (DeleteAccountTypeTypeRequest) returns (DeleteAccountTypeResponse)'

/****************************************************/
// Execute user request: 'UpdateCompanyInformation (UpdateCompanyInformationRequest) returns (UpdateCompanyInformationResponse)'

/****************************************************/
// Initiate database. If already exists then use it otherwise create a new one and fill with standardized data
func (userRequestsServerObject *userRequestsServerObjectStruct) initializeSqlDB() {
	var err error

	// Open connection towards database
	userRequestsServerObject.sqlDbObject, err = sql.Open("sqlite3", userRequestsServerObject.databaseName)

	// If database was not found then create and initiate a database from scratch
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id": "0a58dc65-e1e1-4127-8334-8e55376c6320",
		}).Info("Couldn't open existing database, will create a new one and fill with standardized data.")

		// Create the new database with data included.
		// That part will open the database
		userRequestsServerObject.createNewDatabase()

	}

}

/****************************************************/
// Create a new database and fill with standardized data
func (userRequestsServerObject *userRequestsServerObjectStruct) createNewDatabase() {
	var err error

	// Create the database-file
	userRequestsServerObject.logger.WithFields(logrus.Fields{
		"Id":                                    "77ca5d18-b9b6-4ef7-978c-f08bda35f86d",
		"userRequestsServerObject.databaseName": userRequestsServerObject.databaseName,
	}).Info("Creating a new database")

	databaseFile, err := os.Create(userRequestsServerObject.databaseName)

	// If not succeeded then exit program because something is not as intended.
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":          "15ab56b9-48c1-4b0b-8b83-2fd7d469d67e",
			"err.Error()": err.Error(),
		}).Panic("Exiting because couldn't create a new database file")
	} else {

		//Success in creating database
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":                                    "77ca5d18-b9b6-4ef7-978c-f08bda35f86d",
			"userRequestsServerObject.databaseName": userRequestsServerObject.databaseName,
		}).Debug("Success in creating a new database")

		// Close file
		err = databaseFile.Close()
		// If not succeeded then exit program because something is not as intended.
		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"Id":          "8e627225-1478-4fd8-ae14-56df59653863",
				"err.Error()": err.Error(),
			}).Panic("Exiting because couldn't close the new database file")
		} else {

			// Fill newly create database with standardized data
			userRequestsServerObject.FillDatabaseWithStandardizedData()
		}
	}

}

/****************************************************/
// Fill database with standardized data
func (userRequestsServerObject *userRequestsServerObjectStruct) FillDatabaseWithStandardizedData() {
	var err error

	// Retry to open connection towards newly database
	userRequestsServerObject.sqlDbObject, err = sql.Open("sqlite3", userRequestsServerObject.databaseName)

	// If database was not found then create exit program due to something is wrong
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":          "ef63447f-00b0-414b-bb46-1b416d30270d",
			"err.Error()": err.Error(),
		}).Panic("Couldn't open newly create database, will exit program.")

	}

	// Load SQL from file
	fileContent, err := ioutil.ReadFile(userRequestsServerObject.sqlFile)
	// IfSQL-file couldn't be read then exit program
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":          "63a3e2df-a8ba-488b-8c74-f2b2e0b96e3d",
			"err.Error()": err.Error(),
		}).Panic("Couldn't read SQL-file, will exit program.")
	}

	// Success in reading SQL-file. Convert []byte to string
	sqlText := string(fileContent)

	// Create a sql statement
	sqlStatement, err := userRequestsServerObject.sqlDbObject.Prepare(sqlText) // Prepare SQL Statement
	if err != nil {
		userRequestsServerObject.logger.WithFields(logrus.Fields{
			"Id":          "89580262-2144-48f7-ad9e-3937bc3b4c5d",
			"err.Error()": err.Error(),
		}).Panic("Couldn't crete sql-statement, will exit program.")
	} else {

		// Execute SQL Statements
		sqlResults, err := sqlStatement.Exec()
		// If not succeeded then exit program because something is not as intended.
		if err != nil {
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"Id":                                    "334f7ce8-a610-4667-979f-40153d3d8fec",
				"err.Error()":                           err.Error(),
				"userRequestsServerObject.databaseName": userRequestsServerObject.databaseName,
			}).Panic("Exiting because couldn't execute sql to create and initialize database")
		} else {
			// Success in executing sqlStatement
			userRequestsServerObject.logger.WithFields(logrus.Fields{
				"Id":                                    "85efe0ff-3645-47fa-a014-a2747f176381",
				"userRequestsServerObject.databaseName": userRequestsServerObject.databaseName,
				"sqlResults":                            sqlResults,
			}).Debug("Success in executing sql to create and initialize database")
		}
	}

}
