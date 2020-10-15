package main

import (
	"banco/accounts"
	"fmt"
)

func main() {
	alexAccount := accounts.CheckingAccount{}
	alexAccount.Holder = "Alex"
	alexAccount.Balance = 500

	fmt.Println(alexAccount.Balance)

	fmt.Println(alexAccount.Withdraw(100))

	fmt.Println(alexAccount.Balance)

	status, value := alexAccount.Deposit(600)

	fmt.Println(status, value)

	nycolleAccount := accounts.CheckingAccount{Holder: "Nycolle", Balance: 200}

	fmt.Println(alexAccount)
	fmt.Println(nycolleAccount)

	statusTransfer := alexAccount.Transfer(200, &nycolleAccount)
	fmt.Println("Status transfer: ", statusTransfer)

	fmt.Println(alexAccount)
	fmt.Println(nycolleAccount)
}

func comparations() {
	alexAccount := accounts.CheckingAccount{Holder: "Alex", AgencyNumber: 589, AccountNumber: 123456, Balance: 125.5}
	alexAccount2 := accounts.CheckingAccount{Holder: "Alex", AgencyNumber: 589, AccountNumber: 123456, Balance: 125.5}

	fmt.Println(alexAccount == alexAccount2)

	nycolleAccount := accounts.CheckingAccount{"Nycolle", 222, 111222, 200.0}
	fmt.Println(nycolleAccount)

	var betoAccount *accounts.CheckingAccount
	betoAccount = new(accounts.CheckingAccount)
	betoAccount.Holder = "Beto"
	betoAccount.Balance = 300.5

	var betoAccount2 *accounts.CheckingAccount
	betoAccount2 = new(accounts.CheckingAccount)
	betoAccount2.Holder = "Beto"
	betoAccount2.Balance = 300.5

	fmt.Println(&betoAccount)
	fmt.Println(&betoAccount2)
	fmt.Println(betoAccount2 == betoAccount)
	fmt.Println(*betoAccount == *betoAccount2)
}
