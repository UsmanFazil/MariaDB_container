package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	create("ENotary")
}

func create(DBname string) {
	Usertable := "create table Users ( userid varchar(40) , email varchar (40) NOT NULL, password varchar(40) NOT NULL, name varchar(40) NOT NULL, company varchar(40) NOT NULL, phone varchar(30) NOT NULL,picture varchar(500) NOT NULL, sign varchar(500) NOT NULL, initials varchar(500) NOT NULL, creationTime varchar(50) NOT NULL,verified bool NOT NULL,UNIQUE(email),PRIMARY KEY (userid) );"
	ContractTable := "create table Contract (ContractID varchar(40), filepath varchar(500) NOT NULL,status varchar(40) NOT NULL, creationTime varchar(40) , Creator varchar(40) NOT NULL, delStatus bool NOT NULL, updateTime varchar(40), contractName varchar(100) NOT NULL,  ExpirationTime varchar(40),Blockchain bool NOT NULL,Message varchar(200),PRIMARY KEY (ContractID), FOREIGN KEY (Creator) REFERENCES Users (userid) );"
	SignerTable := "create table Signer ( ContractID varchar(40),userID varchar (40),name varchar (40) NOT NULL,SignStatus varchar(40) NOT NULL, SignDate varchar(40), DeleteApprove bool NOT NULL,Message varchar(500),Access bool,CC bool,PRIMARY KEY (ContractID,userID), FOREIGN KEY (ContractID) REFERENCES Contract (ContractID), FOREIGN KEY (userID) REFERENCES Users (userid) );"
	WalletTable := "create table WalletInfo ( userid varchar(40), publicAddress varchar (100), PRIMARY KEY (userid), FOREIGN KEY (userid) REFERENCES Users (userid) );"
	CoOrdinatesTable := "create table SignerCoordinates ( ContractID varchar(40), userID varchar(40), name varchar(100), email varchar(100), sign varchar(100), date varchar(100), company varchar(100), text varchar(100), initals varchar(100),checkBox varchar (100),PRIMARY KEY (ContractID,userID), FOREIGN KEY (ContractID) REFERENCES Contract (ContractID), FOREIGN KEY (userID) REFERENCES Users (userid) );"
	ContractFolderTable := "create table ContractFolder ( folderID varchar(40), ContractID varchar(40), PRIMARY KEY (folderID,ContractID), FOREIGN KEY (folderID) REFERENCES Folder (folderID), FOREIGN KEY (ContractID) REFERENCES Contract (contractID) );"
	FolderTable := "create table Folder ( folderID varchar(40), folderName varchar(100) NOT NULL, folderType varchar(400) NOT NULL, userID varchar(40),parentFolder varchar(40),PRIMARY KEY (folderID), FOREIGN KEY (userID) REFERENCES Users (userid) );"
	VerificationTable := "create table Verification (userid varchar(40),VerificationCode varchar(40) NOT NULL,exptime varchar(40) NOT NULL, PRIMARY KEY (userid), FOREIGN KEY (userid) REFERENCES Users (userid));"
	ContactsTable := "Create table contacts ( userid varchar(40),contact varchar(40), PRIMARY KEY(userid,contact), FOREIGN KEY (userid) REFERENCES Users (userid) );"
	TokenBlacklist := "Create table BlackList (token varchar(800),exptime varchar(40) NOT NULL, PRIMARY KEY(token));"

	db, err := sql.Open("mysql", "root:mypass@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}
	_, err = db.Exec("CREATE DATABASE " + DBname)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}
	_, err = db.Exec("USE " + DBname)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}
	stmt1, _ := db.Prepare(Usertable)
	_, _ = stmt1.Exec()

	stmt2, _ := db.Prepare(ContractTable)
	_, _ = stmt2.Exec()

	stmt3, _ := db.Prepare(SignerTable)
	_, _ = stmt3.Exec()

	stmt4, _ := db.Prepare(WalletTable)
	_, _ = stmt4.Exec()

	stmt5, _ := db.Prepare(CoOrdinatesTable)
	_, _ = stmt5.Exec()

	stmt6, _ := db.Prepare(FolderTable)
	_, _ = stmt6.Exec()

	stmt7, _ := db.Prepare(CCTable)
	_, _ = stmt7.Exec()

	stmt8, _ := db.Prepare(ContractFolderTable)
	_, _ = stmt8.Exec()

	stmt9, _ := db.Prepare(VerificationTable)
	_, err = stmt9.Exec()

	stmt10, _ := db.Prepare(ContactsTable)
	_, _ = stmt10.Exec()

	stmt11, _ := db.Prepare(TokenBlacklist)
	_, _ = stmt11.Exec()

	fmt.Println("Schema created successfully..")

	defer db.Close()
}
