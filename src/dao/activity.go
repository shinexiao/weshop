package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
)

//活动管理访问对象
type ActivityDao struct {
	db *gorm.DB
}

//构造一个对象传入
func NewActivityDao(db *gorm.DB) *ActivityDao {
	return &ActivityDao{
		db: db,
	}
}

func (this *ActivityDao) InsterActivity(active *model.Activity)(int64,error) {
	db := this.db.Save(active)
	return db.RowsAffected, db.Error
}
