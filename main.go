package main

import (
	"fmt"

	"github.com/kmgsusan/learngo/accounts"
)

// main - mgkim
func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println("ok")
}
