package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"fmt"
	"github.com/rapidloop/skv"

	"errors"
)

func Register(user *User, args []string) error {
	fmt.Println("Registering...")
	if len(args) != 2 {
		return custError.InvalidNumArguments
	}

	username := args[0]
	password := args[1]

	// If username already exists, return error
	var val string
	err := db.GlobalPasswordTable.Get(username, &val)
	if err != skv.ErrNotFound {
		return custError.AccountExistsError
	}

	// Add username to DB
	err = db.GlobalPasswordTable.Put(username, password)
	if err != nil {
		return custError.InternalDBError
	}

	// Login user
	user.Username = username

	return nil
}

func Login(user *User, args []string) error {
	fmt.Println("Logging in...")
	if len(args) != 2 {
		return custError.InvalidNumArguments
	}

	username := args[0]
	user.Username = username

	return nil
}

func Deposit(user *User, args []string) error {
	fmt.Println("Depositing...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	//amount := args[0]

	return nil
}

func Withdraw(user *User, args []string) error {
	fmt.Println("Withdrawing...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	//amount := args[0]

	return nil
}

func Send(user *User, args []string) error {
	fmt.Println("Sending...")
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	//username := args[0]
	//amount := args[1]

	return nil
}

func Balance(user *User) error {
	fmt.Println("Retrieving Balance...")

	return nil
}

func Logout(user *User) error {
	fmt.Println("Logging out...")

	*user = User{}
	return nil
}

func Accounts(user *User) error {
	fmt.Println("Retrieving Accounts...")

	return nil
}
