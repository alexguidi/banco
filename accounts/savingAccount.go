package accounts

import "banco/customers"

type SavingAccount struct {
	Holder                                  customers.Holder
	AgencyNumber, AccountNumnber, Operation int
	balance                                 float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.balance

	if canWithdraw {
		c.balance -= withdrawValue
		return "Withdraw performed successfully"
	} else {
		return "There is no balance available to withdraw"
	}
}

func (c *SavingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.balance += value
		return "Deposit performed successfully", c.balance
	} else {
		return "Failure to deposit", c.balance
	}
}

func (c *SavingAccount) GetBalance() float64 {
	return c.balance
}
