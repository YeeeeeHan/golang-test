package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"bufio"
	"fmt"
	"github.com/rapidloop/skv"
	"os"
	"strconv"
	"strings"

	"errors"
)

func Register(wallet *Wallet, args []string) error {
	fmt.Println("Registering...")
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

// Deposit gets the curr value, adds deposit, and updates DB
func Deposit(wallet *Wallet, args []string) error {
	fmt.Println("Depositing...")
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	amountStr := args[0]

	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(wallet.Username, &bal)

	// Convert deposit value to int
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	// Update balance into DB
	err = db.GlobalBalanceTable.Put(wallet.Username, bal+amount)
	if err != nil {
		return custError.InternalDBError
	}

	fmt.Println(fmt.Sprintf("Successfully deposited $%v", amountStr))
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

func Balance(wallet *Wallet) error {
	fmt.Println("Retrieving Balance...")

	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(wallet.Username, &bal)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Current balance for (%s): %d", wallet.Username, bal))

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
		// Get balance from DB
		err = db.GlobalBalanceTable.Get(strings.TrimSpace(scanner.Text()), &bal)
		if err != nil {
			return err
		}

		fmt.Println(fmt.Sprintf("Username: %s, Balance: %d", scanner.Text(), bal))
	}

	return nil
}
