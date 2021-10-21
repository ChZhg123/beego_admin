package admin

import (
	"blog/validate"
)

type LoginForm struct {
	Username  string `form:"username" alias:"用户名" valid:"Required;MaxSize(20)"`
	Password  string `form:"password" alias:"密码" valid:"Required;MinSize(6)"`
	Captcha   string `form:"captcha" alias:"验证码" valid:"Required"`
	CaptchaId string `form:"captcha_id" valid:"Required"`
	Remember  int    `form:"remember" valid:"Required;Range(1,2)"`
}

type ChangePwdForm struct {
	Username  string `form:"username" alias:"用户名" valid:"Required;MaxSize(20)"`
	OldPassword  string `form:"old-password" alias:"旧密码" valid:"Required;MinSize(6)"`
	NewPassword  string `form:"new-password" alias:"新密码" valid:"Required;MinSize(6)"`
	Captcha   string `form:"captcha" alias:"验证码" valid:"Required"`
	CaptchaId string `form:"captcha_id" valid:"Required"`
}

func (param *ChangePwdForm) Validate() error  {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,ChangePwdForm{})
	return err
}

func (param *LoginForm) Validate() error {
	var baseValidate validate.BaseValidate
	err := baseValidate.Validate(param,LoginForm{})
	return err
}
