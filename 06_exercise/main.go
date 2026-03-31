package main

import (
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) (err error) {
	if amount > a.Balance {
		err = fmt.Errorf("insufficient balance")
		return err
	}
	a.Balance -= amount
	return nil
}

func (a *BankAccount) String() string {
	return fmt.Sprintf("[%s] Balance: %.2f", a.Owner, a.Balance)
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		Owner:   owner,
		Balance: initialBalance,
	}
}

func main() {
	account := NewBankAccount("Alice", 1000)

	account.Deposit(100)
	if err := account.Withdraw(200); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	if err := account.Withdraw(2000); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	str := account.String()
	fmt.Print(str)
}
