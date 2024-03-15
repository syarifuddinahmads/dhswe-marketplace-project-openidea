package response

import (
	"encoding/json"
	"net/http"

	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/internal/dto"
)

type successResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Success struct {
	Response successResponse `json:"response"`
	Code     int             `json:"code"`
}

var SuccessConstant = Success{
	Response: successResponse{
		Message: "Request successfully proceed",
		Data:    nil,
	},
	Code: http.StatusOK,
}

func SuccessBuilder(res *Success, data interface{}) *Success {
	res.Response.Data = data
	return res
}

func CustomSuccessBuilder(code int, data interface{}, message string, info *dto.PaginationInfo) *Success {
	return &Success{
		Response: successResponse{
			Message: message,
			Data:    data,
		},
		Code: code,
	}
}

func SuccessResponse(data interface{}) *Success {
	return SuccessBuilder(&SuccessConstant, data)
}

func (s *Success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.Code)
	json.NewEncoder(w).Encode(s.Response)
}
