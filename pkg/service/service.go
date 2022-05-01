package service

import (
	"TechnicalAssignment/pkg/constants"
	"fmt"

	"errors"
)

func Login(user *User, args []string) error {
	fmt.Println("Processing 'Login'...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	username := args[0]
	user.username = username

	return nil
}

func Deposit(user *User, args []string) error {
	fmt.Println("Processing 'Deposit'...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	amount := args[0]

	return nil
}

func Withdraw(user *User, args []string) error {
	fmt.Println("Processing 'Withdraw'...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	amount := args[0]

	return nil
}

func Send(user *User, args []string) error {
	fmt.Println("Processing 'Send'...")
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	username := args[0]
	amount := args[1]

	return nil
}

func Balance(user *User) error {
	fmt.Println("Processing 'Balance'...")

	return nil
}

func Logout(user *User) error {
	fmt.Println("Processing 'Logout'...")

	*user = User{}
	return nil
}

func Accounts(user *User) error {
	fmt.Println("Processing 'Accounts'...")

	return nil
}
