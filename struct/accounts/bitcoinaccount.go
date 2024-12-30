package accounts

import "fmt"

// BitcoinAccount struct
type BitcoinAccount struct {
    balance int
}

// NewBitcoinAccount constructor
func NewBitcoinAccount(balance int) *BitcoinAccount {
    return &BitcoinAccount{balance: balance}
}

// GetBalance method for BitcoinAccount
func (b *BitcoinAccount) GetBalance() int {
    return b.balance
}

// Deposit method for BitcoinAccount
func (b *BitcoinAccount) Deposit(amount int) {
    b.balance += amount
}

// Withdraw method for BitcoinAccount
func (b *BitcoinAccount) Withdraw(amount int) error {
    fee := 300
    newBalance := b.balance - amount - fee
    if newBalance < 0 {
        return fmt.Errorf("insufficient funds")
    }
    b.balance = newBalance
    return nil
}