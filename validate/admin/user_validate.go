package admin

import (
	"blog/validate"
)

type AddUserForm struct {
	RoleId  int64 `form:"role_id" alias:"角色" valid:"Required"`
	UserName  string `form:"username" alias:"用户名" valid:"Required;MaxSize(20)"`
	NickName   string `form:"nickname" alias:"名称" valid:"Required;MaxSize(20)"`
	Phone string `form:"phone" alias:"联系方式" valid:"Phone"`
	PassWord  string `form:"password" alias:"登录密码" valid:"Required;MinSize(6)"`
}

type EditUserForm struct {
	Id  int64 `form:"id" alias:"用户ID" valid:"Required"`
	RoleId  int64 `form:"role_id" alias:"角色" valid:"Required"`
	UserName  string `form:"username" alias:"用户名" valid:"Required;MaxSize(20)"`
	NickName   string `form:"nickname" alias:"名称" valid:"Required;MaxSize(20)"`
	Phone string `form:"phone" alias:"联系方式" valid:"Phone"`
	PassWord  string `form:"password" alias:"登录密码" valid:"MinSize(6)"`
}

type DeleteUserForm struct {
	Id  int64 `form:"id" alias:"用户ID" valid:"Required"`
}

func (param *AddUserForm) Validate() error  {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,AddUserForm{})
	return err
}

func (param *EditUserForm) Validate() error {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,EditUserForm{})
	return err
}

func (param *DeleteUserForm) Validate() error {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,DeleteUserForm{})
	return err
}
