package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pojo/model"
	"os"
)
func TestNewActiveDao(t *testing.T) {
	var db *gorm.DB

	activety := NewAdminDao(nil)

	assert.Equal(t, db, activety.db)
}

func TestActivetyDao_InsterActivity(t *testing.T) {

	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	//创建表
	db.CreateTable(model.Activity{})

	activity := &model.Activity{
		Type:1,
		NewCash:12.0,
		OldCash:11.0,
	}
	dao := NewActivityDao(db)
	result ,err:= dao.InsterActivity(activity)
	assert.Equal(t,err,nil)
	assert.Equal(t,result,int64(1))
	os.Remove("unit_test_gorm.db")
}
