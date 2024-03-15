package dto

type BankAccountRequest struct {
	Bank_Name           string `valid:"required"`
	Bank_Account_Name   string `valid:"required"`
	Bank_Account_Number string `valid:"required"`
}

type BankAccountResponse struct {
	Bank_Name           string `json:"bank_name"`
	Bank_Account_Name   string `json:"bank_account_name"`
	Bank_Account_Number string `json:"bank_account_number"`
}
