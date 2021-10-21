package routers

import (
	"blog/controllers/admin"
	"blog/controllers/index"
	"blog/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	middleware.InitAuth()
    beego.Router("/", &index.IndexController{},"get:Index")
    admin := beego.NewNamespace("/admin",
    	//授权
		beego.NSRouter("/",&admin.LoginController{},"get:Login"),
		beego.NSRouter("/login",&admin.LoginController{},"post:Login"),
		beego.NSRouter("/logout",&admin.LoginController{},"get:Logout"),
		beego.NSRouter("/change_password",&admin.LoginController{},"get:ChangePwd;post:ChangePwd"),
		beego.NSRouter("/index",&admin.IndexController{},"get:Index"),
		beego.NSRouter("/iframe",&admin.IndexController{},"get:Iframe"),
		//系统管理
		//用户模块
		beego.NSRouter("/system/userList",&admin.SystemController{},"get:UserList"),
		beego.NSRouter("/system/get_user_list",&admin.SystemController{},"get:GetUserList"),
		beego.NSRouter("/system/add_user",&admin.SystemController{},"get:AddUser;post:AddUser"),
		beego.NSRouter("/system/edit_user",&admin.SystemController{},"get:EditUser;post:EditUser"),
		beego.NSRouter("/system/delete_user",&admin.SystemController{},"post:DeleteUser"),
    	//角色模块
		beego.NSRouter("/system/roleList",&admin.SystemController{},"get:RoleList"),
		beego.NSRouter("/system/get_role_list",&admin.SystemController{},"get:GetRoleList"),
		beego.NSRouter("/system/add_role",&admin.SystemController{},"get:AddRole;post:AddRole"),
		beego.NSRouter("/system/edit_role",&admin.SystemController{},"get:EditRole;post:EditRole"),
		beego.NSRouter("/system/delete_role",&admin.SystemController{},"post:DeleteRole"),
		beego.NSRouter("/system/authority",&admin.SystemController{},"get:Authority;post:Authority"),
    )
    beego.AddNamespace(admin)
}
