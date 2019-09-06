package model

import (
	"app/db"
)

//Admin 管理员
type Admin struct {
	ID        int      `gorm:"primary_key;AUTO_INCREMENT;NOT NULL"`
	Username  string   `gorm:"type:varchar(100);unique_index;NOT NULL" json:"username" binding:"required"`
	Password  string   `gorm:"type:varchar(100);NOT NULL" json:"password" form:"password" binding:"required"`
	Role      int      `gorm:"type:tinyint(4);not null;default:1"`
	UpdatedAt Datetime `gorm:"ASSOCIATION_AUTOUPDATE" json:"update_at"`
	CreatedAt Datetime `gorm:"ASSOCIATION_AUTOCREATE" json:"create_at"`
}

//自动建表
func init() {
	table := db.G.HasTable(Admin{})
	if !table {
		db.G.CreateTable(Admin{})
	}
}

//TableName 表名admin
func (admin Admin) TableName() string {
	return "admin"
}

//FirstByUsername 根据用户名查找
func (admin Admin) FirstByUsername(username string) Admin {
	db.G.First(&admin, "username = ?", username)
	return admin
}

//FirstByID 根据用户ID查找记录
func (admin Admin) FirstByID(id int) Admin {
	db.G.First(&admin, id)
	return admin
}

//Create 新增Admin
func (admin *Admin) Create() int {
	db.G.Create(&admin)
	return admin.ID
}

//UpdatePass 更新密码
func (admin *Admin) UpdatePass(pass string) {
	admin.Password = pass
	db.G.Save(&admin)
}
