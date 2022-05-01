package db

import (
	"github.com/rapidloop/skv"
)

var GlobalBalanceTable *skv.KVStore
var GlobalPasswordTable *skv.KVStore

func InitBalanceTable() (*skv.KVStore, error) {
	store, err := skv.Open("cmd/db/balance.db")
	if err != nil {
		return nil, err
	}

	return store, nil
}

func InitPasswordTable() (*skv.KVStore, error) {
	store, err := skv.Open("cmd/db/password.db")
	if err != nil {
		return nil, err
	}

	return store, nil
}

func CloseDB(store *skv.KVStore) {
	store.Close()
}
