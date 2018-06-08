package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"pojo/model"
	"os"
)

// 单元测试，测试advDao对象的构造
func TestNewAdvDao(t *testing.T) {
	// 创建一个空的DB
	var db *gorm.DB

	// new 一个 NewAdvDao
	advDao := NewAdvDao(db)

	//断言 NewAdvDao 是否正确
	assert.Equal(t, db, advDao.db)
}

// 单元测试，测试保存adv对象
func TestAdvDao_InsertAdv(t *testing.T) {
	// 连接数据库
	db, _ := gorm.Open("sqlite3", "adv_test_gorm.db")

	// 创建数据库表
	db.CreateTable(model.Adv{})

	// 创建一个adv 实例
	adv := &model.Adv{
		Img: "图片",
		Url: "地址",
	}

	// 新建一个advDao
	advDao := NewAdvDao(db)

	// 把adv实例 保存入数据库
	result, err := advDao.InsertAdv(adv)

	// 关闭数据库
	db.Close()

	// 断言 错误信息
	assert.Equal(t, err, nil)

	// 断言 正确信息
	assert.Equal(t, result, int64(1))

	os.Remove("adv_test_gorm.db")
}
