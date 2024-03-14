package bank

import (
	"net/http"

	bankService "github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/bank/service"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
)

func (s service) Create() http.HandlerFunc {
	type request struct {
		BankName          string `json:"bank_name"`
		BankAccountName   string `json:"bank_account_name"`
		BankAccountNumber string `json:"bank_account_number"`
	}

	type response struct {
		BankAccountId string `json:"bank_account_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := request{}
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := utils.Decode(r, &req)
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}

		id, err := s.bankService.Create(r.Context(), bankService.CreateParams{
			BankName:          req.BankName,
			BankAccountName:   req.BankAccountName,
			BankAccountNumber: req.BankAccountNumber,
		})
		if err != nil {
			utils.Respond(w, err, 0)
			return
		}
		utils.Respond(w, response{BankAccountId: id}, http.StatusOK)
	}
}
