package entity

import (
	"context"
	"errors"
)

type Account struct {
	Id      int
	Balance float64
}

func (acc *Account) Deposit(ctx context.Context, amount float64) error {
	if amount < 0 {
		return errors.New("cannot deposit negative value")
	}
	acc.Balance += amount
	return nil
}

func (acc *Account) Withdraw(ctx context.Context, amount float64) error {
	if acc.Balance-amount < 0 {
		return errors.New("insufficient funds")
	}
	acc.Balance -= amount
	return nil
}

func (acc *Account) GetBalance(ctx context.Context) float64 {
	return acc.Balance
}
