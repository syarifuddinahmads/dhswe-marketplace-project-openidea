package dto

type CreateBankParams struct {
	BankName          string `valid:"required"`
	BankAccountName   string `valid:"required"`
	BankAccountNumber string `valid:"required"`
}

type BankAccountResponse struct {
	BankName          string `json:"bank_name"`
	BankAccountName   string `json:"bank_account_name"`
	BankAccountNumber string `json:"bank_account_number"`
}
