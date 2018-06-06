package dao

import "github.com/jinzhu/gorm"

type active struct {
	gorm.Model
	IsDeleted   int32   `gorm:"column:is_deleted;type:int(1);not null" json:"is_deleted"`     //是否删除
	NewCash     float64 `gorm:"column:new_cash;type:double(64);null" json:"new_cash"`         //新客户返现金卷
	OldCash     float64 `gorm:"column:old_cash;type:double(64);null" json:"old_cash"`         //老客户返现金卷
	Type        int32   `gorm:"column:type;type:int(1);not null" json:"type"`                 //1、推广有礼 or  2、购买返现
	OneRebate   float32 `gorm:"column:one_rebate;type:double(64);null" json:"one_rebate"`     //第一级返利
	FourRebate  float32 `gorm:"column:four_rebate;type:double(64);null" json:"four_rebate"`   //第二级返利
	ThreeRebate float32 `gorm:"column:three_rebate;type:double(64);null" json:"three_rebate"` //第三级返利
	goodsId     uint64  `gorm:"column:goods_id;type:bigint(32);null" json:"goods_id"`         //商品id
}


