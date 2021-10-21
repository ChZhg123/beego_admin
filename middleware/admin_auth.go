package middleware

import (
	"blog/global"
	"blog/models"
	"blog/services/admin_services"
	"blog/utils"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	context2 "github.com/beego/beego/v2/server/web/context"
	"strings"
)

var authExcept map[string]int = map[string]int{
	"admin/":0,
	"admin/login":1,
}

func InitAuth()  {
	var FilterUser = func(ctx *context.Context) {
		url := strings.TrimLeft(ctx.Input.URI(),"/")
		if _,ok := authExcept[url];!ok{
			_,isLogin := IsLogin((*context2.Context)(ctx))
			if !isLogin{
				ctx.Redirect(302,"/admin/")
			}
		}
	}
	beego.InsertFilter("/admin/*",beego.BeforeRouter,FilterUser)
}

func IsLogin(ctx *context2.Context) (*models.AdminInfo,bool) {
	adminUser,ok := ctx.Input.Session(global.LoginUser).(*models.AdminInfo)
	if !ok{
		adminToken := ctx.GetCookie(global.RememberUserToken)
		if adminToken==""{
			return nil,false
		}
		username := utils.Base64Decode(adminToken)
		var adminModel models.AdminInfo
		admin := adminModel.GetUserByUserName(username)
		if admin==nil {
			return nil,false
		}else{
			var authServer admin_services.AuthService
			if admin.Role.Power != "" {
				authServer.LoginBehavior(admin, admin.Role.Power, 2, ctx) //登录行为
				return admin,true
			} else {
				return nil, false
			}
		}
	}else{
		return adminUser,true
	}
}
