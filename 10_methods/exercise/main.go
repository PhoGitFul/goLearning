package main

import (
	"fmt"
	"goLearning/10_methods/exercise/Txns"
)

func main() {
	ba1 := &Txns.BankAccount{}
	var cred1, cred2, deb1, deb2, deb3 float64 = 100.5, 10.5, -50, 112, 111

	ba1.Deposit(cred1)
	fmt.Printf("After depositing %.2f, balance: %.2f\n", cred1, ba1.GetBalance())

	err := ba1.Withdraw(deb1)
	if err != nil {
		fmt.Printf("Attempting to Withdraw %.2f, Error: %s", deb1, err)
	} else {
		fmt.Printf("After withdrawing %.2f, balance: %.2f\n", deb1, ba1.GetBalance())
	}

	ba1.Deposit(cred2)
	fmt.Printf("After depositing %.2f, balance: %.2f\n", cred2, ba1.GetBalance())

	err2 := ba1.Withdraw(deb2)
	if err2 != nil {
		fmt.Printf("Attempting to Withdraw %.2f, Error: %s", deb2, err2)
		fmt.Printf("After withdrawing %.2f, balance: %.2f\n", deb2, ba1.GetBalance())
	}

	err3 := ba1.Withdraw(deb3)
	if err3 != nil {
		fmt.Printf("Attempting to Withdraw %.2f, Error: %s", deb3, err3)
		fmt.Printf("After withdrawing %.2f, balance: %.2f\n", deb3, ba1.GetBalance())
	}
}
