package dao

import (
	"os"
	"testing"

	"pojo/model"

	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 单元测试，测试AdminDao对象的构造
func TestNewAdminDao(t *testing.T) {

	adminDao := NewAdminDao(nil)

	assert.Nil(t, adminDao.db)
}

// 单元测试，测试保存Admin对象
func TestAdminDao_InsertAdmin(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	defer os.Remove("unit_test_gorm.db")
	defer db.Close()

	// 创建数据库表
	db.CreateTable(model.Admin{})

	//单元测试， 覆盖正确逻辑
	admin := &model.Admin{
		Username: "xiaoka",
		Password: "xxxx111111",
	}

	adminDao := NewAdminDao(db)

	result, err := adminDao.InsertAdmin(admin)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	// 单元测试，覆盖错误逻辑
	admin = &model.Admin{
		Username: "xiaoka",
		Password: "xxx22222",
	}

	_, err = adminDao.InsertAdmin(admin)
	assert.NotNil(t, err)
}

// 单元测试，查询保存的Admin对象
func TestAdminDao_GetAdmin(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	defer os.Remove("unit_test_gorm.db")
	defer db.Close()

	// 创建数据库表
	db.CreateTable(model.Admin{})

	//单元测试， 覆盖正确逻辑
	admin := &model.Admin{
		Username: "xiaoka",
		Password: "xxxx111111",
	}

	adminDao := NewAdminDao(db)

	result, err := adminDao.InsertAdmin(admin)

	assert.Nil(t, err)
	assert.Equal(t, result, int64(1))

	local, err := adminDao.GetAdmin(admin.Username)

	assert.Nil(t, err)
	assert.Equal(t, local.Username, "xiaoka")

	_, err = adminDao.GetAdmin("notexists")
	assert.NotNil(t, err)

}
