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

		command, _ := utils.ParseInput(text)

		if command != constants.Login && user == (service.User{}) {
			fmt.Println(constants.NotLoggedInMsg)
			continue
		}

		switch command {
		case constants.Login:
			fmt.Println("Prcoessing 'Login'...")
		case constants.Withdraw:
			fmt.Println("Prcoessing 'Withdraw'...")
		case constants.Send:
			fmt.Println("Prcoessing 'Send'...")
		case constants.Balance:
			fmt.Println("Prcoessing 'Balance'...")
		case constants.Logout:
			fmt.Println("Prcoessing 'Logout'...")
		case constants.Accounts:
			fmt.Println("Prcoessing 'Accounts'...")
		default:
			fmt.Println(constants.InvalidCommandMsg)
		}
	}
}
