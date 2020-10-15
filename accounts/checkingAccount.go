package accounts

import "banco/customers"

type CheckingAccount struct {
	Holder                      customers.Holder
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *CheckingAccount) Transfer(value float64, targetAccount *CheckingAccount) bool {
	if value < c.balance && value > 0 {
		c.balance -= value
		targetAccount.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw performed successfully"
	} else {
		return "There is no balance available to withdraw"
	}
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit performed successfully", c.balance
	} else {
		return "Failure to deposit", c.balance
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
