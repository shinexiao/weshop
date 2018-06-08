package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
活动结构体
 */
type Activity struct {
	gorm.Model
	NewCash     float64   `gorm:"column:new_cash;type:double(64);null" json:"new_cash"`         // 新客户返现金卷
	OldCash     float64   `gorm:"column:old_cash;type:double(64);null" json:"old_cash"`         // 老客户返现金卷
	Type        int32     `gorm:"column:type;type:int(1);not null" json:"type"`                 // 1、推广有礼 or  2、购买返现
	OneRebate   float64   `gorm:"column:one_rebate;type:double(64);null" json:"one_rebate"`     // 第一级返利
	TwoRebate   float64   `gorm:"column:two_rebate;type:double(64);null" json:"twos_rebate"`    // 第二级返利
	ThreeRebate float64   `gorm:"column:three_rebate;type:double(64);null" json:"three_rebate"` // 第三级返利
	GoodsId     uint64    `gorm:"column:goods_id;type:bigint(32);null" json:"goods_id"`         // 商品id
	StartTime   time.Time `json:"start_time"`        // 活动开始时间
	EndTime     time.Time `json:"end_time"`            // 活动结束时间
	//StartTime   time.Time `gorm:"column:start_time;type:datatime;null json:"start_time"`        // 活动开始时间
	//EndTime     time.Time `gorm:"column:end_time;type:datatime;null json:"end_time"`            // 活动结束时间
}

/**
查询活动结构体
 */
type ActivityQY struct {
	gorm.Model
	GoodsName string    `json:"goods_name" form:"goods_name"` // 商品名称
	Type      int32     `json:"type" form:"type"`             // 1、推广有礼 or  2、购买返现
	GoodsId   uint64    `json:"goods_id" form:"goods_id"`     // 商品id
	Page      int64     `json:"page" form:"page" `            // 分页
	Size      int64     `json:"size" form:"size"`             // 分页
	StartTime time.Time `json:"start_time" form:"start_time"` // 活动开始时间
	EndTime   time.Time `json:"end_time" form:"end_time"`     // 活动结束时间
}

/**
指定数据库表名称
 */
func (Activity) TableName() string {
	return "activity"
}
