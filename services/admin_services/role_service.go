package admin_services

import (
	"blog/global"
	"blog/models"
	"blog/services"
	"blog/validate/admin"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pkg/errors"
	"strings"
)

type RoleService struct {
	services.BaseService
}

func (this *RoleService) AddRole(param admin.AddRoleForm) error {
	data := models.AdminRole{
		Name:param.Name,
		Desc:param.Desc,
	}
	result := global.DB.Create(&data)
	if result.Error != nil{
		logs.Error(result.Error)
		return errors.New("新增失败")
	}
	return nil
}

func (this *RoleService) EditRole(param admin.EditRoleForm) error {
	data := models.AdminRole{
		Name:param.Name,
		Desc:param.Desc,
	}
	result := global.DB.Model(&models.AdminRole{}).Where("id = ?",param.Id).Updates(&data)
	if result.Error != nil{
		logs.Error(result.Error)
		return errors.New("编辑失败")
	}
	return nil
}

func (this *RoleService) DeleteRole(param admin.DeleteRoleForm) error {
	//result := global.DB.Where("id = ?",param.Id).Delete(&models.AdminInfo{})
	result := global.DB.Model(&models.AdminRole{}).Where("id = ?",param.Id).Update("status",1)
	if result.Error != nil{
		return errors.New("删除失败")
	}
	return nil
}

func (this *RoleService) Authority(param admin.AuthorityForm) error  {
	var powerSlice []string
	json.Unmarshal([]byte(param.Power),&powerSlice)
	result := global.DB.Model(&models.AdminRole{}).Where("id = ?",param.Id).Update("power",strings.Join(powerSlice,","))
	if result.Error != nil{
		return errors.New("提交失败")
	}
	return nil
}