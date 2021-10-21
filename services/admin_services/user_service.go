package admin_services

import (
	"blog/global"
	"blog/models"
	"blog/services"
	"blog/utils"
	"blog/validate/admin"
	"github.com/beego/beego/v2/core/logs"
	"github.com/pkg/errors"
)

type UserService struct {
	services.BaseService
}

func (this *UserService) AddUser(param admin.AddUserForm) error {
	var user models.AdminInfo
	result := global.DB.Model(&models.AdminInfo{}).Where("username = ?",param.UserName).First(&user)
	if result.RowsAffected != 0{
		return errors.New("用户名已存在")
	}
	data := models.AdminInfo{
		NickName:param.NickName,
		UserName:param.UserName,
		RoleId:param.RoleId,
		PassWord:utils.Md5(utils.Md5(param.PassWord)),
		Phone: param.Phone,
	}
	result = global.DB.Create(&data)
	if result.Error != nil{
		logs.Error(result.Error)
		return errors.New("新增失败")
	}
	return nil
}

func (this *UserService) EditUser(param admin.EditUserForm) error {
	var user models.AdminInfo
	result := global.DB.Model(&models.AdminInfo{}).Where("username = ? AND id <> ?",param.UserName,param.Id).First(&user)
	if result.RowsAffected != 0{
		return errors.New("用户名已存在")
	}
	data := models.AdminInfo{
		NickName:param.NickName,
		UserName:param.UserName,
		RoleId:param.RoleId,
		Phone: param.Phone,
	}
	if param.PassWord != ""{
		data.PassWord = utils.Md5(utils.Md5(param.PassWord))
	}
	result = global.DB.Model(&models.AdminInfo{}).Where("id = ?",param.Id).Updates(&data)
	if result.Error != nil{
		logs.Error(result.Error)
		return errors.New("编辑失败")
	}
	return nil
}

func (this *UserService) DeleteUser(param admin.DeleteUserForm) error {
	//result := global.DB.Where("id = ?",param.Id).Delete(&models.AdminInfo{})
	result := global.DB.Model(&models.AdminInfo{}).Where("id = ?",param.Id).Update("status",1)
	if result.Error != nil{
		return errors.New("删除失败")
	}
	return nil
}