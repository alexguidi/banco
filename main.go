package main

import "fmt"

type CheckingAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	alexAccount := CheckingAccount{holder: "Alex", agencyNumber: 589, accountNumber: 123456, balance: 125.5}
	nycolleAccount := CheckingAccount{"Nycolle", 222, 111222, 200.0}
	fmt.Println(alexAccount)
	fmt.Println(nycolleAccount)
}
