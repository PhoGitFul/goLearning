package Txns

import (
	"errors"
)

type BankAccount struct {
	balance float64
}

func (ba *BankAccount) Deposit(credit float64) {
	ba.balance += credit
}

func (ba *BankAccount) Withdraw(debit float64) error {
	if debit < 0 {
		return errors.New("cannot withdraw a negative amount")
	}
	if ba.balance < debit {
		return errors.New("insufficient funds")
	}
	ba.balance -= debit
	return nil
}
func (ba BankAccount) GetBalance() float64 {
	return ba.balance
}
