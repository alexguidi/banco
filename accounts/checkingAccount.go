package accounts

type CheckingAccount struct {
	Holder        string
	AgencyNumber  int
	AccountNumber int
	Balance       float64
}

func (c *CheckingAccount) Transfer(value float64, targetAccount *CheckingAccount) bool {
	if value < c.Balance && value > 0 {
		c.Balance -= value
		targetAccount.Deposit(value)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue > 0 && withdrawValue <= c.Balance

	if canWithdraw {
		c.Balance -= withdrawValue
		return "Withdraw performed successfully"
	} else {
		return "There is no Balance available to withdraw"
	}
}

func (c *CheckingAccount) Deposit(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "Deposit performed successfully", c.Balance
	} else {
		return "Failure to deposit", c.Balance
	}
}
