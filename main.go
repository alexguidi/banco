package main

import (
	"banco/accounts"
	"banco/customers"
	"fmt"
)

func main() {

	exampleAccount := accounts.CheckingAccount{}
	exampleAccount.Deposit(100)
	fmt.Println(exampleAccount.GetBalance())

	customerBruno := customers.Holder{"Bruno", "35809690807", "Musician"}
	brunoAccount := accounts.CheckingAccount{Holder: customerBruno, AgencyNumber: 123, AccountNumber: 123456}
	brunoAccount.Deposit(100)
	fmt.Println(brunoAccount)

	customerAlex := customers.Holder{"Alex", "123456", "Developer"}
	alexAccount := accounts.CheckingAccount{}
	alexAccount.Holder = customerAlex
	alexAccount.Deposit(500)

	fmt.Println(alexAccount.GetBalance())

	fmt.Println(alexAccount.Withdraw(100))

	fmt.Println(alexAccount.GetBalance())

	status, value := alexAccount.Deposit(600)

	fmt.Println(status, value)

	customerNycolle := customers.Holder{"Nycolle", "1234567", "Scrum Master"}
	nycolleAccount := accounts.CheckingAccount{Holder: customerNycolle}
	nycolleAccount.Deposit(200)

	fmt.Println(alexAccount)
	fmt.Println(nycolleAccount)

	statusTransfer := alexAccount.Transfer(200, &nycolleAccount)
	fmt.Println("Status transfer: ", statusTransfer)

	fmt.Println(alexAccount)
	fmt.Println(nycolleAccount)
}

func comparations() {
	customerAlex := customers.Holder{"Alex", "123123123", "Developer"}
	alexAccount := accounts.CheckingAccount{Holder: customerAlex, AgencyNumber: 589, AccountNumber: 123456}
	alexAccount.Deposit(125.5)
	alexAccount2 := accounts.CheckingAccount{Holder: customerAlex, AgencyNumber: 589, AccountNumber: 123456}
	alexAccount2.Deposit(125.5)

	fmt.Println(alexAccount == alexAccount2)

	customerNycolle := customers.Holder{"Nycolle", "123123123", "Scrum Master"}
	nycolleAccount := accounts.CheckingAccount{Holder: customerNycolle, AgencyNumber: 222, AccountNumber: 111222}
	nycolleAccount.Deposit(200)
	fmt.Println(nycolleAccount)

	betoCustomer := customers.Holder{"Beto", "32123", "Retired"}
	var betoAccount *accounts.CheckingAccount
	betoAccount = new(accounts.CheckingAccount)
	betoAccount.Holder = betoCustomer
	betoAccount.Deposit(300.5)

	betoCustomer2 := customers.Holder{"Beto", "32123", "Retired"}
	var betoAccount2 *accounts.CheckingAccount
	betoAccount2 = new(accounts.CheckingAccount)
	betoAccount2.Holder = betoCustomer2
	betoAccount2.Deposit(300.5)

	fmt.Println(&betoAccount)
	fmt.Println(&betoAccount2)
	fmt.Println(betoAccount2 == betoAccount)
	fmt.Println(*betoAccount == *betoAccount2)
}
