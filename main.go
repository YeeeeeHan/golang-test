package main

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/cmd/server"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/service"
	"fmt"
)

//db.GlobalBalanceTable.Put("def", 456)
//var bal int
//_ = db.GlobalBalanceTable.Get("def", &bal)
//fmt.Println(bal)

func main() {

	// Init Tables
	db.GlobalUsernameTable = constants.UsernameFile
	var initerr error
	db.GlobalBalanceTable, db.GlobalPasswordTable, initerr = db.InitTables(constants.BalanceFile, constants.PasswordFile)
	if initerr != nil {
		panic("DB init error")
	}

	// Close DB
	defer db.CloseDB(db.GlobalBalanceTable)
	defer db.CloseDB(db.GlobalPasswordTable)

	// Init Wallet
	sessionUser := service.Wallet{}

	// Register Admin
	db.CreateUser("admin", "password123")

	fmt.Println(
		"Please Enter a command: " +
			"\n register <username> <password> - adds an account with username and password" +
			"\n login <username> <password>- login with username and password" +
			"\n deposit <x> - deposit $x to your account" +
			"\n withdraw <x> - withdraw $x from your account" +
			"\n send <username> <x> - sends $x to <username>" +
			"\n balance - shows your acccount balance" +
			"\n logout - logout" +
			"\n accounts - view all account information (admin only)")

	server.ListenAndServe(&sessionUser)
}
