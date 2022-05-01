package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/custError"
)

func drawDown(username string, amount int) error {
	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(username, &bal)

	// Sanity check
	if amount > bal {
		return custError.InsufficientFunds
	}

	// Update balance into DB
	err = db.GlobalBalanceTable.Put(username, bal-amount)
	if err != nil {
		return custError.InternalDBError
	}

	return nil
}

func topUp(username string, amount int) error {
	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(username, &bal)

	// Update balance into DB
	err = db.GlobalBalanceTable.Put(username, bal+amount)
	if err != nil {
		return custError.InternalDBError
	}

	return nil
}
