package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 订单模型
type Order struct {
	gorm.Model
	OrderNo     string    `gorm:"Column:order_no;UNIQUE;NOT NULL"`  //单号
	UserId      uint      `gorm:"Column:user_id;NOT NULL"`          //客户id
	UserName    string    `gorm:"Column:user_name;NOT NULL"`        //客户姓名
	UserPhone   string    `gorm:"Column:user_phone;NOT NULL"`       //客户手机号
	price       float32   `gorm:"Column:price;NOT NULL"`            //价格
	Status      int       `gorm:"Column:status;NOT NULL;DEFAULT:0"` //状态
	PayTime     time.Time `gorm:"Column:pay_time"`                  //支付时间
	payNo       string    `gorm:"Column:pay_no"`                    //支付交易号
	LogisticsNo string    `gorm:"Column:logistics_no"`              //物流单号
	GoodsMemo   string    `gorm:"Column:goods_memo;"`               //收货信息
}

// 返回表名
func (Order) TableName() string {
	return "shop_order"
}
