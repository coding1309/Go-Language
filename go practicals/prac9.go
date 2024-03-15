package main

import (
	"fmt"
)

type Account struct {
	balance float64
	closed  bool
}

func (a *Account) deposit(amount float64) (newBalance float64, succeeded bool) {
	if a.closed {
		return 0, false
	}
	newBalance = a.balance + amount
	a.balance = newBalance
	return newBalance, true
}

func (a *Account) withdraw(amount float64) (newBalance float64, succeeded bool) {
	if a.closed {
		return 0, false
	}
	if amount > a.balance {
		return 0, false
	}
	newBalance = a.balance - amount
	a.balance = newBalance
	return newBalance, true
}

func (a *Account) balances() (balance float64, succeeded bool) {
	if a.closed {
		return 0, false
	}
	balance = a.balance
	return balance, true
}

func (a *Account) close() (finalBalance float64, succeeded bool) {
	if a.closed {
		return 0, false
	}
	a.closed = true
	finalBalance = a.balance
	a.balance = 0
	return finalBalance, true
}

func main() {
	account := &Account{balance: 1000}
	fmt.Println("Hello !!! Welcome to Deposit&Withdrawal Machine\n")
	fmt.Printf("By Default Amount :  1000\n\n")
	var depositAmount, withdrawAmount float64
	fmt.Print("Enter amount to be deposited: ")
	fmt.Scan(&depositAmount)
	newBalance, _ := account.deposit(depositAmount)
	fmt.Printf("Amount Deposited: %.2f\n", depositAmount)
	fmt.Printf("New Balance: %.2f\n", newBalance)
	fmt.Print("Enter amount to be withdrawn: ")
	fmt.Scan(&withdrawAmount)
	newBalance, _ = account.withdraw(withdrawAmount)
	fmt.Printf("You Withdrew: %.2f\n", withdrawAmount)
	fmt.Printf("Net Available Balance = %.2f\n", newBalance)
}
