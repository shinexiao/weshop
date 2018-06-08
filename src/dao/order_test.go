package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pojo/model"
	"os"
)

// 单元测试，测试OrderDao对象的构造
func TestNewOrderDao(t *testing.T) {

	var db *gorm.DB

	orderDao := NewOrderDao(nil)

	assert.Equal(t, db, orderDao.db)
}

// 单元测试，测试保存Order对象
func TestOrderDao_InsertOrder(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")

	// 创建数据库表
	db.CreateTable(model.Order{})


	//单元测试， 覆盖正确逻辑
	order := &model.Order{
		OrderNo: "ORDER123456789",
		UserId: 46,
		UserName: "fzh",
		UserPhone: "18208160972",
		Price: float32(200.0),
	}

	orderDao := NewOrderDao(db)

	result, err := orderDao.InsertOrder(order)

	assert.Equal(t, err, nil)
	assert.Equal(t, result, int64(1))

	// 单元测试，覆盖错误逻辑
	order = &model.Order{
		OrderNo: "ORDER123456789",
		UserId: 46,
		UserName: "fzh",
		UserPhone: "18208160972",
		Price: float32(100.0),
	}

	_, err = orderDao.InsertOrder(order)
	assert.NotEqual(t, err, nil)

	db.Close()

	// 清理数据库，以备下次使用
	os.Remove("unit_test_gorm.db")

}
