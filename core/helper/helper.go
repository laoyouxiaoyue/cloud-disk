package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"log/slog"
	"math/rand"
	"time"
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

func MailSendCode(mail string, code string) error {
	slog.Info(fmt.Sprintf("Mail send code  %s t %s", code, mail))
	return nil
}

func RandCode() string {
	s := "123456789"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += fmt.Sprintf("%c", s[rand.Int()%9])
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}
