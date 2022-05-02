package service

import (
	"TechnicalAssignment/cmd/db"
	"TechnicalAssignment/pkg/custError"
)

// drawDown abstracts the logic to:
// 1. Get user's current balance
// 2. Subtract amount from balance
// 3. Update balance to DB
func drawDown(username string, amount int) error {
	if amount <= 0 {
		return custError.NegativeValueError
	}

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

// topUp abstracts the logic to:
// 1. Get user's current balance
// 2. Add amount to balance
// 3. Update balance to DB
func topUp(username string, amount int) error {
	if amount < 0 {
		return custError.NegativeValueError
	}

	// Get balance from DB
	var bal int
	err := db.GlobalBalanceTable.Get(username, &bal)
	if err != nil {
		return custError.InternalDBError
	}

	// Update balance into DB
	err = db.GlobalBalanceTable.Put(username, bal+amount)
	if err != nil {
		return custError.InternalDBError
	}

	return nil
}
