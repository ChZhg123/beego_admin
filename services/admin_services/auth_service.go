package admin_services

import (
	"blog/global"
	"blog/models"
	"blog/services"
	"blog/utils"
	"blog/validate/admin"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type AuthService struct {
	services.BaseService
}

func (a *AuthService) CheckLogin(data admin.LoginForm, ctx *context.Context) (*models.AdminInfo, error) {
	var AdminInfo models.AdminInfo
	admin := AdminInfo.GetUserByUserName(data.Username)
	if admin == nil {
		// 没有找到记录
		return nil, errors.New("用户名不存在")
	}
	encryptPassword := utils.Md5(utils.Md5(data.Password))
	if admin.PassWord != encryptPassword {
		return nil, errors.New("密码不正确")
	}
	updateData := map[string]interface{}{
		"LoginTime":time.Now(),
		"LoginIp":ctx.Request.RemoteAddr,
	}
	result := global.DB.Model(&admin).Updates(updateData)
	if result.RowsAffected ==0 {
		return nil, errors.New("登录失败")
	}
	if admin.Role.Power != "" {
		a.LoginBehavior(admin, admin.Role.Power, data.Remember, ctx) //登录行为
	} else {
		return nil, errors.New("没有权限")
	}
	return admin, nil
}

func (a *AuthService) LoginBehavior(admin *models.AdminInfo, power string, remember int, ctx *context.Context) {
	powers := a.GetPower(power)
	ctx.Output.Session(global.LoginUser,admin)
	ctx.Output.Session(global.LoginUserPowers,powers)
	if remember == 2 {
		ctx.SetCookie(global.RememberUserToken, utils.Base64Encode(admin.UserName), 3600*24)
	}
}

func (a *AuthService) GetPower(power string) map[int64]interface{} {
	var powers []*models.AdminPower
	data := make(map[int64]interface{})
	result := global.DB.Model(&models.AdminPower{}).Where("id in (?)",strings.Split(power,",")).Find(&powers)
	if result.RowsAffected != 0 {
		for _, v := range powers {
			value := make(map[string]interface{})
			value["title"] = v.Name
			value["icon"] = v.Icon
			value["href"] = "/" + v.Url
			value["spread"] = false
			value["children"] = make([]interface{}, 0)
			if _, ok := data[v.Id]; ok == false && v.Pid == 0 {
				data[v.Id] = value
			}
			if v.Pid != 0 {
				if _, ok1 := data[v.Pid]; ok1 {
					child := data[v.Pid].(map[string]interface{})["children"]
					data[v.Pid].(map[string]interface{})["children"] = append(child.([]interface{}), value)
				}
			}
		}
	}
	return data
}

func (a *AuthService) ChangePwd(data admin.ChangePwdForm) error  {
	var adminModel models.AdminInfo
	admin := adminModel.GetUserByUserName(data.Username)
	if admin == nil{
		return errors.New("用户不存在")
	}
	encryptPassword := utils.Md5(utils.Md5(data.OldPassword))
	if admin.PassWord != encryptPassword{
		return errors.New("旧密码不正确")
	}
	admin.PassWord = utils.Md5(utils.Md5(data.NewPassword))
	if result := global.DB.Save(admin); result.RowsAffected == 0 {
		return errors.New("修改失败")
	}
	return nil
}
