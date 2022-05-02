package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

var sessionUser Wallet

func TestMain(m *testing.M) {
	// Creating files before test
	_, e := os.Create(constants.UsernameFileTest)
	if e != nil {
		log.Fatal(e)
	}
	_, e = os.Create(constants.BalanceFileTest)
	if e != nil {
		log.Fatal(e)
	}
	_, e = os.Create(constants.PasswordFileTest)
	if e != nil {
		log.Fatal(e)
	}

	db.GlobalUsernameTable = constants.UsernameFileTest
	var initerr error
	db.GlobalBalanceTable, db.GlobalPasswordTable, initerr = db.InitTables(constants.BalanceFileTest, constants.PasswordFileTest)
	if initerr != nil {
		log.Println(initerr)
	}

	// Register Admin
	err := db.CreateUser("admin", "password123")
	if err != nil {
		log.Println(err)
	}

	// Register Receiver
	err = db.CreateUser("receiver", "password123")
	if err != nil {
		log.Println(err)
	}

	exitVal := m.Run()

	// Removing files after test
	e = os.Remove(constants.UsernameFileTest)
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(constants.BalanceFileTest)
	if e != nil {
		log.Fatal(e)
	}
	e = os.Remove(constants.PasswordFileTest)
	if e != nil {
		log.Fatal(e)
	}

	os.Exit(exitVal)

	////log.Println("SUCCESSFUL INIT")
	////// If username already exists in passwordTable, return error
	//var getPw string
	//err = db.GlobalPasswordTable.Get("adminzzz", &getPw)
	//if err != nil {
	//	log.Println("errrr??", err)
	//}
	//log.Println("adminzzz PW", getPw)
}

func TestTopUp(t *testing.T) {
	testTable := []struct {
		name        string
		inputAmount int
		username    string
		outputErr   error
	}{
		{
			name:        "successfully topUp",
			inputAmount: 10,
			username:    "admin",
			outputErr:   nil,
		},
		{
			name:        "user does not exist",
			inputAmount: 10,
			username:    "adminzzzzz",
			outputErr:   custError.AccountsDoesNotExistError,
		},
		{
			name:        "negative amount",
			inputAmount: -10,
			username:    "adminzzzzz",
			outputErr:   custError.NegativeValueError,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			err := topUp(tt.username, tt.inputAmount)
			assert.Equal(t, tt.outputErr, err)
		})
	}
}

func TestDrawDown(t *testing.T) {
	testTable := []struct {
		name        string
		inputAmount int
		username    string
		outputErr   error
	}{
		{
			name:        "successfully drawdown",
			inputAmount: 10,
			username:    "admin",
			outputErr:   nil,
		},
		{
			name:        "user does not exist",
			inputAmount: 10,
			username:    "adminzzzzz",
			outputErr:   custError.AccountsDoesNotExistError,
		},
		{
			name:        "negative amount",
			inputAmount: -10,
			username:    "adminzzzzz",
			outputErr:   custError.NegativeValueError,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			err := drawDown(tt.username, tt.inputAmount)
			assert.Equal(t, tt.outputErr, err)
		})
	}
}

