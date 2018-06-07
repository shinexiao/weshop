package form

import "github.com/pkg/errors"

// 管理员表单
type AdminForm struct {
	Id       uint   `form:"id"`       // 主键
	Username string `form:"username"` //用户名
	Password string `form:"password"` //密码
}

// 进行参数合法性验证
func (this AdminForm) Validate() (bool, error) {

	if len(this.Username) < 4 || len(this.Username) > 20 {

		return false, errors.New("用户名长度为4-20个字符")
	}

	if len(this.Password) < 6 || len(this.Password) > 30 {

		return false, errors.New("密码长度为6-30个字符")
	}

	return true, nil
}
