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

func TestMain(m *testing.M)

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
	os.Exit(exitVal)

	////log.Println("SUCCESSFUL INIT")
	////// If outputUsername already exists in passwordTable, return error
	//var getPw string
	//err = db.GlobalPasswordTable.Get("adminzzz", &getPw)
	//if err != nil {
	//	log.Println("errrr??", err)
	//}
	//log.Println("adminzzz PW", getPw)
}

func TestLogin(t *testing.T) {
	testTable := []struct {
		name           string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "parse valid args and credentials",
			inputArgs:      []string{"admin", "password123"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "parse invalid args number",
			inputArgs:      []string{"admin"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "",
		},
		{
			name:           "parse wrong credentials",
			inputArgs:      []string{"adminzzzz", "password123"},
			outputErr:      custError.WrongCredentialsError,
			outputUsername: "",
		},
		{
			name:           "parse wrong credentials",
			inputArgs:      []string{"admin", "password123zzzzz"},
			outputErr:      custError.WrongCredentialsError,
			outputUsername: "",
		},
		{
			name:           "parse wrong credentials",
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

//func TestBalance(t *testing.T) {
//	testTable := []struct {
//		name           string
//		outputErr      error
//		outputBalance  int
//		outputUsername string
//	}{
//		{
//			name:           "parse valid balance arguments",
//			outputErr:      nil,
//			outputUsername: "admin",
//		},
//	}
//
//	for _, tt := range testTable {
//		// Login Wallet
//		sessionUser = Wallet{tt.outputUsername}
//
//		t.Run(tt.name, func(t *testing.T) {
//			bal, err := Balance(&sessionUser)
//			assert.Equal(t, tt.outputErr, err)
//			assert.Equal(t, tt.outputBalance, bal)
//			assert.Equal(t, tt.outputUsername, sessionUser.Username)
//		})
//	}
//}

func TestDeposit(t *testing.T) {
	testTable := []struct {
		name           string
		inputArgs      []string
		outputErr      error
		outputUsername string
	}{
		{
			name:           "parse valid deposit value",
			inputArgs:      []string{"10"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "parse invalid deposit value",
			inputArgs:      []string{"$10"},
			outputErr:      custError.InvalidArguments,
			outputUsername: "admin",
		},
		{
			name:           "parse negative deposit value",
			inputArgs:      []string{"-10"},
			outputErr:      custError.NegativeValueError,
			outputUsername: "admin",
		},
		{
			name:           "parse invalid argument numbers",
			inputArgs:      []string{"10", "10"},
			outputErr:      custError.InvalidNumArguments,
			outputUsername: "admin",
		},
		{
			name:           "parse no account exists",
			inputArgs:      []string{"10"},
			outputErr:      custError.InternalDBError,
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
			name:           "parse valid withdraw value",
			depositArgs:    []string{"10"},
			inputArgs:      []string{"10"},
			outputErr:      nil,
			outputUsername: "admin",
		},
		{
			name:           "parse withdraw more than deposit",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"99999999"},
			outputErr:      custError.InsufficientFunds,
			outputUsername: "admin",
		},
		{
			name:           "parse withdraw negative value",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"-10"},
			outputErr:      custError.NegativeValueError,
			outputUsername: "admin",
		},
		{
			name:           "parse withdraw invalid value",
			depositArgs:    []string{"0"},
			inputArgs:      []string{"$10"},
			outputErr:      custError.InvalidArguments,
			outputUsername: "admin",
		},
		{
			name:           "parse withdraw wrong argument numbers",
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

//func TestSend(t *testing.T) {
//	testTable := []struct {
//		name           string
//		depositArgs    []string
//		inputArgs      []string
//		outputErr      error
//		outputUsername string
//	}{
//		{
//			name:           "parse valid send value",
//			depositArgs:    []string{"10"},
//			inputArgs:      []string{"receiver", "10"},
//			outputErr:      nil,
//			outputUsername: "admin",
//		},
//	}
//
//	for _, tt := range testTable {
//		// Login Wallet
//		sessionUser = Wallet{tt.outputUsername}
//
//		t.Run(tt.name, func(t *testing.T) {
//			err := Deposit(&sessionUser, tt.depositArgs)
//			assert.Equal(t, nil, err)
//
//			err = Send(&sessionUser, tt.inputArgs)
//			assert.Equal(t, tt.outputErr, err)
//			assert.Equal(t, tt.outputUsername, sessionUser.Username)
//		})
//	}
//}
