package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       int64
	Identity string
	Name     string
	jwt.RegisteredClaims
}

var JwtKey = "cloud-disk-key"

// CodeLength 验证码长度
var CodeLength = 4

// CodeExpire 验证码过期时间
var CodeExpire = 300

// TencentSECRETID 腾讯云对象存储
var TencentSECRETID = ""
var TencentSECRETKEY = ""
var TencentURL = ""

// PageSize 分页默认参数
var PageSize = 20
