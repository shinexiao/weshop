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

	// 判断是否有错误信息
	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	// 成功返回 行数
	return result.RowsAffected, nil
}

// 查询新闻信息
func (this *NewsDao) GetNewsList(news *model.News, page, limit int32) ([]*model.News, error) {
	news_list := make([] *model.News, limit)
	// 把新闻信息更新到数据库  返回一个信息
	result := this.db.Table("shop_news").Where(map[string]interface{}{"title": news.Title}).Order("created_at").Limit(limit).Offset((page - 1) * limit).Find(&news_list)

	// 判断是否有错误信息
	if result.Error != nil {
		// 保存出错了
		return nil, result.Error
	}

	// 成功返回 行数
	return news_list, nil
}

// 根据ID查询新闻
func (this *NewsDao) GetNews(id int64) (*model.News, error) {
	news := &model.News{}

	result := this.db.Table("shop_news").Where("id = ?", id).Find(news)

	if result.Error != nil {
		return nil, result.Error
	}

	return news, nil
}

// 根据id 修改对于的 新闻
func (this *NewsDao) PutNews(news *model.News, id int64) (int64, error) {
	result := this.db.Table("shop_news").Where("id = ?", id).Updates(map[string]interface{}{
		"title":   news.Title,
		"content": news.Content,
		"img":     news.Img,
	})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
