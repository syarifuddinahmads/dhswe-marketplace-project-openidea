package model

type BankAccount struct {
	BankId            int    `json:"bank_id"`
	BankName          string `json:"bank_name"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
}
