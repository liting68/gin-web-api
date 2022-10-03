package model

import (
	"app/config"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Bearer struct {
	Admin  string
	User   string
	Issuer string
}

var BEARER *Bearer

// 登录类型
func init() {
	BEARER = &Bearer{Admin: "admin", User: "user", Issuer: "Issuer"}
}

// LoginUser 登录者
type LoginUser interface {
	GetID() int
	GetUsername() string
	LoginType() string
}

// CreateJWT 创建JsonWebToken 过期时间2小时
func (b Bearer) CreateJWT(user LoginUser) (string, error) {
	claims := jwt.RegisteredClaims{
		Audience:  []string{user.GetUsername()},                      // 受众
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 失效时间2小时
		ID:        strconv.Itoa(user.GetID()),                        // 编号
		IssuedAt:  jwt.NewNumericDate(time.Now()),                    // 签发时间
		Issuer:    BEARER.Issuer,                                     // 签发人
		NotBefore: jwt.NewNumericDate(time.Now()),                    // 生效时间
		Subject:   user.LoginType(),                                  // 主题
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
