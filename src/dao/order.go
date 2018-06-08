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

// 根据订单号查询订单
func (this *OrderDao) GetOrder(orderno string) (*model.Order, error) {

	order := &model.Order{}

	result := this.db.Find(order, "order_no = ?", orderno)

	if result.Error != nil {

		return nil, result.Error
	}

	return order, nil
}

// 删除订单信息
func (this *OrderDao) DeleteOrder(id uint) (int64, error) {

	order := &model.Order{}

	result := this.db.Delete(order, "id = ?", id)

	if result.Error != nil {

		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// 修改订单信息
func (this *OrderDao) UpdateOrder(order *model.Order) (int64, error) {

	result := this.db.Save(order)

	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

// 修改订单价格信息
func (this *OrderDao) UpdateOrderPrice(order *model.Order) (int64, error) {

	result := this.db.Model(order).Where("id = ?", order.ID).Update("price", order.Price)

	if result.Error != nil {
		// 保存出错了
		return 0, result.Error
	}

	return result.RowsAffected, nil
}