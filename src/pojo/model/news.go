package model

import (
	"github.com/jinzhu/gorm"
)

// 新闻数据访库问对象
type News struct {
	gorm.Model
	// 标题
	Title string `gorm:"column:title;type:varchar(30);not null" json:"title"`
	// 内容
	Content string `gorm:"column:content;type:varchar(500);not null" json:"content"`
	// 图片
	Img string `gorm:"column:img;type:varchar(100);not null" json:"img"`
}

// 定义新闻表名
func (News) TableName() string {

	return "shop_news"
}
