package model

import (
	"github.com/jinzhu/gorm"
)

// 管理员数据库访问对象
type Admin struct {
	gorm.Model
	Username string `gorm:"column:username;type:char(20);unique;not null"` // 用户名
	Password string `gorm:"column:password;type:char(30);not null"`        // 密码，密码经过sha1加密
}

// 定义admin表名
func (Admin) TableName() string {

	return "shop_admin"
}
