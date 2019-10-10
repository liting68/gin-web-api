package middleware

import (
	"app/config"
	"app/model"
	"app/result"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Loginner interface {
	GetID() int
	GetUsername() string
	LoginType() string
}

//Logging 登录者信息
var Logging Loginner

//CreateJWT 创建JsonWebToken 过期时间10小时
func CreateJWT(user Loginner) (string, error) {
	expiresTime := time.Now().Unix() + int64(10*time.Hour.Seconds()) //10小时
	claims := jwt.StandardClaims{
		Audience:  user.GetUsername(),         // 受众
		ExpiresAt: expiresTime,                // 失效时间
		Id:        strconv.Itoa(user.GetID()), // 编号
		IssuedAt:  time.Now().Unix(),          // 签发时间
		Issuer:    "EPR",                      // 签发人
		NotBefore: time.Now().Unix(),          // 生效时间
		Subject:   user.LoginType(),           // 主题
	}
	var jwtSecret = []byte(config.Info.Middleware.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var err error
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		return "Bearer " + token, err
	}
	return "", err
}

//Auth 验证中间件函数
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		auth := context.Request.Header.Get("Authorization")
		if len(auth) == 0 {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{"code": result.AuthErr, "errMsg": "需要 token"})
			return
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		claims, err := ParseToken(auth)
		if err != nil || claims == nil || time.Now().Unix() > claims.ExpiresAt {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{"code": result.AuthErr, "errMsg": "token 过期"})
			return
		}
		LoginID, _ := strconv.Atoi(claims.Id)
		var d model.Doctor
		var a model.Admin
		switch claims.Subject {
		case d.LoginType():
			Logging = d.FirstByID(LoginID)
		case a.LoginType():
			Logging = a.FirstByID(LoginID)
		}
		if Logging == nil || Logging.GetID() == 0 {
			context.Abort()
			context.JSON(http.StatusOK, gin.H{"code": result.AuthErr, "errMsg": "用户不合法"})
			return
		}
		context.Next()
	}
}
func ParseToken(token string) (*jwt.StandardClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Info.Middleware.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
