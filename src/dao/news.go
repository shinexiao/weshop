package dao

import (
	"github.com/jinzhu/gorm"
)

// 新闻数据访库问对象
type New struct {
	gorm.Model
	Title   string `gorm:"column:title"`   // 标题
	Content string `gorm:"column:content"` // 内容
	Img     string `gorm:"column:img"`     // 图片
}

// 定义新闻表名
func (New) TableName() string {

	return "shop_news"
}
