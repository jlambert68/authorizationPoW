package userAuthorizationEngine

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jlambert/authorizationPoW/grpc_api/userAuthorizationEngine_grpc_api"
	"os"
)

/****************************************************/
// List users authorized accounts
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) sqlListUsersAuthorizedAccounts(userAuthorizedAccountsRequest *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsRequest) *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse {
	var err error
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse

	// SQl for 'List users authorized accounts'
	sqlText := "SELECT AccountNumber "
	sqlText += "FROM AuthorizedAccounts "
	sqlText += "WHERE "
	sqlText += "UserName = '" + userAuthorizedAccountsRequest.UserId + "' AND "
	sqlText += "Company = '" + userAuthorizedAccountsRequest.CompanyId + "' "
	sqlText += "ORDER BY AccountNumber "

	// Execute a sql quesry
	sqlResponseRows, err := userAuthorizationServerObject.sqlDbObject.Query(sqlText)
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "0d391723-b683-42ec-8b13-8672dd5be58b",
			"err.Error()": err.Error(),
			"sqlText":     sqlText,
		}).Warning("Couldn't execute sql-query")

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse{
			UserId:    userAuthorizedAccountsRequest.UserId,
			CompanyId: userAuthorizedAccountsRequest.CompanyId,
			Acknack:   false,
			Comments:  "Error While executing SQL",
			Accounts:  nil,
		}
		return returnMessage

	} else {

		// Success in executing sqlStatement
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":              "89b97829-96e5-468e-b369-dc9867fe6412",
			"sqlResponseRows": sqlResponseRows,
		}).Debug("Success in executing sql for 'List users authorized accounts'")

		// Extract data from SQL results and create response object
		var accountsList []*userAuthorizationEngine_grpc_api.Account
		var AccountNumber string

		// Iterate and fetch the records from result cursor
		for sqlResponseRows.Next() {
			sqlResponseRows.Scan(&AccountNumber)
			convertedAccount := &userAuthorizationEngine_grpc_api.Account{Account: AccountNumber}
			accountsList = append(accountsList, convertedAccount)
		}

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountsResponse{
			UserId:    userAuthorizedAccountsRequest.UserId,
			CompanyId: userAuthorizedAccountsRequest.CompanyId,
			Acknack:   true,
			Comments:  "",
			Accounts:  accountsList,
		}
	}

	return returnMessage
}

/****************************************************/
// List users authorized account types
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) sqlListUsersAuthorizedAccountTypes(userAuthorizedAccountTypesRequest *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesRequest) *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse {
	var err error
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse

	// SQl for 'List users authorized accounts'
	sqlText := "SELECT AccountType "
	sqlText += "FROM AuthorizedAccountTypes "
	sqlText += "WHERE "
	sqlText += "UserName = '" + userAuthorizedAccountTypesRequest.UserId + "' AND "
	sqlText += "Company = '" + userAuthorizedAccountTypesRequest.CompanyId + "' "
	sqlText += "ORDER BY AccountType "

	// Execute a sql quesry
	sqlResponseRows, err := userAuthorizationServerObject.sqlDbObject.Query(sqlText)
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "011b5759-ef35-479a-bd17-a9ffac0baaf8",
			"err.Error()": err.Error(),
			"sqlText":     sqlText,
		}).Warning("Couldn't execute sql-query")

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse{
			UserId:       userAuthorizedAccountTypesRequest.UserId,
			CompanyId:    userAuthorizedAccountTypesRequest.CompanyId,
			Acknack:      false,
			Comments:     "Error While executing SQL",
			AccountTypes: nil,
		}
		return returnMessage

	} else {

		// Success in executing sqlStatement
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":              "5a24c99c-341e-44bd-89dc-42aa2c4072d6",
			"sqlResponseRows": sqlResponseRows,
		}).Debug("Success in executing sql for 'List users authorized account types'")

		// Extract data from SQL results and create response object
		var accountTypesList []*userAuthorizationEngine_grpc_api.AccountType
		var AccountType string

		// Iterate and fetch the records from result cursor
		for sqlResponseRows.Next() {
			sqlResponseRows.Scan(&AccountType)
			convertedAccountType := &userAuthorizationEngine_grpc_api.AccountType{AccountType: AccountType}
			accountTypesList = append(accountTypesList, convertedAccountType)
		}

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedAccountTypesResponse{
			UserId:       userAuthorizedAccountTypesRequest.UserId,
			CompanyId:    userAuthorizedAccountTypesRequest.CompanyId,
			Acknack:      true,
			Comments:     "",
			AccountTypes: accountTypesList,
		}
	}

	return returnMessage
}

