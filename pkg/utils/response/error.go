package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	Duplicate                   Error
	NotFound                    Error
	RouteNotFound               Error
	UnprocessableEntity         Error
	Unauthorized                Error
	BadRequest                  Error
	Validation                  Error
	InternalServerError         Error
	EmailOrPasswordIncorrect    Error
	UsernameOrPasswordIncorrect Error
	ConvertionNotFound          Error
	NotEnoughStock              Error
}

var ErrorConstant errorConstant = errorConstant{
	Duplicate: Error{
		Response: errorResponse{
			Message: "Created value already exists",
			Error:   E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	EmailOrPasswordIncorrect: Error{
		Response: errorResponse{
			Message: "Email or password is incorrect",
			Error:   E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	UsernameOrPasswordIncorrect: Error{
		Response: errorResponse{
			Message: "Username or password is incorrect",
			Error:   E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	NotFound: Error{
		Response: errorResponse{
			Message: "Data not found",
			Error:   E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Message: "Route not found",
			Error:   E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Message: "Invalid parameters or payload",
			Error:   E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Message: "Unauthorized, please login",
			Error:   E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Message: "Bad Request",
			Error:   E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	Validation: Error{
		Response: errorResponse{
			Message: "Invalid parameters or payload",
			Error:   E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Message: "Something bad happened",
			Error:   E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res *Error, message error) *Error {
	res.ErrorMessage = message
	return res
}

func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Message: message,
			Error:   err,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	err := json.NewEncoder(w).Encode(e.Response)
	if err != nil {
		logrus.Errorf("Error encoding JSON response: %v", err)
	}
}
