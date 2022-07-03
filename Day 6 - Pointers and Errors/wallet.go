package main

import (
	"fmt"
	"errors"
)


type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= value
	
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(value Bitcoin) {
 	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}