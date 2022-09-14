package bearer

import (
	"app/config"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 登录类型
const (
	LoginAdminType = "Admin"
	LoginUserType  = "User"
)

// LoginUser 登录者
type LoginUser interface {
	GetID() int
	GetUsername() string
	LoginType() string
}

// CreateJWT 创建JsonWebToken 过期时间10小时
func CreateJWT(user LoginUser) (string, error) {
	expiresTime := time.Now().Unix() + int64(10*time.Hour.Seconds()) //10小时
	claims := jwt.StandardClaims{
		Audience:  user.GetUsername(),         // 受众
		ExpiresAt: expiresTime,                // 失效时间
		Id:        strconv.Itoa(user.GetID()), // 编号
		IssuedAt:  time.Now().Unix(),          // 签发时间
		Issuer:    "gin-web-api",              // 签发人
		NotBefore: time.Now().Unix(),          // 生效时间
		Subject:   user.LoginType(),           // 主题
	}
	var jwtSecret = []byte(config.Info.Middleware.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	token := ""
	if token, err = tokenClaims.SignedString(jwtSecret); err == nil {
		return "Bearer " + token, err
	}
	return "", err
}

// ParseToken 解析token
func ParseToken(headAuth string) (*jwt.StandardClaims, error) {
	arr := strings.Fields(headAuth)
	if len(arr) > 1 {
		headAuth = arr[1]
	} else {
		headAuth = arr[0]
	}
	jwtToken, err := jwt.ParseWithClaims(headAuth, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Info.Middleware.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, errors.New("token 错误")
}
