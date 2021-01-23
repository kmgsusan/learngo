package main

import (
	"fmt"
	"log"

	"github.com/kmgsusan/learngo/accounts"
)

// main - mgkim
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
	account.ChangeOwner("mgkim")
	account.Deposit(30)
	fmt.Println(account)
	err := account.Withdraw(100)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account)
}
