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
	result := this.db.Create(&news)

	this.db.Table().Where("id = ?", id).Update(
	{
		map[string]interface{}{
			"title", news.Title,
		}
	})

	// 判断是否有错误信息
	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	// 成功返回 行数
	return result.RowsAffected, nil
}

// 查询新闻信息
func (this *NewsDao) FirstNews(news *model.News, page, limit int32) (int64, error) {
	// 把新闻信息更新到数据库  返回一个信息
	result := this.db.Table("news").Where("title like %?%", news.Title).Order("title DESC").Limit(20).Offset((page - 1) * limit).Find(news)

	// 判断是否有错误信息
	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	// 成功返回 行数
	return result.RowsAffected, nil
}
