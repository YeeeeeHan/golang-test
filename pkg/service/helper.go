package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/custError"
	"strconv"
)

func drawDown(username, amountStr string) error {
	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(username, &bal)

	// Convert drawDown value to int
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

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

func topUp(username, amountStr string) error {
	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(username, &bal)

	// Convert topUp value to int
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return err
	}

	// Update balance into DB
	err = db.GlobalBalanceTable.Put(username, bal+amount)
	if err != nil {
		return custError.InternalDBError
	}

	return nil
}
