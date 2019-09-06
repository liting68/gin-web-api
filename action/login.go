package action

import (
	"app/middleware"
	"app/model"
	"app/result"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Login 登录
func Login(c *gin.Context) {
	var admin model.Admin
	username := c.PostForm("username")
	password := c.PostForm("password")
	admin = admin.FirstByUsername(username)
	if admin.ID == 0 {
		c.JSON(http.StatusOK, gin.H{"code": result.LoginErr, "errMsg": "未找到此用户"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": result.LoginPassErr, "errMsg": "密码错误"})
		return
	}
	if token, err := middleware.CreateJWT(&admin); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": result.AuthErr, "errMsg": "授权失败" + token})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": result.OK, "data": token})
	}

}
