package service

import (
	"os"
	"testing"

	"dao"
	"pojo/form"
	"pojo/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

// 单元测试，管理员服务层构造
func TestNewAdminService(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")

	defer os.Remove("unit_test_gorm.db")
	defer db.Close()

	db.CreateTable(model.Admin{})

	adminDao := dao.NewAdminDao(db)

	assert.NotNil(t, adminDao)

	adminService := NewAdminService(adminDao)

	assert.NotNil(t, adminService)
	assert.Equal(t, adminService.dao, adminDao)
}

// 单元测试，管理员服务层保存管理信息
func TestAdminService_InsertAdmin(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	defer os.Remove("unit_test_gorm.db")
	defer db.Close()

	db.CreateTable(model.Admin{})

	adminDao := dao.NewAdminDao(db)

	adminService := NewAdminService(adminDao)

	assert.Equal(t, adminService.dao, adminDao)

	adminForm := &form.AdminForm{
		Username: "xiaoka",
		Password: "xiaoka256",
	}

	err := adminService.InsertAdmin(adminForm)
	assert.Nil(t, err)


	//测试用户名已经存在
	err = adminService.InsertAdmin(adminForm)
	assert.NotNil(t, err)
}
