package main

import (
    "fmt"
	"github.com/nazurudeen23/struct/accounts"
)

func main() {
    wf := accounts.NewWellsFargo(100)
    wf.Deposit(200)
    wf.Deposit(370)
    err := wf.Withdraw(117)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("wf balance is", wf.GetBalance())

    myAccounts := []accounts.BankAccount{
        accounts.NewWellsFargo(0),
        accounts.NewBitcoinAccount(0),
    }

    for _, account := range myAccounts {
        account.Deposit(500)
        err := account.Withdraw(70)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Println("balance is", account.GetBalance())
    }

    // Example of type assertion
    wfAccount, ok := myAccounts[0].(*accounts.WellsFargo)
    if ok {
        fmt.Println("This is a Wells Fargo account:", wfAccount)
    }

        btcAccount, ok := myAccounts[1].(*accounts.BitcoinAccount)
        if ok {
                fmt.Println("This is a Bitcoin account:", btcAccount)
        }
}