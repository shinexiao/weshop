package dao

import (
	"github.com/jinzhu/gorm"
)

// 订单明细模型
type OrderDetail struct {
	gorm.Model
	OrderId      uint    `gorm:"Column:order_id;NOT NULL"`      //订单id
	GoodsId      uint    `gorm:"Column:goods_id;NOT NULL"`      //商品id
	GoodsName    string  `gorm:"Column:goods_name;NOT NULL"`    //商品名称
	GoodsPicture string  `gorm:"Column:goods_picture;NOT NULL"` //商品图片
	GoodsAmount  int     `gorm:"Column:goods_amount;NOT NULL"`  //商品数量
	GoodsPrice   float32 `gorm:"Column:goods_price;NOT NULL"`   //商品单价
}

// 返回表名
func (OrderDetail) TableName() string {
	return "shop_order_detail"
}
