package dao

import (
	"pojo/model"

	"github.com/jinzhu/gorm"
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

// 根据用户名查询管理员
func (this *AdminDao) GetAdmin(username string) (*model.Admin, error) {

	admin := &model.Admin{}

	result := this.db.Find(admin, "username = ?", username)

	if result.Error != nil {

		return nil, result.Error
	}

	return admin, nil
}
