package main

import (
	"fmt"
)

type Accounts struct {
	balance int
}

func (a *Accounts) Deposit(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid deposit amount: %d", amount)
	}
	a.balance += amount
	return nil
}

func (a *Accounts) Withdraw(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid withdrawal amount: %d", amount)
	}
	if a.balance < amount {
		return fmt.Errorf("insufficient funds for withdrawal of: %d", amount)
	}
	a.balance -= amount
	return nil
}

func (a *Accounts) Balances() int {
	return a.balance
}

func main() {
	accounts := &Accounts{balance: 1000}

	// Deposit $500
	err := accounts.Deposit(500)
	if err != nil {
		fmt.Println("there was an error:", err)
	} else {
		fmt.Println("deposit successful, new balance:", accounts.Balances())
	}

	// Withdraw $1500
	err = accounts.Withdraw(1500)
	if err != nil {
		fmt.Println("there was an error:", err)
	} else {
		fmt.Println("withdrawal successful, new balance:", accounts.Balances())
	}

	// Withdraw $400
	err = accounts.Withdraw(400)
	if err != nil {
		fmt.Println("there was an error:", err)
	} else {
		fmt.Println("withdrawal successful, new balance:", accounts.Balances())
	}
}
