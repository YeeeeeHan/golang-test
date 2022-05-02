package main

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/cmd/server"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/service"
	"fmt"
)

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
			"\n 1) register <username> <password> - adds an account with username and password" +
			"\n 2) login <username> <password>- login with username and password" +
			"\n 3) deposit <x> - deposit $x to your account" +
			"\n 4) withdraw <x> - withdraw $x from your account" +
			"\n 5) send <username> <x> - sends $x to <username>" +
			"\n 6) balance - shows your acccount balance" +
			"\n 7) logout - logout" +
			"\n 8) accounts - view all account information (admin only)")

	server.ListenAndServe(&sessionUser)
}
