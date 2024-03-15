package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/syarifuddinahmads/dhswe-marketplace-project-openidea/pkg/utils/constant"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId   int    `json:"user_id"`
	Name     string `json:"name;size:200;not null"`
	Username string `json:"username;size:200;not null;unique"`
	Password string `json:"password"`
}

// HashPassword is a method for struct User for Hashing password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

// GenerateToken is a method for struct User for creating new jwt token
func (u *User) GenerateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  u.UserId,
		"username": u.Username,
		"name":     u.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // we set expired in 72 hour
	})

	tokenString, err := token.SignedString([]byte(constant.JWT_SECRET))
	return tokenString, err
}
