package action

import (
	"app/middleware"
	"app/model"
	"app/result"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//AddManagerAccount 新增管理员
func AddManagerAccount(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	var admin model.Admin
	admin.Username = username
	admin.Password = string(hash)
	id := admin.Create()
	if id == 0 {
		c.JSON(http.StatusOK, gin.H{"code": result.Err, "errMsg": "添加账号失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": result.OK, "insertID": id})
	}
}

func UpdateManagerPass(c *gin.Context) {
	password := c.PostForm("password")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	var admin = middleware.LoginAdmin
	admin.UpdatePass(string(hash))
	c.JSON(http.StatusOK, gin.H{"code": result.OK, "msg": "success"})
}
