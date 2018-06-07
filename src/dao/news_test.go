package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"pojo/model"
	"fmt"
	"runtime/debug"
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

	//创建 数据库表
	//db.CreateTable(model.News{})

	// 创建 一个新闻对象
	news := make([] *model.News, 2)
	new1 := &model.News{
		Title:   "成都新闻",
		Content: "成都新闻内容",
		Img:     "成都新闻图片",
	}
	new2 := &model.News{
		Title:   "北京新闻",
		Content: "北京新闻内容",
		Img:     "北京新闻图片",
	}

	news[0] = new1
	news[1] = new2

	for key, val := range news {
		fmt.Print(key, val)
	}

	// 构造 新闻数据库对象
	newsDao := NewNewsDao(db)

	// 新闻 保存事件
	result, err := newsDao.InsertNews(new1)
	result, err = newsDao.InsertNews(new2)

	// 关闭 数据库
	db.Close()

	// 断言 错误信息
	assert.Equal(t, err, nil)

	// 断言 正确信息
	assert.Equal(t, result, int64(1))

	//os.Remove("news_test_gorm.db")
}

func TestNewsDao_FirstNews(t *testing.T) {
	defer func()(){
		if err := recover();nil != err {
			debug.PrintStack()
		}
	}()
	// 连接 数据库
	db, _ := gorm.Open("sqlite3", "news_test_gorm.db")

	//创建 数据库表
	//db.CreateTable(model.News{})

	// 创建 一个新闻对象
	var news []model.News

	// 构造 新闻数据库对象
	//newsDao := NewNewsDao(db)
	//
	//// 新闻 保存事件
	//result, err := newsDao.InsertNews(news)
	//
	//// 创建 一个新闻对象
	//news = &model.News{
	//	Title:   "新闻表头2",
	//	Content: "新闻内容2",
	//	Img:     "新闻图片2",
	//}
	//
	//// 新闻 保存事件
	//result, err = newsDao.InsertNews(news)

	db.Limit(10).Offset(0).Find(&news)

	fmt.Print(news)
	// 关闭 数据库
	defer db.Close()

	// 断言 错误信息
	//assert.Equal(t, err, nil)
	//
	//// 断言 正确信息
	//assert.Equal(t, result, int64(1))

	//os.Remove("news_test_gorm.db")
}
