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
		return errors.New("cannot withdraw a negative amount\n")
	}
	if ba.balance < debit {
		return errors.New("insufficient funds\n")
	}
	ba.balance -= debit
	return nil
	// if ba.balance < debit {
	// 	ba.balance = 0
	// } else {
	// 	ba.balance -= debit
	// }
}
func (ba BankAccount) GetBalance() float64 {
	return ba.balance
}

// package Txns

// import (
// 	"errors"
// 	"fmt"
// )

// // Deposit adds the specified amount to the account balance.
// func (ba *BankAccount) Deposit(amount float64) {
// 	if amount < 0 {
// 		fmt.Println("Cannot deposit a negative amount.")
// 		return
// 	}
// 	ba.balance += amount
// }

// // Withdraw subtracts the specified amount from the account balance.
// // Returns an error if the withdrawal amount exceeds the current balance.
// func (ba *BankAccount) Withdraw(amount float64) error {
// 	if amount < 0 {
// 		return errors.New("cannot withdraw a negative amount")
// 	}
// 	if ba.balance < amount {
// 		return errors.New("insufficient funds")
// 	}
// 	ba.balance -= amount
// 	return nil
// }
