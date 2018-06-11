package dao

import (
	"github.com/jinzhu/gorm"
	"pojo/model"
	"bytes"
	"fmt"
)

//活动管理访问对象
type ActivityDao struct {
	db *gorm.DB
}

// 构造一个对象传入
func NewActivityDao(db *gorm.DB) *ActivityDao {
	return &ActivityDao{
		db: db,
	}
}

// 添加活动
func (this *ActivityDao) InsterActivity(active *model.Activity) (int64, error) {
	db := this.db.Create(active)
	return db.RowsAffected, db.Error
}

// 修改活动
func (this *ActivityDao) UpdateActivity(activity *model.Activity) (int64, error) {
	db := this.db.Model(&model.Activity{}).Update(activity)
	return db.RowsAffected, db.Error
}

// 删除活动
func (this *ActivityDao) DeleteActivity(id int64) (int64, error) {
	db := this.db.Where("id = ?",id).Delete(&model.Activity{})
	return db.RowsAffected, db.Error
}

// 查询活动
func (this *ActivityDao) QueryActivity(query *model.ActivityQY) ([]*model.Activity, error) {
	buf := new(bytes.Buffer)         	// 创建一个字符缓冲
	params := make([]interface{}, 0) 	//定义一个切片
	buf.WriteString("type = ?")   //组合判断条件
	params = append(params, 2)
	//if 0 != len(query.GoodsName) {
	//	buf.WriteString("and goods_name = ?")
	//	params = append(params, query.GoodsName)
	//}
	if !query.StartTime.IsZero() {
		buf.WriteString("and start_time >= ?")
		params = append(params, query.StartTime)
	}
	if !query.EndTime.IsZero() {
		buf.WriteString("and end_time <= ?")
		params = append(params, query.EndTime)
	}
	if 0 != query.GoodsId{
		buf.WriteString("and goods_id <= ?")
		params = append(params, query.GoodsId)
	}
	activity := []*model.Activity{} //定义一个空的数组
	db := this.db.Where(buf.String(), params...).Limit(query.Size).Offset(query.Size * (query.Page - 1)).Find(&activity)
	fmt.Println(db.Error)
	return activity, db.Error
}

// 查询单个活动
func (this *ActivityDao)QueryActivityById(id uint64)(model.Activity,error){
	activity := model.Activity{}
	first := this.db.First(&activity, id)
	return activity, first.Error
}

