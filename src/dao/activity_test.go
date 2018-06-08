package dao

import (
	"testing"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"pojo/model"
	_ "os"
	"fmt"
	"math/rand"
	"time"
	"os"
)

/**
单元测试 测试数据库连接
 */
func TestNewActiveDao(t *testing.T) {
	var db *gorm.DB

	activety := NewAdminDao(nil)

	assert.Equal(t, db, activety.db)
}

/**
单元测试 测试批量添加操作
 */
func TestActivetyDao_InsterActivity(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	//创建表
	db.CreateTable(model.Activity{})

	activity := &model.Activity{
		Type:    1,
		NewCash: 12.0,
		OldCash: 11.0,
	}
	dao := NewActivityDao(db)
	result, err := dao.InsterActivity(activity)
	assert.Equal(t, err, nil)
	assert.Equal(t, result, int64(1))
	//批量插入数据库
	for i := 0.01; i < 0.1; i = i + 0.01 {
		activity = &model.Activity{
			Type:        2,
			OneRebate:   0.1 + i,
			TwoRebate:   0.2 + i,
			ThreeRebate: 0.3 + i,
			StartTime:   time.Now(),
			EndTime:     time.Now(),
		}
		result, err = dao.InsterActivity(activity)
		fmt.Println(result, i)
		assert.Equal(t, nil, err)
		assert.Equal(t, int64(1), result)
	}
	//os.Remove("unit_test_gorm.db")
}

/**
单元测试 测试批量查询
 */
func TestActivityDao_QueryActivity(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	dao := NewActivityDao(db)
	query := &model.ActivityQY{
		Type: 2,
		Page: 1,
		Size: 10,
		GoodsId:15,
		StartTime:time.Unix(time.Now().Unix()-10000,2222),
		EndTime:time.Unix(time.Now().Unix()+10000,2222),
	}
	activities, e := dao.QueryActivity(query)
	fmt.Println(e)
	assert.Equal(t, e, nil)
	assert.Equal(t, true, len(activities) > 0)
	//fmt.Println(activities)
	for _, activity := range activities {
		fmt.Println(activity)
	}
}

/*
单元测试 测试通过id查询
 */
func TestActivityDao_QueryActivityById(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	dao := NewActivityDao(db)
	//产生随机数字
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(10)
	activity, e := dao.QueryActivityById(uint64(intn + 1))
	assert.Equal(t, e, nil)
	assert.NotEqual(t, activity, nil)
}

/**
单元测试 测试修改参数
 */
func TestActivityDao_UpdateActivity(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	dao := NewActivityDao(db)
	//产生随机数字
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(10)
	fmt.Println(intn)
	activity := &model.Activity{
		GoodsId: 15,
	}
	updateActivity, e := dao.UpdateActivity(activity)
	assert.Equal(t, nil, e)
	assert.Equal(t, int64(11), updateActivity)

	id, e := dao.QueryActivityById(uint64(intn+1))
	assert.Equal(t, nil, e)
	fmt.Println(id)
	assert.NotEqual(t, nil, id)
}

/**
单元测试 删除操作
 */
func TestActivityDao_DeleteActivity(t *testing.T) {
	db, _ := gorm.Open("sqlite3", "unit_test_gorm.db")
	dao := NewActivityDao(db)
	//产生随机数字
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(10)
	fmt.Println(intn)
	i, e := dao.DeleteActivity(int64(intn+1))
	assert.Equal(t, nil, e)
	assert.NotEqual(t, 1, i)

	// 数据查询
	id,err := dao.QueryActivityById(uint64(intn+1))
	fmt.Println(id,err)
	assert.NotEqual(t, err, nil)
	assert.NotEqual(t, id, nil)
	// 删除数据库
	defer os.Remove("unit_test_gorm.db")
}
