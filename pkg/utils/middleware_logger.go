package utils

import (
	"fmt"
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
