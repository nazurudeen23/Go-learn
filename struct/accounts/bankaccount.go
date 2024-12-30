package accounts

import "fmt"

// BankAccount interface
type BankAccount interface {
    GetBalance() int
    Deposit(amount int)
    Withdraw(amount int) error
}

// WellsFargo struct
type WellsFargo struct {
    balance int
}

// NewWellsFargo constructor
func NewWellsFargo(balance int) *WellsFargo {
    return &WellsFargo{balance: balance}
}

// GetBalance method for WellsFargo
func (w *WellsFargo) GetBalance() int {
    return w.balance
}

// Deposit method for WellsFargo
func (w *WellsFargo) Deposit(amount int) {
    w.balance += amount
}

// Withdraw method for WellsFargo
func (w *WellsFargo) Withdraw(amount int) error {
    newBalance := w.balance - amount
    if newBalance < 0 {
        return fmt.Errorf("insufficient funds")
    }
    w.balance = newBalance
    return nil
}