package model

import (
	"app/db"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// User 用户信息
type User struct {
	ID        int      `gorm:"primary_key;AUTO_INCREMENT;NOT NULL" json:"ID"`
	NickName  string   `gorm:"type:varchar(64);NULL;comment:'昵称'" json:"NickName"`
	Avatar    string   `gorm:"type:varchar(128);NULL;comment:'头像'" json:"Avatar"`
	Username  string   `gorm:"type:varchar(64);NOT NULL" json:"Username"`
	Password  string   `gorm:"type:varchar(128);NOT NULL" json:"Password"`
	Name      string   `gorm:"type:varchar(32);NOT NULL;comment:'姓名'" json:"Name"`
	Status    int      `gorm:"type:tinyint(1);not null;default:1;comment:'账号状态1启用2禁用'" json:"Status"`
	UpdatedAt Datetime `gorm:"ASSOCIATION_AUTOUPDATE" json:"UpdatedAt"`
	CreatedAt Datetime `gorm:"ASSOCIATION_AUTOCREATE" json:"CreatedAt"`
}

// 自动建表
func init() {
	table := db.DB.HasTable(User{})
	if !table {
		db.DB.CreateTable(User{})
		hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		new := User{Username: "user", Password: string(hash)}
		db.DB.Save(&new)
	}
}

// TableName 表名
func (u User) TableName() string {
	return "user"
}

// LoginType 登录用户类型
func (u User) LoginType() string {
	return BEARER.User
}

// GetID 登录用户ID
func (u User) GetID() int {
	return u.ID
}

// GetUsername 登录用户账户
func (u User) GetUsername() string {
	return u.Username
}

// FirstByID 根据ID查找记录
func (u User) FirstByID(id int) User {
	db.DB.First(&u, "id=?", id)
	return u
}

// FirstByID 根据用户名查找记录
func (u User) FirstByUsername(username string) User {
	db.DB.First(&u, "username=?", username)
	return u
}

// Login 登录
func (u User) Login(c *gin.Context) (any, error) {
	if e := c.Bind(&u); e != nil {
		return nil, e
	}
	var user User
	db.DB.First(&user, "username = ? ", u.Username)
	if user.ID == 0 {
		return nil, fmt.Errorf("未找到此用户")
	}
	if user.Status != 1 {
		return nil, fmt.Errorf("此账号被禁用")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		return nil, fmt.Errorf("密码错误")
	}
	if token, err := BEARER.CreateJWT(user); err != nil {
		return nil, fmt.Errorf("授权失败" + token)
	} else {
		return token, nil
	}
}

// Info 用户信息
func (u User) Info(c *gin.Context, login User) (any, error) {
	var o User
	db.DB.First(&o, "id=?", login.ID)
	if o.ID == 0 {
		return nil, fmt.Errorf("未找到此用户")
	}
	return o, nil
}
