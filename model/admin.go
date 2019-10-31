package model

/*
 * @Author: hiwein.lucus
 * @Date: 2019-10-12 14:10:54
 * @Last Modified by: hiwein.lucus
 * @Last Modified time: 2019-10-12 18:19:10
 */

import (
	"app/bearer"
	"app/db"
	"app/resp"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//Admin 管理员
type Admin struct {
	ID        int      `gorm:"primary_key;AUTO_INCREMENT;NOT NULL"`
	Username  string   `gorm:"type:varchar(64);unique_index;NOT NULL" json:"Username"`
	Password  string   `gorm:"type:varchar(128);NOT NULL" json:"password" form:"Password"`
	UpdatedAt Datetime `gorm:"ASSOCIATION_AUTOUPDATE" json:"UpdatedAt"`
	CreatedAt Datetime `gorm:"ASSOCIATION_AUTOCREATE" json:"CreatedAt"`
}

//TableName 表名admin
func (admin Admin) TableName() string {
	return bearer.LoginAdminType
}

//自动建表
func init() {
	table := db.DB.HasTable(Admin{})
	if !table {
		db.DB.CreateTable(Admin{})
	}
}

//LoginAdmin 登录者Admin
func (admin Admin) LoginAdmin(c *gin.Context) Admin {
	sess := sessions.Default(c)
	id := sess.Get("LoginAdminID")
	if id != nil {
		db.DB.First(&admin, id.(int))
	}
	return admin
}

//SetSession 设置登录者ID
func (admin Admin) SetSession(c *gin.Context, id int) {
	sess := sessions.Default(c)
	db.DB.First(&admin, id)
	if admin.ID != 0 {
		sess.Set("LoginAdminID", admin.ID)
	} else {
		sess.Set("LoginAdminID", 0)
	}
	sess.Save()
}

//LoginType 登录用户类型
func (admin Admin) LoginType() string {
	return "Admin"
}

//GetID 登录用户ID
func (admin Admin) GetID() int {
	return admin.ID
}

//GetUsername 登录用户账户
func (admin Admin) GetUsername() string {
	return admin.Username
}

//FirstByID 根据ID查找记录
func (admin Admin) FirstByID(id int) Admin {
	db.DB.First(&admin, id)
	return admin
}

//Create 新增Admin
func (admin Admin) Create(c *gin.Context) {
	err := c.Bind(&admin)
	if err != nil {
		resp.Fail(c, err.Error())
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(hash)
	b := db.DB.Create(&admin)
	for _, e := range b.GetErrors() {
		resp.Fail(c, e.Error())
		return
	}
	resp.Succ(c, admin)
}

//UpdatePass 更新密码
func (admin Admin) UpdatePass(c *gin.Context) {
	loginAdmin := admin.LoginAdmin(c)
	c.Bind(&admin)
	hash, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	b := db.DB.Model(&loginAdmin).Update("password", string(hash))
	for _, e := range b.GetErrors() {
		resp.Fail(c, e.Error())
		return
	}
	resp.Succ(c, "修改成功")
}

//Login 登录
func (admin Admin) Login(c *gin.Context) {
	var loginUser Admin
	if e := c.Bind(&loginUser); e != nil {
		resp.Fail(c, e.Error())
	}
	var o Admin
	db.DB.First(&o, "username = ?", loginUser.Username)
	if o.ID == 0 {
		resp.LoginFail(c, "未找到此用户")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(loginUser.Password)); err != nil {
		resp.LoginFail(c, "密码错误")
		return
	}
	if token, err := bearer.CreateJWT(&o); err != nil {
		resp.AuthFail(c, "授权失败"+token)
	} else {
		resp.Succ(c, token)
	}
}
