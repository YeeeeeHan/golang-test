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

	// Init Balance Table
	var initerr error
	db.GlobalBalanceTable, initerr = db.InitBalanceTable()
	if initerr != nil {
		panic("balance DB init error")
	}

	// Init Password Table
	db.GlobalPasswordTable, initerr = db.InitPasswordTable()
	if initerr != nil {
		panic("password DB init error")
	}

	// Close DB
	defer db.CloseDB(db.GlobalBalanceTable)
	defer db.CloseDB(db.GlobalPasswordTable)

	// Init User
	user := service.User{}

	// Register Admin
	db.GlobalPasswordTable.Put("admin", "password123")

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
		fmt.Print(fmt.Sprintf("\nUser: %s > ", user.GetUsername()))
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command, args := utils.ParseInput(text)

		// Handle Register command
		if command == constants.Register {
			err := service.Register(&user, args)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(constants.RegisterSuccess)
			}
			continue
		}

		// If user is not logged in, they cannot initiate any commands other than 'Login'
		if user == (service.User{}) && command != constants.Login {
			fmt.Println(constants.NotLoggedInMsg)
			continue
		}

		// Process commands
		var err error
		switch command {
		case constants.Login:
			err = service.Login(&user, args)
		case constants.Deposit:
			err = service.Deposit(&user, args)
		case constants.Withdraw:
			err = service.Withdraw(&user, args)
		case constants.Send:
			err = service.Send(&user, args)
		case constants.Balance:
			err = service.Balance(&user)
		case constants.Logout:
			err = service.Logout(&user)
		case constants.Accounts:
			err = service.Accounts(&user)
		default:
			fmt.Println(constants.InvalidCommandMsg)
		}
		if err != nil {
			fmt.Println(err)
		}
	}
}
