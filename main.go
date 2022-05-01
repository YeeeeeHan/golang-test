package main

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/service"
	"TechnicalAssignment/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

//db.GlobalBalanceTable.Put("def", 456)
//var bal int
//_ = db.GlobalBalanceTable.Get("def", &bal)
//fmt.Println(bal)

func main() {

	// Init Tables
	var initerr error
	db.GlobalBalanceTable, db.GlobalPasswordTable, initerr = db.InitTables()
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

	for {
		fmt.Print(fmt.Sprintf("\n(%s) > ", sessionUser.GetUsername()))
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command, args := utils.ParseInput(text)

		// Handle Register command
		if command == constants.Register {
			err := service.Register(&sessionUser, args)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.RegisterSuccess)
			}
			continue
		}

		// If user is not logged in, they cannot initiate any commands other than 'Login'
		if sessionUser == (service.Wallet{}) && command != constants.Login {
			fmt.Println(constants.NotLoggedInMsg)
			continue
		}

		// Process commands
		var err error
		switch command {
		case constants.Login:
			err = service.Login(&sessionUser, args)
		case constants.Deposit:
			err = service.Deposit(&sessionUser, args)
		case constants.Withdraw:
			err = service.Withdraw(&sessionUser, args)
		case constants.Send:
			err = service.Send(&sessionUser, args)
		case constants.Balance:
			err = service.Balance(&sessionUser)
		case constants.Logout:
			err = service.Logout(&sessionUser)
		case constants.Accounts:
			err = service.Accounts(&sessionUser)
		default:
			fmt.Println(constants.InvalidCommandMsg)
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("ERROR: %s!", err))
		}
	}
}
