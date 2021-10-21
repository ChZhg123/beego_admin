package admin

import (
	"blog/global"
	"blog/middleware"
	"blog/models"
	"blog/services/admin_services"
	"blog/validate/admin"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/server/web/captcha"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.ChallengeNums = 4
	cpt.StdHeight = 80
	cpt.StdWidth = 240
}

type LoginController struct {
	BaseController
}

type LoginData struct {
	Username  string `json:"username" form:"username"`
	Password  string `form:"password"`
	Captcha   string `form:"captcha"`
	CaptchaId string `form:"captcha_id"`
	Remember  int    `form:"remember"`
}

func (this *LoginController) Login() {
	if this.Ctx.Request.Method == "GET" {
		_, isLogin := middleware.IsLogin(this.Ctx)
		if isLogin {
			this.Ctx.Redirect(302, "/admin/index")
		}
		admin := map[string]interface{}{
			"name":  global.CONFIG.Base.Name,
			"title": "登录",
		}
		this.Data["admin"] = admin
		this.TplName = "admin/auth/login.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.LoginForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		//验证码
		if !cpt.VerifyReq(this.Ctx.Request) {
			this.Error("验证码不正确")
		}
		var authService admin_services.AuthService
		_, err2 := authService.CheckLogin(param, this.Ctx)
		if err2 != nil {
			this.Error(err2.Error())
		}
		this.Success("登录成功", global.Data)
	}
}

func (l *LoginController) Logout() {
	l.DelSession(global.LoginUser)
	l.DelSession(global.LoginUserPowers)
	l.Ctx.SetCookie(global.RememberUserToken, "", -1)
	l.Ctx.Redirect(302, "/admin/")
}

func (this *LoginController) ChangePwd() {
	if this.Ctx.Request.Method == "GET" {
		admin, _ := this.Ctx.Input.Session(global.LoginUser).(models.AdminInfo)
		this.Data["username"] = admin.UserName
		this.TplName = "admin/auth/password.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.ChangePwdForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		//新密码不能与旧密码一致
		if param.NewPassword == param.OldPassword {
			this.Error("新密码不能与旧密码一致")
		}
		//验证码
		if !cpt.VerifyReq(this.Ctx.Request) {
			this.Error("验证码不正确")
		}
		var authService admin_services.AuthService
		err2 := authService.ChangePwd(param)
		if err2 != nil {
			this.Error(err2.Error())
		}
		this.Success("修改成功", global.Data)
	}
}
