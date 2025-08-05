package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func GenerateToken(id int64, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Name:     name,
		Identity: identity,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", errors.New("GenerateTokenErr")
	}
	return tokenString, nil
}
