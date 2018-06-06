package dao

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pojo/model"
	"os"
)

// 单元测试，测试AdminDao对象的构造
func TestNewAdminDao(t *testing.T) {

	var db *gorm.DB

	adminDao := NewAdminDao(nil)

	assert.Equal(t, db, adminDao.db)
}

// 单元测试，测试保存Admin对象
func TestAdminDao_InsertAdmin(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	// 创建数据库表
	db.CreateTable(model.Admin{})

	//单元测试， 覆盖正确逻辑
	admin := &model.Admin{
		Username: "xiaoka",
		Password: "xxxx111111",
	}

	adminDao := NewAdminDao(db)

	result, err := adminDao.InsertAdmin(admin)

	assert.Equal(t, err, nil)
	assert.Equal(t, result, int64(1))

	// 单元测试，覆盖错误逻辑
	admin = &model.Admin{
		Username: "xiaoka",
		Password: "xxx22222",
	}

	_ , err = adminDao.InsertAdmin(admin)
	assert.NotEqual(t, err, nil)

	db.Close()

	// 清理数据库，以备下次使用
	os.Remove("unit_test_gorm.db")

}
