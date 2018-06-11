package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
)

// 广告需要访问数据库
type AdvDao struct {
	db *gorm.DB
}

// 构造一个广告访问数据库的对象
func NewAdvDao(db *gorm.DB) *AdvDao {
	// 返回一个对象
	return &AdvDao{
		db: db,
	}
}

// 保存 广告信息
func (this *AdvDao) InsertAdv(adv *model.Adv) (int64, error) {
	// 往数据库插入广告信息  返回一个结果
	result := this.db.Save(adv)

	// 判断是否是错误信息
	if result.Error != nil {
		// 错误了
		return 0, result.Error
	}

	// 正确的话 返回正确条数
	return result.RowsAffected, nil
}

// 分页查询 广告
func (this *AdvDao) GetAdvs( page, limit int32) ([] *model.Adv, error) {
	advs := make([] *model.Adv, limit)
	result := this.db.Limit(limit).Offset((page - 1) * limit).Find(&advs)
	if result.Error != nil {
		return nil, result.Error
	}
	return advs, nil
}
