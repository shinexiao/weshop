package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"pojo/model"
)

// 单元测试，测试newsDao对象的构造
func TestNewNewsDao(t *testing.T) {
	var db *gorm.DB
	newsDao := NewNewsDao(nil)
	assert.Equal(t, db, newsDao.db)
}

func TestNewsDao_InsertNews(t *testing.T) {
	// 连接 数据库
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")

	//创建 数据库
	db.CreateTable(model.News{})

	// 创建 一个新闻对象
	news := &model.News{
		Title:   "新闻表头",
		Content: "新闻内容",
		Img:     "新闻图片",
	}

	// 构造 新闻数据库对象
	newsDao := NewNewsDao(db)

	// 新闻 保存事件
	result, err := newsDao.InsertNews(news)

	// 关闭 数据库
	db.Close()

	// 断言 错误信息
	assert.Equal(t, err, nil)

	// 断言 正确信息
	assert.Equal(t, result, int64(1))
}
