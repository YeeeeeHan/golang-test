package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"bufio"
	"fmt"
	"github.com/rapidloop/skv"
	"os"

	"errors"
)

func Register(wallet *Wallet, args []string) error {
	fmt.Println("Registering...")
	if len(args) != 2 {
		return custError.InvalidNumArguments
	}

	username := args[0]
	password := args[1]

	err := db.CreateUser(username, password)
	if err != nil {
		return err
	}

	// Login wallet
	wallet.Username = username

	return nil
}

func Login(wallet *Wallet, args []string) error {
	fmt.Println("Logging in...")
	if len(args) != 2 {
		return custError.InvalidNumArguments
	}

	username := args[0]
	password := args[1]

	var val string
	// Username does not exist
	err := db.GlobalPasswordTable.Get(username, &val)
	if err == skv.ErrNotFound {
		return custError.WrongCredentialsError
	}
	// Wrong password
	if val != password {
		return custError.WrongCredentialsError
	}

	// Login usr
	wallet.Username = username

	return nil
}

func Deposit(wallet *Wallet, args []string) error {
	fmt.Println("Depositing...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	//amount := args[0]

	return nil
}

func Withdraw(user *Wallet, args []string) error {
	fmt.Println("Withdrawing...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	//amount := args[0]

	return nil
}

func Send(user *Wallet, args []string) error {
	fmt.Println("Sending...")
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}
	//username := args[0]
	//amount := args[1]

	return nil
}

func Balance(user *Wallet) error {
	fmt.Println("Retrieving Balance...")

	return nil
}

func Logout(user *Wallet) error {
	fmt.Println("Logging out...")

	*user = Wallet{}
	return nil
}

func Accounts(user *Wallet) error {
	fmt.Println("Retrieving Accounts...")

	if user.Username != "admin" {
		return custError.PermissionError
	}

	// Open file
	file, err := os.Open(constants.UsernameFile)
	if err != nil {
		return custError.InternalDBError
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// For each username, print balance
	for scanner.Scan() {
		var bal int
		err = db.GlobalBalanceTable.Get(scanner.Text(), &bal)
		fmt.Println(fmt.Sprintf("Username: %s, Balance: %v", scanner.Text(), bal))
	}

	if err := scanner.Err(); err != nil {
		return custError.InternalDBError
	}

	return nil
}
