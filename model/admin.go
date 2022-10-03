package model

import (
	"app/db"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Admin 管理员
type Admin struct {
	ID        int      `gorm:"primary_key;AUTO_INCREMENT;NOT NULL"`
	Username  string   `gorm:"type:varchar(64);unique_index;NOT NULL" json:"Username"`
	Password  string   `gorm:"type:varchar(128);NOT NULL" json:"Password"`
	UpdatedAt Datetime `gorm:"ASSOCIATION_AUTOUPDATE" json:"UpdatedAt"`
	CreatedAt Datetime `gorm:"ASSOCIATION_AUTOCREATE" json:"CreatedAt"`
}

// TableName 表名admin
func (admin Admin) TableName() string {
	return "admin"
}

// 自动建表
func init() {
	table := db.DB.HasTable(Admin{})
	if !table {
		db.DB.CreateTable(Admin{})
		hash, _ := bcrypt.GenerateFromPassword([]byte("admin@1234"), bcrypt.DefaultCost)
		new := Admin{Username: "admin", Password: string(hash)}
		db.DB.Save(&new)
	}
}

// LoginType 登录用户类型
func (admin Admin) LoginType() string {
	return BEARER.Admin
}

// GetID 登录用户ID
func (admin Admin) GetID() int {
	return admin.ID
}

// GetUsername 登录用户账户
func (admin Admin) GetUsername() string {
	return admin.Username
}

// FirstByID 根据ID查找记录
func (admin Admin) FirstByID(id int) Admin {
	db.DB.First(&admin, "id=?", id)
	return admin
}

// FirstByUsername 根据用户名查找记录
func (admin Admin) FirstByUsername(username string) Admin {
	db.DB.First(&admin, "username=?", username)
	return admin
}

// Create 新增Admin
func (admin Admin) Create(c *gin.Context) (any, error) {
	err := c.Bind(&admin)
	if err != nil {
		return nil, err
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(hash)
	b := db.DB.Create(&admin)
	for _, err = range b.GetErrors() {
		return nil, err
	}
	return admin, nil
}

// UpdatePass 更新密码
func (admin Admin) UpdatePass(c *gin.Context, loginAdmin Admin) (any, error) {
	c.Bind(&admin)
	hash, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	b := db.DB.Model(&loginAdmin).Update("password", string(hash))
	for _, e := range b.GetErrors() {
		return nil, e
	}
	return "修改成功", nil
}

// ModifyPass 修改密码
func (admin Admin) ModifyPass(c *gin.Context, login Admin) (any, error) {
	type requestPass struct {
		OldPass string `json:"OldPass"`
		NewPass string `json:"NewPass"`
	}
	var req requestPass
	c.Bind(&req)
	user := admin.FirstByID(login.ID)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPass)); err != nil {
		return nil, fmt.Errorf("旧密码错误")
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(req.NewPass), bcrypt.DefaultCost)
	user.Password = string(hash)
	if err := db.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return "密码修改成功", nil
}

// Login 登录
func (admin Admin) Login(c *gin.Context) (any, error) {
	var loginUser Admin
	if e := c.Bind(&loginUser); e != nil {
		return nil, e
	}
	var o Admin
	db.DB.First(&o, "username = ?", loginUser.Username)
	if o.ID == 0 {
		return nil, fmt.Errorf("未找到此用户")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(o.Password), []byte(loginUser.Password)); err != nil {
		return nil, fmt.Errorf("密码错误")
	}
	if token, err := BEARER.CreateJWT(&o); err != nil {
		return nil, fmt.Errorf("授权失败" + token)
	} else {
		log.Printf("%+v \n", o)
		return token, nil
	}
}
