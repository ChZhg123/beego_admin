package global

import (
	"blog/conf"
	"database/sql"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/server/web/session"
	"gorm.io/gorm"
)

const ErrorMsg = "服务异常"
const LoginUser = "LoginUser" //登录用户
const LoginUserPowers = "LoginUserPowers" //登录用户权限
const RememberUserToken = "remember_user_token"

var(
	CONFIG conf.Server
	Data struct{}
	Bm cache.Cache
	DB *gorm.DB
	SqlDB *sql.DB
	Session session.Store
)
