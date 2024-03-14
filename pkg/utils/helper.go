package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ResponseFormat struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    Pagination  `json:"meta"`
}

func Respond(w http.ResponseWriter, data interface{}, status int) {
	var respData interface{}
	var message string

	switch v := data.(type) {
	case nil:
		message = "success"
	case ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
		message = "error"
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
		message = "error"
	default:
		respData = data
		message = "success"
	}

	response := ResponseFormat{
		Message: message,
		Data:    respData,
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if data != nil {
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			http.Error(w, "Could not encode in JSON", http.StatusInternalServerError)
			return
		}
	}
}

func Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ReadRequestBody(r *http.Request) ([]byte, error) {
	var bodyBytes []byte
	var err error
	if r.Body != nil {
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			err := errors.New("could not read request body")
			return nil, err
		}
	}
	return bodyBytes, nil
}

func RestoreRequestBody(r *http.Request, bodyBytes []byte) {
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
}
