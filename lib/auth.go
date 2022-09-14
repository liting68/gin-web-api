package lib

/*
 * @Author: hiwein.lucus
 * @Date: 2019-10-12 17:25:17
 * @Last Modified by: hiwein.lucus
 * @Last Modified time: 2019-10-12 17:37:38
 */

import (
	"app/bearer"
	"app/resp"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// ManageAuth 管理者验证
func ManageAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := VerifyToken(c, bearer.LoginAdminType)
		if err != nil {
			resp.AuthFail(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

// UserAuth 用户端端验证
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := VerifyToken(c, bearer.LoginUserType)
		if err != nil {
			resp.AuthFail(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

// VerifyToken 验证token
func VerifyToken(context *gin.Context, sub string) (claim *jwt.StandardClaims, err error) {
	headAuth := context.Request.Header.Get("Authorization")
	if len(headAuth) == 0 {
		return nil, errors.New("需要 token")
	}
	claim, err = bearer.ParseToken(headAuth)
	if err != nil {
		return nil, err
	}
	if time.Now().Unix() > claim.ExpiresAt {
		return nil, errors.New("token 过期")
	}
	if claim.Subject != sub {
		return nil, errors.New("token 权限不合法")
	}
	return claim, nil
}
