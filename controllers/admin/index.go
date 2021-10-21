package admin

import (
	"blog/global"
	"blog/models"
	"blog/utils"
)

type IndexController struct {
	BaseController
}

func (i *IndexController) Index()  {
	var attrPowers []interface{}
	powers := i.Ctx.Input.Session(global.LoginUserPowers).(map[int64]interface{})
	for _,v := range powers{
		attrPowers = append(attrPowers,v)
	}
	admin,_ := i.Ctx.Input.Session(global.LoginUser).(models.AdminInfo)
	var base map[string]interface{} = map[string]interface{}{
		"powers" : attrPowers,
		"admin" : global.CONFIG.Base.Name,
		"username" : admin.UserName,
		"domain" : i.Ctx.Request.Host,
		"date" : utils.Date(),
	}
	i.Data["base"] = base
	i.TplName = "admin/public/base.html"
}

func (i *IndexController) Iframe()  {
	i.TplName = "admin/index/index.html"
}
