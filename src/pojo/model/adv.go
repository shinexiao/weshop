package model

import "github.com/jinzhu/gorm"

// 广告数据库访问
type Adv struct {
	gorm.Model
	// 图片
	Img string `gorm:"column:img;type:varchar(30);not null" json:"img"`
	// 连接
	Url string `gorm:"column:url;type:varchar(300)" json:"url"`
}

// 数据库表名
func (Adv) TableName() string {
	return "shop_adv"
}