func TestRegister(t *testing.T) {
	testTable := []struct {
		name           string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "register successful and login",
			inputArgs:      []string{"account1", "password1"},
			outputErr:      nil,
			outputUsername: "account1",
		},
		{
			name:           "register invalid num args",
			inputArgs:      []string{"account1", "password1", "123"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "",
		},
		{
			name:           "register user already exists",
			inputArgs:      []string{"admin", "password123"},
			outputErr:      custError.AccountAlreadyExistsError,
			outputUsername: "",
		},
	}

	for _, tt := range testTable {
		// Init Wallet
		sessionUser = Wallet{}

		t.Run(tt.name, func(t *testing.T) {
			err := Register(&sessionUser, tt.inputArgs)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestLogin(t *testing.T) {
	testTable := []struct {
		name           string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "login successfully with correct credentials",
			inputArgs:      []string{"admin", "password123"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "login with invalid args number",
			inputArgs:      []string{"admin"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "",
		},
		{
			name:           "login with wrong credentials",
			inputArgs:      []string{"adminzzzz", "password123"},
			outputErr:      custError.WrongCredentialsError,
			outputUsername: "",
		},
		{
			name:           "login with wrong credentials",
			inputArgs:      []string{"admin", "password123zzzzz"},
			outputErr:      custError.WrongCredentialsError,
			outputUsername: "",
		},
		{
			name:           "login with wrong credentials",
			inputArgs:      []string{"admin", "password123zzzzz"},
			outputErr:      custError.WrongCredentialsError,
			outputUsername: "",
		},
	}

	for _, tt := range testTable {
		// Init Wallet
		sessionUser = Wallet{}

		t.Run(tt.name, func(t *testing.T) {
			err := Login(&sessionUser, tt.inputArgs)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestBalance(t *testing.T) {
	testTable := []struct {
		name           string
		outputErr      error
		outputBalance  int
		outputUsername string
	}{
		{
			name:           "successfully request balance",
			outputErr:      nil,
			outputUsername: "admin",
			outputBalance:  0,
		},
		{
			name:           "no account exists",
			outputErr:      custError.AccountsDoesNotExistError,
			outputUsername: "adminzzz",
			outputBalance:  0,
		},
	}

	for _, tt := range testTable {
		// Login Wallet
		sessionUser = Wallet{tt.outputUsername}

		t.Run(tt.name, func(t *testing.T) {
			bal, err := Balance(&sessionUser)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputBalance, bal)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestDeposit(t *testing.T) {
	testTable := []struct {
		name           string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "successfully deposit",
			inputArgs:      []string{"10"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "deposit invalid value format",
			inputArgs:      []string{"$10"},
			outputErr:      custError.InvalidArguments,
			outputUsername: "admin",
		},
		{
			name:           "deposit negative value",
			inputArgs:      []string{"-10"},
			outputErr:      custError.NegativeValueError,
			outputUsername: "admin",
		},
		{
			name:           "deposit invalid argument numbers",
			inputArgs:      []string{"10", "10"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "admin",
		},
		{
			name:           "no account exists",
			inputArgs:      []string{"10"},
			outputErr:      custError.AccountsDoesNotExistError,
			outputUsername: "adminzzz",
		},
	}

	for _, tt := range testTable {
		// Login Wallet
		sessionUser = Wallet{tt.outputUsername}

		t.Run(tt.name, func(t *testing.T) {
			err := Deposit(&sessionUser, tt.inputArgs)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestWithdraw(t *testing.T) {
	testTable := []struct {
		name           string
		depositArgs    []string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "successfully withdraw",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"10"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "withdraw more than funds",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"99999999"},
			outputErr:      custError.InsufficientFunds,
			outputUsername: "admin",
		},
		{
			name:           "withdraw negative value",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"-10"},
			outputErr:      custError.NegativeValueError,
			outputUsername: "admin",
		},
		{
			name:           "withdraw invalid value format",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"$10"},
			outputErr:      custError.InvalidArguments,
			outputUsername: "admin",
		},
		{
			name:           "withdraw with wrong argument numbers",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"10", "10"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "admin",
		},
	}

	for _, tt := range testTable {
		// Login Wallet
		sessionUser = Wallet{tt.outputUsername}

		t.Run(tt.name, func(t *testing.T) {
			err := Deposit(&sessionUser, tt.depositArgs)
			assert.Equal(t, nil, err)

			err = Withdraw(&sessionUser, tt.inputArgs)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestSend(t *testing.T) {
	testTable := []struct {
		name                  string
		depositArgs           []string
		inputArgs             []string
		inputReceiverUsername string
		outputErr             error
		outputUsername        string
		outputSenderBalance   int
		outputReceiverBalance int
	}{
		{
			name:                  "successfully send and receive",
			depositArgs:           []string{"20"},
			inputArgs:             []string{"receiver", "9"},
			inputReceiverUsername: "receiver",
			outputErr:             nil,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 9,
		},
		{
			name:                  "invalid send value format",
			depositArgs:           []string{"0"},
			inputArgs:             []string{"receiver", "$9"},
			inputReceiverUsername: "receiver",
			outputErr:             custError.InvalidArguments,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 9,
		},
		{
			name:                  "invalid number of args",
			depositArgs:           []string{"0"},
			inputArgs:             []string{"receiver", "9", "9"},
			inputReceiverUsername: "receiver",
			outputErr:             custError.InvalidNumArguments,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 9,
		},
		{
			name:                  "insufficient funds",
			depositArgs:           []string{"0"},
			inputArgs:             []string{"receiver", "99999"},
			inputReceiverUsername: "receiver",
			outputErr:             custError.InsufficientFunds,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 9,
		},
		{
			name:                  "send negative value",
			depositArgs:           []string{"0"},
			inputArgs:             []string{"receiver", "-9"},
			inputReceiverUsername: "receiver",
			outputErr:             custError.NegativeValueError,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 9,
		},
		{
			name:                  "receiver account does not exist",
			depositArgs:           []string{"0"},
			inputArgs:             []string{"receiverzzzzz", "-9"},
			inputReceiverUsername: "receiverzzzzz",
			outputErr:             custError.AccountsDoesNotExistError,
			outputUsername:        "admin",
			outputSenderBalance:   11,
			outputReceiverBalance: 0,
		},
	}

	for _, tt := range testTable {
		// Login Wallet
		sessionUser = Wallet{tt.outputUsername}

		// Login Wallet
		receiverUser := Wallet{tt.inputReceiverUsername}

		t.Run(tt.name, func(t *testing.T) {
			err := Deposit(&sessionUser, tt.depositArgs)
			assert.Equal(t, nil, err)

			err = Send(&sessionUser, tt.inputArgs)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)

			bal, err := Balance(&sessionUser)
			assert.Equal(t, tt.outputSenderBalance, bal)

			balRec, err := Balance(&receiverUser)
			assert.Equal(t, tt.outputReceiverBalance, balRec)
		})
	}
}

func TestLogout(t *testing.T) {
	testTable := []struct {
		name           string
		sessionUser    Wallet
		outputErr      error
		outputUsername string
	}{
		{
			name:           "successfully log out",
			outputErr:      nil,
			outputUsername: "",
		},
	}

	for _, tt := range testTable {
		// Init Wallet
		sessionUser = Wallet{}

		t.Run(tt.name, func(t *testing.T) {
			err := Logout(&sessionUser)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}

func TestAccounts(t *testing.T) {
	testTable := []struct {
		name           string
		sessionUser    Wallet
		outputErr      error
		outputUsername string
	}{
		{
			name:           "successfully request all accounts",
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "calling function with non-admin account",
			outputErr:      custError.PermissionError,
			outputUsername: "receiver",
		},
	}

	for _, tt := range testTable {
		// Init Wallet
		sessionUser = Wallet{Username: tt.outputUsername}

		t.Run(tt.name, func(t *testing.T) {
			err := Accounts(&sessionUser)
			assert.Equal(t, tt.outputErr, err)
			assert.Equal(t, tt.outputUsername, sessionUser.Username)
		})
	}
}
