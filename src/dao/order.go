package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
)

// 订单数据库访问对象
type OrderDao struct {
	db *gorm.DB
}

// 构造OrderDao需要传入数据库对象
func NewOrderDao(db *gorm.DB) *OrderDao {

	return &OrderDao{
		db: db,
	}
}

// 保存订单信息
func (this *OrderDao) InsertOrder(order *model.Order) (int64, error) {

	result := this.db.Save(order)

	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
