package service

import (
	"TechnicalAssignment/pkg/constants"

	"errors"
)

func Login(user *User, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	return nil
}

func Deposit(user *User, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	return nil
}

func Withdraw(user *User, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	return nil
}

func Send(user *User, args []string) error {
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	return nil
}

func Balance(user *User) error {

	return nil
}

func Logout(user *User) error {

	return nil
}

func Accounts() error {

	return nil
}
