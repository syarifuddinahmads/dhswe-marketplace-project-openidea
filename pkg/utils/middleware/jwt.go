package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/constant"
)

var (
	jwtKey = []byte(constant.JWT_SECRET)
)

func JwtMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Missing authorization header")
				return
			}
			tokenString = tokenString[len("Bearer "):]

			err := VerifyToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid token")
				return
			}

			// Jika token valid, lanjutkan ke handler berikutnya
			next.ServeHTTP(w, r)
		})
	}
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
