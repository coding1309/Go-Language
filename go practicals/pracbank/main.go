package main

import (
	"errors"
	"fmt"
)

// Define a struct for the bank account
type BankAccount struct {
	owner   string
	balance float64
}

// Method to deposit funds into the account
func (acc *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	acc.balance += amount
	fmt.Printf("Deposited %.2f. New balance: %.2f\n", amount, acc.balance)
	return nil
}

// Method to withdraw funds from the account
func (acc *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if amount > acc.balance {
		return errors.New("insufficient funds")
	}
	acc.balance -= amount
	fmt.Printf("Withdrawn %.2f. New balance: %.2f\n", amount, acc.balance)
	return nil
}

// Method to check the balance of the account
func (acc *BankAccount) Balance() float64 {
	return acc.balance
}

func main() {
	// Create a new bank account
	account := BankAccount{owner: "Neel Patel", balance: 10000.0}

	// Perform transactions with error handling
	err := account.Deposit(-20.0)
	if err != nil {
		fmt.Println("Error depositing funds:", err)
	}

	err = account.Withdraw(2000.0)
	if err != nil {
		fmt.Println("Error withdrawing funds:", err)
	}

	fmt.Println("Account balance:", account.Balance())
}
