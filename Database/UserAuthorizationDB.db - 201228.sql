BEGIN TRANSACTION;
DROP TABLE IF EXISTS "AuthorizedAccounts";
CREATE TABLE IF NOT EXISTS "AuthorizedAccounts" (
	"AccountNumber"	TEXT NOT NULL,
	"UserName"	TEXT NOT NULL,
	"Company"	TEXT,
	PRIMARY KEY("AccountNumber","UserName","Company")
);
DROP TABLE IF EXISTS "AuthorizedAccounttypes";
CREATE TABLE IF NOT EXISTS "AuthorizedAccounttypes" (
	"AccountType"	TEXT NOT NULL,
	"UserName"	TEXT NOT NULL,
	"Company"	TEXT,
	PRIMARY KEY("AccountType","UserName","Company")
);
DROP TABLE IF EXISTS "AuthorizedCompanies";
CREATE TABLE IF NOT EXISTS "AuthorizedCompanies" (
	"Company"	TEXT NOT NULL,
	"UserName"	TEXT NOT NULL,
	PRIMARY KEY("Company","UserName")
);
DROP TABLE IF EXISTS "AuthorizedUsers";
CREATE TABLE IF NOT EXISTS "AuthorizedUsers" (
	"UserName"	TEXT NOT NULL UNIQUE,
	PRIMARY KEY("UserName")
);
INSERT INTO "AuthorizedAccounts" ("AccountNumber","UserName","Company") VALUES ('12412412-31','Alice','SEB AB'),
 ('12412412-31','Bob','SEB AB'),
 ('283592388-31','Alice','SEB AB');
INSERT INTO "AuthorizedCompanies" ("Company","UserName") VALUES ('SEB AB','Alice'),
 ('SEB AB','Bob'),
 ('SEB AB','Carol');
INSERT INTO "AuthorizedUsers" ("UserName") VALUES ('Alice'),
 ('Bob'),
 ('Carol'),
 ('Chuck'),
 ('Sybil');
COMMIT;
