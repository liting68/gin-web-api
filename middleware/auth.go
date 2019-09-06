package middleware

import (
	"app/config"
	"app/model"
	"app/result"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//LoginAdmin 登录者信息
var LoginAdmin model.Admin

//CreateJWT 创建JsonWebToken 过期时间10小时
func CreateJWT(user *model.Admin) (string, error) {
	expiresTime := time.Now().Unix() + 10*3600 //10小时
	claims := jwt.StandardClaims{
		Audience:  user.Username,         // 受众
		ExpiresAt: expiresTime,           // 失效时间
		Id:        strconv.Itoa(user.ID), // 编号
		IssuedAt:  time.Now().Unix(),     // 签发时间
		Issuer:    "signer",              // 签发人
		NotBefore: time.Now().Unix(),     // 生效时间
		Subject:   "login",               // 主题
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
			context.JSON(http.StatusUnauthorized, gin.H{"code": result.AuthErr, "errMsg": "需要 token"})
		}
		auth = strings.Fields(auth)[1]
		// 校验token
		claims, err := parseToken(auth)
		if err != nil || time.Now().Unix() > claims.ExpiresAt {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{"code": result.AuthErr, "errMsg": "token 过期"})
		}
		fmt.Printf("%#v\n", claims)
		LoginID, _ := strconv.Atoi(claims.Id)
		fmt.Printf("%#v\n", LoginID)
		LoginAdmin = LoginAdmin.FirstByID(int64(LoginID))
		fmt.Printf("%#v\n", LoginAdmin)
		if LoginAdmin.ID == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, gin.H{"code": result.AuthErr, "errMsg": "用户不合法"})
		}
		context.Next()
	}
}
func parseToken(token string) (*jwt.StandardClaims, error) {
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
