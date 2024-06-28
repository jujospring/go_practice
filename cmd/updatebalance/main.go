package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(c *customer, t transaction) error {
	if t.transactionType == transactionDeposit {
		c.balance += t.amount
	} else if t.transactionType == transactionWithdrawal {
		if t.amount > c.balance {
			var err error = errors.New("insufficient funds")
			return err
		}
		c.balance -= t.amount
	} else {
		var err error = errors.New("unknown transaction type")
		return err
	}
	return nil
}
