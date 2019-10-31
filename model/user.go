package model

import (
	"app/bearer"
	"app/db"
	"app/resp"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//User 用户信息
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

//自动建表
func init() {
	table := db.DB.HasTable(User{})
	if !table {
		db.DB.CreateTable(User{})
	}
}

//TableName 表名
func (u User) TableName() string {
	return bearer.LoginUserType
}

//LoginDoctor 登录者Doctor
func (u User) LoginDoctor(c *gin.Context) User {
	sess := sessions.Default(c)
	id := sess.Get("LoginDoctorID")
	if id != nil {
		db.DB.First(&u, id.(int))
	}
	return u
}

//SetSession 设置登录者ID
func (u User) SetSession(c *gin.Context, id int) {
	sess := sessions.Default(c)
	db.DB.First(&u, id)
	if u.ID != 0 {
		sess.Set("LoginDoctorID", u.ID)
	} else {
		sess.Set("LoginDoctorID", 0)
	}
	sess.Save()
}

//LoginType 登录用户类型
func (u User) LoginType() string {
	return "User"
}

//GetID 登录用户ID
func (u User) GetID() int {
	return u.ID
}

//GetUsername 登录用户账户
func (u User) GetUsername() string {
	return u.Username
}

//FirstByID 根据ID查找记录
func (u User) FirstByID(id int) User {
	db.DB.First(&u, id)
	return u
}

//Login 登录
func (u User) Login(c *gin.Context) {
	if e := c.Bind(&u); e != nil {
		resp.Fail(c, e.Error())
	}
	var user User
	db.DB.First(&user, "username = ? ", u.Username)
	if user.ID == 0 {
		resp.LoginFail(c, "未找到此用户")
		return
	}
	if user.Status != 1 {
		resp.LoginFail(c, "此账号被禁用")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		resp.LoginPassFail(c, "密码错误")
		return
	}
	if token, err := bearer.CreateJWT(user); err != nil {
		resp.AuthFail(c, "授权失败"+token)
	} else {
		resp.Succ(c, token)
	}
}
