package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"bufio"
	"fmt"
	"github.com/rapidloop/skv"
	"os"
	"strings"

	"errors"
)

func Register(wallet *Wallet, args []string) error {
	if len(args) != 2 {
		return custError.InvalidNumArguments
	}

	username := args[0]
	password := args[1]

	fmt.Println(fmt.Sprintf("Username: %s, password: %s", username, password))

	err := db.CreateUser(username, password)
	if err != nil {
		return err
	}

	// Login wallet
	wallet.Username = username

	fmt.Println(fmt.Sprintf("Successfully registered %s", username))
	return nil
}

func Login(wallet *Wallet, args []string) error {
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

	fmt.Println(fmt.Sprintf("Successfully logged into %s", username))
	return nil
}

// Deposit gets the curr value, adds deposit, and updates DB
func Deposit(wallet *Wallet, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	amountStr := args[0]

	err := topUp(wallet.Username, amountStr)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully deposited $%v", amountStr))
	return nil
}

func Withdraw(wallet *Wallet, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	amountStr := args[0]

	err := drawDown(wallet.Username, amountStr)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully withdrew $%v", amountStr))
	return nil
}

func Send(wallet *Wallet, args []string) error {
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	username := args[0]
	amount := args[1]

	//drawDown from source account
	err := drawDown(wallet.Username, amount)
	if err != nil {
		return err
	}

	// topUp to destination account
	err = topUp(username, amount)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully sent (%s): $%s", username, amount))
	return nil
}

func Balance(wallet *Wallet) error {

	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(wallet.Username, &bal)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Current balance for (%s): $%d", wallet.Username, bal))
	return nil
}

func Logout(wallet *Wallet) error {

	// Logout
	*wallet = Wallet{}

	fmt.Println(fmt.Sprintf("Successfully logged out of %s", wallet.Username))
	return nil
}

func Accounts(wallet *Wallet) error {

	if wallet.Username != "admin" {
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
		// Get balance from DB
		err = db.GlobalBalanceTable.Get(strings.TrimSpace(scanner.Text()), &bal)
		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("Username: %s, Balance: $%d", scanner.Text(), bal))
	}

	return nil
}
