package admin

import (
	"blog/validate"
)

type AddRoleForm struct {
	Name  string `form:"name" alias:"角色名称" valid:"Required;MaxSize(20)"`
	Desc   string `form:"desc" alias:"描述" valid:"MaxSize(255)"`
}

type EditRoleForm struct {
	Id  int64 `form:"id" alias:"角色ID" valid:"Required"`
	Name  string `form:"name" alias:"角色名称" valid:"Required;MaxSize(20)"`
	Desc   string `form:"desc" alias:"描述" valid:"MaxSize(255)"`
}

type DeleteRoleForm struct {
	Id  int64 `form:"id" alias:"角色ID" valid:"Required"`
}

type AuthorityForm struct {
	Id  int64 `form:"id" alias:"角色ID" valid:"Required"`
	Power string `form:"power"`
}

func (param *AddRoleForm) Validate() error  {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,AddRoleForm{})
	return err
}

func (param *EditRoleForm) Validate() error {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,EditRoleForm{})
	return err
}

func (param *DeleteRoleForm) Validate() error {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,DeleteRoleForm{})
	return err
}

func (param *AuthorityForm) Validate() error  {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,AuthorityForm{})
	return err
}
