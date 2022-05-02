package server

import (
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/service"
	"TechnicalAssignment/pkg/utils"
	"bufio"
	"fmt"
	"os"
)

// ListenAndServe is a blocking function that repeatedly listens for commands from user
func ListenAndServe(sessionUser *service.Wallet) {
	for {
		fmt.Print(fmt.Sprintf("\n(%s) > ", sessionUser.GetUsername()))
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		command, args := utils.ParseInput(text)
		if command == "" {
			fmt.Println(constants.InvalidNumArgumentsMsg)
			continue
		}

		// Handle Register command
		if command == constants.Register {
			err := service.Register(sessionUser, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		// If user is not logged in, they cannot initiate any commands other than 'Login'
		if *sessionUser == (service.Wallet{}) && command != constants.Login {
			fmt.Println(constants.NotLoggedInMsg)
			continue
		}

		// Process commands
		var err error
		switch command {
		case constants.Login:
			err = service.Login(sessionUser, args)
		case constants.Deposit:
			err = service.Deposit(sessionUser, args)
		case constants.Withdraw:
			err = service.Withdraw(sessionUser, args)
		case constants.Send:
			err = service.Send(sessionUser, args)
		case constants.Balance:
			_, err = service.Balance(sessionUser)
		case constants.Logout:
			err = service.Logout(sessionUser)
		case constants.Accounts:
			err = service.Accounts(sessionUser)
		default:
			fmt.Println("ERROR: ", constants.InvalidCommandMsg)
		}
		if err != nil {
			fmt.Println(fmt.Sprintf("ERROR: %s!", err))
		}
	}
}
