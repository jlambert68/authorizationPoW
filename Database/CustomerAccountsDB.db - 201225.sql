BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "CompanyData" (
	"Name"	TEXT NOT NULL UNIQUE,
	"Address"	TEXT NOT NULL,
	"Country"	TEXT NOT NULL,
	"ContactAddress"	TEXT NOT NULL,
	PRIMARY KEY("Name")
);
CREATE TABLE IF NOT EXISTS "AccountTypes" (
	"Id"	INTEGER NOT NULL UNIQUE,
	"Type"	TEXT NOT NULL UNIQUE,
	"TypeDescription"	TEXT NOT NULL,
	PRIMARY KEY("Id" AUTOINCREMENT)
);
CREATE TABLE IF NOT EXISTS "Accounts" (
	"AccountNumber"	TEXT NOT NULL UNIQUE,
	"Description"	TEXT NOT NULL,
	"Amount"	NUMERIC NOT NULL,
	"Type"	INTEGER NOT NULL,
	FOREIGN KEY("Type") REFERENCES "AccountTypes"("Id"),
	PRIMARY KEY("AccountNumber")
);
INSERT INTO "CompanyData" ("Name","Address","Country","ContactAddress") VALUES ('SEB AB','MAll of Scandinaviagången 11','Sweden','info@seb.se'),
 ('Drakon Sverige AB','Telegramvägen 57','Sweden','info@drakon.se'),
 ('Scania AB','Södertäljevägen 4','Sweden','info@scania.se');
INSERT INTO "AccountTypes" ("Id","Type","TypeDescription") VALUES (1,'MOMS - 25%	','Moms 25%'),
 (2,'MOMS  -12%','Moms 12%'),
 (3,'Ränteintäkter','Ränteintäkter'),
 (4,'Försäljing','Försäljning'),
 (5,'Konstnader','Kostnader'),
 (6,'Deleteable','Can be deleted, and used for demo purpose');
INSERT INTO "Accounts" ("AccountNumber","Description","Amount","Type") VALUES ('283592388-31','Moms',23235.37,1),
 ('12412412-31','FÖrsäljning',1223.44,1);
COMMIT;
