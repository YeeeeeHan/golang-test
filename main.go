package main

import (
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/service"
	"TechnicalAssignment/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {

	user := service.User{}

	fmt.Println(
		"Please Enter a command: " +
			"\n login <username> - login" +
			"\n deposit <x> - deposit $x to your account" +
			"\n withdraw <x> - withdraw $x from your account" +
			"\n send <username> <x> - sends $x to <username>" +
			"\n balance - shows your acccount balance" +
			"\n logout - logout" +
			"\n accounts - view all account information (admin only)")
	for {
		fmt.Print("\n>")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command, args := utils.ParseInput(text)

		if command != constants.Login && user == (service.User{}) {
			fmt.Println(constants.NotLoggedInMsg)
			continue
		}

		switch command {
		case constants.Login:
			_ = service.Login(&user, args)
		case constants.Deposit:
			_ = service.Deposit(&user, args)
		case constants.Withdraw:
			_ = service.Withdraw(&user, args)
		case constants.Send:
			_ = service.Send(&user, args)
		case constants.Balance:
			_ = service.Balance(&user)
		case constants.Logout:
			_ = service.Logout(&user)
		case constants.Accounts:
			_ = service.Accounts(&user)
		default:
			fmt.Println(constants.InvalidCommandMsg)
		}
	}
}
