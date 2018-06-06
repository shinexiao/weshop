package dao

import (
	"github.com/jinzhu/gorm"
)

// 订单模型
type Order struct {
	gorm.Model
	OrderNo string `gorm:column:order_no` //单号
	Status  int8   `gorm:column:status`  // 状态
}

// 返回表名
func (Order) TableName() string {
	return "shop_order"
}
