package service

import (
	"fmt"

	"dao"
	"pojo/form"
	"pojo/model"
	"util"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// 管理员服务层
type AdminService struct {
	dao *dao.AdminDao
}

// 构造管理员服务层
func NewAdminService(dao *dao.AdminDao) *AdminService {

	return &AdminService{
		dao: dao,
	}
}

// 保存管理员信息
func (this *AdminService) InsertAdmin(adminForm *form.AdminForm) error {

	if _, err := this.dao.GetAdmin(adminForm.Username); err != nil {

		// 没有查询到记录
		if err == gorm.ErrRecordNotFound {
			// 对密码进行加密并保存数据
			admin := &model.Admin{
				Username: adminForm.Username,
				Password: util.Sha1(adminForm.Password),
			}

			_, err = this.dao.InsertAdmin(admin)

			return err

		} else {
			// 保存数据出现其他错误
			return err
		}

	}

	//数据已经存在
	return errors.New(fmt.Sprintf("%s，已经存在", adminForm.Username))
}
