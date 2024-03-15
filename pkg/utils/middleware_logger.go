package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type ResponseWriter struct {
	http.ResponseWriter
	status      int
	body        []byte
	wroteHeader bool
	wroteBody   bool
}

func WrapResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{ResponseWriter: w}
}

func (rw *ResponseWriter) Status() int {
	return rw.status
}

func (rw *ResponseWriter) WriteHeader(code int) {
	if rw.wroteBody {
		return
	}
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true
}

func (rw *ResponseWriter) Write(body []byte) (int, error) {
	if rw.wroteBody {
		return 0, nil
	}
	i, err := rw.ResponseWriter.Write(body)
	if err != nil {
		return 0, err
	}
	rw.body = body
	return i, err
}

func (rw *ResponseWriter) Body() []byte {
	return rw.body
}

func MiddlewareLogger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		requestBody, err := ReadRequestBody(r)
		if err != nil {
			Respond(w, err, 0)
			return
		}
		RestoreRequestBody(r, requestBody)

		logMessage := fmt.Sprintf("path:%s, method: %s, requestBody: %v", r.URL.EscapedPath(), r.Method, string(requestBody))

		start := time.Now()
		wrapped := WrapResponseWriter(w)
		next.ServeHTTP(wrapped, r)

		logMessage = fmt.Sprintf("%s, responseStatus: %d, responseBody: %s", logMessage, wrapped.Status(), string(wrapped.Body()))
		// Assuming s.logger is globally accessible here.
		logrus.Infof("%s, duration: %v", logMessage, time.Since(start))
	}
	return http.HandlerFunc(fn)
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

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func (e ErrorResponse) Error() string {
	return e.ErrorMessage
}

type ErrArgument struct {
	Wrapped error
}

func (e ErrArgument) Error() string {
	return "invalid argument"
}

func (e ErrArgument) Unwrap() error {
	return e.Wrapped
}
