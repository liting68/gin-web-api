package route

import (
	"app/config"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Auth 鉴权验证
func Auth(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := VerifyToken(c, role)
		if err != nil {
			AuthFail(c, err.Error())
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
	claim, err := ParseToken(headAuth)
	if err != nil {
		return err
	}
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		return errors.New("token 过期")
	}
	if claim.Subject != sub {
		return errors.New("token 权限不合法")
	}
	return nil
}

// ParseToken 解析token
func ParseToken(headAuth string) (*jwt.RegisteredClaims, error) {
	arr := strings.Fields(headAuth)
	if len(arr) > 1 {
		headAuth = arr[1]
	} else {
		headAuth = arr[0]
	}
	jwtToken, err := jwt.ParseWithClaims(headAuth, &jwt.RegisteredClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Info.Middleware.Secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, errors.New("token 错误")
}
