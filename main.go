package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	alexAccount := CheckingAccount{}
	alexAccount.holder = "Alex"
	alexAccount.balance = 500

	fmt.Println(alexAccount.balance)

	fmt.Println(alexAccount.withdraw(-100))

	fmt.Println(alexAccount.balance)
}

func (c *CheckingAccount) withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw performed successfully"
	} else {
		return "There is no balance available to withdraw"
	}
}

func comparations() {
	alexAccount := CheckingAccount{holder: "Alex", agencyNumber: 589, accountNumber: 123456, balance: 125.5}
	alexAccount2 := CheckingAccount{holder: "Alex", agencyNumber: 589, accountNumber: 123456, balance: 125.5}

	fmt.Println(alexAccount == alexAccount2)

	nycolleAccount := CheckingAccount{"Nycolle", 222, 111222, 200.0}
	fmt.Println(nycolleAccount)

	var betoAccount *CheckingAccount
	betoAccount = new(CheckingAccount)
	betoAccount.holder = "Beto"
	betoAccount.balance = 300.5

	var betoAccount2 *CheckingAccount
	betoAccount2 = new(CheckingAccount)
	betoAccount2.holder = "Beto"
	betoAccount2.balance = 300.5

	fmt.Println(&betoAccount)
	fmt.Println(&betoAccount2)
	fmt.Println(betoAccount2 == betoAccount)
	fmt.Println(*betoAccount == *betoAccount2)
}
