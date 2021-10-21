package index

import "blog/controllers"

type IndexController struct {
	controllers.BaseController
}

func (i *IndexController) Index()  {
	if i.Ctx.Request.Method == "GET" {
		i.Data["Website"] = "beego.me"
		i.Data["Email"] = "astaxie@gmail.com"
		i.TplName = "index/index/index.tpl"
	}
}
