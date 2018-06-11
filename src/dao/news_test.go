package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"pojo/model"
	"os"
)

// 单元测试，测试newsDao对象的构造
func TestNewNewsDao(t *testing.T) {
	var db *gorm.DB
	newsDao := NewNewsDao(nil)
	assert.Equal(t, db, newsDao.db)
}

func TestNewsDao_InsertNews(t *testing.T) {
	// 连接 数据库
	db, _ := gorm.Open("sqlite3", "news_test_gorm.db")

	// 关闭 数据库
	defer db.Close()

	//创建 数据库表
	db.CreateTable(model.News{})

	newsDao := NewNewsDao(db)

	news := &model.News{
		Title:   "成都新闻",
		Content: "成都新闻内容",
		Img:     "成都新闻图片",
	}
	// 正确测试
	result, err := newsDao.InsertNews(news)

	// 断言 错误信息
	assert.Nil(t, err)

	// 断言 正确信息
	assert.Equal(t, result, int64(1))

	// 错误测试
	news = &model.News{}
	_, err = newsDao.InsertNews(news)
	assert.Nil(t, err)

	//os.Remove("news_test_gorm.db")
}

// 测试 根据id 获取单个新闻
func TestNewsDao_GetNews(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "news_test_gorm.db")

	//删除数据库表
	defer os.Remove("news_test_gorm.db")

	//关闭数据库表
	defer db.Close()

	// 创建数据库表
	db.CreateTable(model.News{})

	// 创建实例
	newsDao := NewNewsDao(db)

	// 新建一个新闻
	news := &model.News{
		Title:   "重庆新闻",
		Content: "重庆新闻内容",
		Img:     "重庆新闻图片",
	}

	// 保存一个 新闻
	result, err := newsDao.InsertNews(news)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	// 正确测试
	news, err = newsDao.GetNews(1)
	assert.Equal(t, err, nil)
	assert.Equal(t, news.Title, "重庆新闻")
	assert.Equal(t, news.Content, "重庆新闻内容")
	assert.Equal(t, news.Img, "重庆新闻图片")

	//错误测试
	news, err = newsDao.GetNews(2)
	assert.NotNil(t, err)
}

// 测试 根据ID 修改单个新闻
func TestNewsDao_PutNews(t *testing.T) {
	// 连接数据库
	db, _ := gorm.Open("sqlite3", "news_test_gorm.db")

	//删除数据库表
	defer os.Remove("news_test_gorm.db")

	//关闭数据库表
	defer db.Close()

	// 创建数据库表
	db.CreateTable(model.News{})

	// 创建实例
	newsDao := NewNewsDao(db)

	// 新建一个新闻
	news := &model.News{
		Title:   "重庆新闻",
		Content: "重庆新闻内容",
		Img:     "重庆新闻图片",
	}

	// 向数据添加一个新闻
	result, err := newsDao.InsertNews(news)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	news = &model.News{
		Title:   "北京新闻",
		Content: "北京新闻内容",
		Img:     "北京新闻图片",
	}

	// 正确测试
	result, err = newsDao.PutNews(news, 1)
	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	// 获取最新的 id 1 的新闻
	db.Table("shop_news").Where("id = ?", 1).Find(news)

	assert.Equal(t, news.Title, "北京新闻")
	assert.Equal(t, news.Content, "北京新闻内容")
	assert.Equal(t, news.Img, "北京新闻图片")

	// 错误测试
	result, err = newsDao.PutNews(news, 12)
	assert.Nil(t, err)
	assert.Equal(t, result, int64(0))
}

// 测试批量删除 新闻
func TestNewsDao_DeleteNews(t *testing.T) {
	// 连接数据库
	db, _ := gorm.Open("sqlite3", "news_test_gorm.db")

	//删除数据库表
	defer os.Remove("news_test_gorm.db")

	//关闭数据库表
	defer db.Close()

	// 创建数据库表
	db.CreateTable(model.News{})

	// 创建实例
	newsDao := NewNewsDao(db)

	// 新建一个新闻
	news := &model.News{
		Title:   "重庆新闻",
		Content: "重庆新闻内容",
		Img:     "重庆新闻图片",
	}

	// 向数据添加一个新闻
	result, err := newsDao.InsertNews(news)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	// 新建一个新闻
	news = &model.News{
		Title:   "成都新闻",
		Content: "成都新闻内容",
		Img:     "成都新闻图片",
	}

	// 向数据添加一个新闻
	result, err = newsDao.InsertNews(news)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	ids := [2] int64{1, 2}

	for i := 0; i < len(ids); i++ {
		result, err = newsDao.DeleteNews(ids[i])
		assert.Nil(t, err)
		assert.Equal(t, result, int64(1))
	}
}
