package lib

import (
	"app/bearer"
	"app/resp"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// ManageAuth 管理者验证
func ManageAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := VerifyToken(c, bearer.LoginAdminType)
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
		err := VerifyToken(c, bearer.LoginUserType)
		if err != nil {
			resp.AuthFail(c, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

// VerifyToken 验证token
func VerifyToken(context *gin.Context, sub string) (err error) {
	headAuth := context.Request.Header.Get("Authorization")
	if len(headAuth) == 0 {
		return errors.New("需要 token")
	}
	claim, err := bearer.ParseToken(headAuth)
	if err != nil {
		return err
	}
	if time.Now().Unix() > claim.ExpiresAt {
		return errors.New("token 过期")
	}
	if claim.Subject != sub {
		return errors.New("token 权限不合法")
	}
	return nil
}