/****************************************************/
// List users authorized companies
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) sqlListUsersAuthorizedCompanies(userAuthorizedCompaniesRequest *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesRequest) *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse {
	var err error
	var returnMessage *userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse

	// SQl for 'List users authorized accounts'
	sqlText := "SELECT Company "
	sqlText += "FROM AuthorizedCompany "
	sqlText += "WHERE "
	sqlText += "UserName = '" + userAuthorizedCompaniesRequest.UserId + "' AND "
	sqlText += "ORDER BY Company "

	// Execute a sql quesry
	sqlResponseRows, err := userAuthorizationServerObject.sqlDbObject.Query(sqlText)
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "6c93ed23-02a0-454c-8975-49906677b83c",
			"err.Error()": err.Error(),
			"sqlText":     sqlText,
		}).Warning("Couldn't execute sql-query")

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse{
			UserId:    userAuthorizedCompaniesRequest.UserId,
			Acknack:   false,
			Comments:  "Error While executing SQL",
			Companies: nil,
		}
		return returnMessage

	} else {

		// Success in executing sqlStatement
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":              "0d3417ef-c952-4ffd-aed4-e7bb2fd4066a",
			"sqlResponseRows": sqlResponseRows,
		}).Debug("Success in executing sql for 'List users authorized companies'")

		// Extract data from SQL results and create response object
		var companiesList []*userAuthorizationEngine_grpc_api.Company
		var Company string

		// Iterate and fetch the records from result cursor
		for sqlResponseRows.Next() {
			sqlResponseRows.Scan(&Company)
			convertedCompany := &userAuthorizationEngine_grpc_api.Company{Company: Company}
			companiesList = append(companiesList, convertedCompany)
		}

		// Create return message
		returnMessage = &userAuthorizationEngine_grpc_api.UserAuthorizedCompaniesResponse{
			UserId:    userAuthorizedCompaniesRequest.UserId,
			Acknack:   true,
			Comments:  "",
			Companies: companiesList,
		}
	}

	return returnMessage
}

/****************************************************/
// Initiate database. If already exists then use it otherwise create a new one and fill with standardized data
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) initializeSqlDB() {
	var err error

	// Open connection towards database
	userAuthorizationServerObject.sqlDbObject, err = sql.Open("sqlite3", userAuthorizationServerObject.databaseName)

	// If database was not found then create and initiate a database from scratch
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id": "05f289a3-5804-4951-92e1-d584e4773a0c",
			"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
		}).Info("Couldn't open existing database, will create a new one and fill with standardized data.")

		// Create the new database with data included.
		// That part will open the database
		userAuthorizationServerObject.createNewDatabase()

	} else {
		// SUccess in opening database
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id": "04b244e7-59bf-4dbb-933d-28b350b461fa",
			"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
		}).Info("Success in opening existing database")

	}

}

/****************************************************/
// Create a new database and fill with standardized data
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) createNewDatabase() {
	var err error

	// Create the database-file
	userAuthorizationServerObject.logger.WithFields(logrus.Fields{
		"Id": "4252df43-4f97-4bd7-bf38-7704482606e7",
		"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
	}).Info("Creating a new database")

	databaseFile, err := os.Create(userAuthorizationServerObject.databaseName)

	// If not succeeded then exit program because something is not as intended.
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "860bdc01-4e47-481f-b8e8-bcfce399538c",
			"err.Error()": err.Error(),
		}).Panic("Exiting because couldn't create a new database file")
	} else {

		//Success in creating database
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id": "31a363a8-8c79-4775-a2c2-6cc60f2e872e",
			"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
		}).Debug("Success in creating a new database")

		// Close file
		err = databaseFile.Close()
		// If not succeeded then exit program because something is not as intended.
		if err != nil {
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"Id":          "7b9455ef-f06b-4af2-a875-201c83bb7f30",
				"err.Error()": err.Error(),
			}).Panic("Exiting because couldn't close the new database file")
		} else {

			// Fill newly create database with standardized data
			userAuthorizationServerObject.FillDatabaseWithStandardizedData()
		}
	}

}

/****************************************************/
// Fill database with standardized data
func (userAuthorizationServerObject *userAuthorizationEngineServerObjectStruct) FillDatabaseWithStandardizedData() {
	var err error

	// Retry to open connection towards newly database
	userAuthorizationServerObject.sqlDbObject, err = sql.Open("sqlite3", userAuthorizationServerObject.databaseName)

	// If database was not found then create exit program due to something is wrong
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "c444d387-e1e2-486a-9bd3-422e64f260bc",
			"err.Error()": err.Error(),
		}).Panic("Couldn't open newly create database, will exit program.")

	}

	// Load SQL from file
	fileContent, err := ioutil.ReadFile(userAuthorizationServerObject.sqlFile)
	// IfSQL-file couldn't be read then exit program
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "4b783396-a1e3-4784-a238-67b2dc57156c",
			"err.Error()": err.Error(),
		}).Panic("Couldn't read SQL-file, will exit program.")
	}

	// Success in reading SQL-file. Convert []byte to string
	sqlText := string(fileContent)

	// Create a sql statement
	sqlStatement, err := userAuthorizationServerObject.sqlDbObject.Prepare(sqlText) // Prepare SQL Statement
	if err != nil {
		userAuthorizationServerObject.logger.WithFields(logrus.Fields{
			"Id":          "13bce9c6-ad95-4c25-b5b4-6f240e19cc56d",
			"err.Error()": err.Error(),
		}).Panic("Couldn't crete sql-statement, will exit program.")
	} else {

		// Execute SQL Statements
		sqlResults, err := sqlStatement.Exec()
		// If not succeeded then exit program because something is not as intended.
		if err != nil {
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"Id":          "79308325-8157-475b-aa78-2f45b6c937a7",
				"err.Error()": err.Error(),
				"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
			}).Panic("Exiting because couldn't execute sql to create and initialize database")
		} else {
			// Success in executing sqlStatement
			userAuthorizationServerObject.logger.WithFields(logrus.Fields{
				"Id": "20388de1-1353-45a8-941d-974d1dc432c0",
				"userAuthorizationServerObject.databaseName": userAuthorizationServerObject.databaseName,
				"sqlResults": sqlResults,
			}).Debug("Success in executing sql to create and initialize database")
		}
	}

}
