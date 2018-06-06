package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
)

// 新闻需要访问数据库的对象
type NewsDao struct {
	db *gorm.DB
}

// 构造NewsDao需要传入数据库的对象
func NewNewsDao(db *gorm.DB) *NewsDao {
	return &NewsDao{
		db: db,
	}
}

// 添加新闻信息
func (this *NewsDao) InsertNews(news *model.News) (int64, error) {
	// 把新闻信息插入数据库 返回一个信息
	result := this.db.Save(news)

	// 判断是否有错误信息
	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	// 成功返回 行数
	return result.RowsAffected, nil
}
