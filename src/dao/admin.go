package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
)

// 管理员数据库访问对象
type AdminDao struct {
	db *gorm.DB
}

// 构造AdminDao需要传入数据库对象
func NewAdminDao(db *gorm.DB) *AdminDao {

	return &AdminDao{
		db: db,
	}
}

// 保存管理员信息
func (this *AdminDao) InsertAdmin(admin *model.Admin) (int64, error) {

	result := this.db.Save(admin)

	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
