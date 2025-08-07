package helper

import (
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"

	"path"
	"time"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func GenerateToken(id int64, identity, name string, second int64) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Name:     name,
		Identity: identity,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),
		},
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

func CosUpload(r *http.Request) (string, error) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://cloud-disk-1315228122.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.TencentSECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: define.TencentSECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go
	//name := "cloud-disk/exampleobject.jpg"
	//fd, err := os.Open("./exampleobject.jpeg")
	//if err != nil {
	//	panic(err)
	//}
	//defer fd.Close()
	//_, err = c.Object.Put(context.Background(), name, fd, nil)
	//if err != nil {
	//	panic(err)
	//}
	file, fileHeader, err := r.FormFile("file")
	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)

	_, err = c.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return key, nil
}

func AnalyzeToken(token string) (*define.UserClaim, error) {
	uc := &define.UserClaim{}
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("TokenInvalidErr")
	}
	return uc, nil
}
