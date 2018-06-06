package dao

import "github.com/jinzhu/gorm"

// 广告数据库访问
type Adv struct {
	gorm.Model
	Img     string `gorm:"column:img;not null"` // 图片
	Url     string `gorm:"column:url"`          // 连接
	Content string `gorm:"column:content"`      // 文本
}

func (Adv) TableName() string {
	return "shop_adv"
}
