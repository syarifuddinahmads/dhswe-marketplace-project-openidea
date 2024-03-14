package model

type BankAcccount struct {
	BankAccountId     string `json:"bank_account_id"`
	BankName          string `json:"bank_name"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
}
