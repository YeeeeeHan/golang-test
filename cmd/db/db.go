package db

import (
	"TechnicalAssignment/pkg/constants"
	"TechnicalAssignment/pkg/custError"
	"github.com/rapidloop/skv"
	"os"
)

var GlobalBalanceTable *skv.KVStore
var GlobalPasswordTable *skv.KVStore

func InitTables() (*skv.KVStore, *skv.KVStore, error) {
	balanceStore, err := skv.Open(constants.BalanceFile)
	if err != nil {
		return nil, nil, err
	}

	passwordStore, err := skv.Open(constants.PasswordFile)
	if err != nil {
		return nil, nil, err
	}

	return balanceStore, passwordStore, nil
}

// CreateUser
func CreateUser(username, password string) error {
	// If username already exists in passwordTable, return error
	var getPw string
	err := GlobalPasswordTable.Get(username, &getPw)
	if err != skv.ErrNotFound {
		return custError.AccountExistsError
	}

	// If username already exists in balanceTable, return error
	var getB string
	err = GlobalBalanceTable.Get(username, &getB)
	if err != skv.ErrNotFound {
		return custError.AccountExistsError
	}

	// Add username to passwordTable
	err = GlobalPasswordTable.Put(username, password)
	if err != nil {
		return custError.InternalDBError
	}

	// Add username to passwordTable
	err = GlobalBalanceTable.Put(username, 0)
	if err != nil {
		return custError.InternalDBError
	}

	// Add username to usernameTable
	f, err := os.OpenFile(constants.UsernameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	_, err = f.WriteString(username + "\n")
	if err != nil {
		return err
	}

	return nil
}

func CloseDB(store *skv.KVStore) {
	store.Close()
}
