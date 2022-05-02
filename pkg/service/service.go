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

// Register creates a user in the DB and logs user in
func Register(wallet *Wallet, args []string) error {
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

	fmt.Println(fmt.Sprintf("Successfully registered %s", username))
	return nil
}

// Login Authenticates user using passwordTable
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

	// Login user
	wallet.Username = username

	fmt.Println(fmt.Sprintf("Successfully logged into %s", username))
	return nil
}

// Deposit adds amount to user's wallet balance
func Deposit(wallet *Wallet, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	amountStr := args[0]
	// Convert topUp value to int
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	err = topUp(wallet.Username, amount)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully deposited $%v", amountStr))
	return nil
}

// Withdraw reduces amount from user's wallet balance
func Withdraw(wallet *Wallet, args []string) error {
	if len(args) != 1 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	// Convert drawDown value to int
	amountStr := args[0]
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	err = drawDown(wallet.Username, amount)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully withdrew $%v", amountStr))
	return nil
}

// Send reduces amount from source wallet balance and increases amount of destination wallet balance
func Send(wallet *Wallet, args []string) error {
	if len(args) != 2 {
		return errors.New(constants.InvalidNumArgumentsMsg)
	}

	destUser := args[0]
	amountStr := args[1]

	// If destUser does not exist
	var bal int
	err := db.GlobalBalanceTable.Get(destUser, &bal)
	if err != nil {
		return custError.AccountsDoesNotExistError
	}

	// Convert amount value to int
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	//drawDown from source account
	err = drawDown(wallet.Username, amount)
	if err != nil {
		return err
	}

	// topUp to destination account
	err = topUp(destUser, amount)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Successfully sent (%s): $%d", destUser, amount))
	return nil
}

// Balance retrieves the balance for the user in the current session
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

// Logout set the current user in the session to an empty Wallet struct end the user session
func Logout(wallet *Wallet) error {
	// Logout
	*wallet = Wallet{}

	fmt.Println(fmt.Sprintf("Successfully logged out of %s", wallet.Username))
	return nil
}

// Accounts reads all the username from the UsernameFile and retrieves all their balances.
// Can only be operated by admin user.
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
