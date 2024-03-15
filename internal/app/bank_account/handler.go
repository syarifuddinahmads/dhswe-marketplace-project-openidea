package bank_account

import (
	"fmt"
	"net/http"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/response"
)

func (s service) CreateBank(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var payload dto.CreateBankParams
	err := utils.Decode(r, &payload)
	if err != nil {
		fmt.Println("Bank Account 1")
		response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(w)
		return
	}

	// Call your service to handle login
	data, err := s.BankAccountService.CreateBank(r.Context(), &payload)
	if err != nil {
		fmt.Println("Bank Account 2")
		response.ErrorResponse(err).Send(w)
		return
	}

	// Send success response
	response.SuccessResponse(data).Send(w)
}

// func (s service) Register(w http.ResponseWriter, r *http.Request) {
// 	// Parse request body
// 	var payload dto.AuthRegisterRequest
// 	err := utils.Decode(r, &payload)
// 	if err != nil {
// 		response.ErrorBuilder(&response.ErrorConstant.BadRequest, err).Send(w)
// 		return
// 	}

// 	// Call your service to handle login
// 	data, err := s.userService.Register(r.Context(), &payload)
// 	if err != nil {
// 		response.ErrorResponse(err).Send(w)
// 		return
// 	}

// 	// Send success response
// 	response.SuccessResponse(data).Send(w)
// }
