package admin

import (
	"blog/global"
	"blog/models"
	"blog/services/admin_services"
	"blog/utils"
	"blog/validate/admin"
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type SystemController struct {
	BaseController
}

//用户模块

func (this *SystemController) UserList()  {
	this.TplName = "admin/system/user_list.html"
}

type UserList struct {
	Id int64 `gorm:"column:id" json:"id"`
	UserName string `gorm:"column:username;unique;size:20;not null" json:"username"`
	NickName string `gorm:"column:nickname;size:20" json:"nickname"`
	Phone string `gorm:"column:phone;size:20" json:"phone"`
	RoleId int64 `gorm:"column:role_id;size:11;default:0" json:"role_id"`
	Role UserRole `json:"role" gorm:"foreignKey:role_id"`
	CreateTime models.MyTime `gorm:"column:create_time" json:"create_time"`
}

type UserRole struct {
	Id int64 `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (this *SystemController) GetUserList()  {
	keywords := this.GetString("keywords","")
	page := this.GetString("page","1")
	per_page := this.GetString("limit","10")
	var admin []UserList
	result := global.DB.Model(&models.AdminInfo{}).Where("admin_info.status = ?",0).Joins("Role").Where("username like ? OR nickname like ?","%"+keywords+"%","%"+keywords+"%").Select("admin_info.id","admin_info.username","admin_info.nickname","admin_info.phone","admin_info.role_id","admin_info.create_time").Scopes(models.Paginate(page,per_page)).Order("create_time desc").Find(&admin)
	this.LayTableApi(result.RowsAffected,admin)
}

/**
新增用户
 */
func (this *SystemController) AddUser()  {
	if this.Ctx.Request.Method == "GET" {
		var role []models.AdminRole
		global.DB.Model(&models.AdminRole{}).Scopes(models.Status).Select("id", "name").Find(&role)
		this.Data["role"] = role
		this.TplName = "admin/system/add_user.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.AddUserForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		var service admin_services.UserService
		err2 := service.AddUser(param)
		if err2 != nil{
			this.Error(err2.Error())
		}
		this.Success("新增成功", global.Data)
	}
}

/**
编辑用户
 */
func (this *SystemController) EditUser()  {
	if this.Ctx.Request.Method == "GET" {
		id := this.GetString("id")
		var user models.AdminInfo
		global.DB.Model(&models.AdminInfo{}).Where("id = ?",id).First(&user)
		var role []models.AdminRole
		global.DB.Model(&models.AdminRole{}).Scopes(models.Status).Select("id", "name").Find(&role)
		this.Data["role"] = role
		this.Data["user"] = user
		this.TplName = "admin/system/edit_user.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.EditUserForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		var service admin_services.UserService
		err2 := service.EditUser(param)
		if err2 != nil{
			this.Error(err2.Error())
		}
		this.Success("编辑成功", global.Data)
	}
}

/**
删除用户
 */
func (this *SystemController) DeleteUser()  {
	var param admin.DeleteUserForm
	if err := this.ParseForm(&param); err != nil {
		this.Error(err.Error())
	}
	//验证数据
	if err1 := param.Validate(); err1 != nil {
		this.Error(err1.Error())
	}
	var service admin_services.UserService
	err2 := service.DeleteUser(param)
	if err2 != nil{
		this.Error(err2.Error())
	}
	this.Success("删除成功", global.Data)
}


//角色模块
func (this *SystemController) RoleList()  {
	this.TplName = "admin/system/role_list.html"
}

func (this *SystemController) GetRoleList()  {
	keywords := this.GetString("keywords","")
	page := this.GetString("page","1")
	per_page := this.GetString("limit","10")
	var role []models.AdminRole
	result := global.DB.Model(&models.AdminRole{}).Scopes(models.Status).Where("name like ?","%"+keywords+"%").Select("id","name","desc","create_time").Scopes(models.Paginate(page,per_page)).Order("create_time desc").Find(&role)
	this.LayTableApi(result.RowsAffected,role)
}

/**
新增角色
*/
func (this *SystemController) AddRole()  {
	if this.Ctx.Request.Method == "GET" {
		this.TplName = "admin/system/add_role.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.AddRoleForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		var service admin_services.RoleService
		err2 := service.AddRole(param)
		if err2 != nil{
			this.Error(err2.Error())
		}
		this.Success("新增成功", global.Data)
	}
}

/**
编辑角色
*/
func (this *SystemController) EditRole()  {
	if this.Ctx.Request.Method == "GET" {
		id := this.GetString("id")
		var role models.AdminRole
		global.DB.Model(&models.AdminRole{}).Where("id = ?",id).First(&role)
		this.Data["role"] = role
		this.TplName = "admin/system/edit_role.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.EditRoleForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		var service admin_services.RoleService
		err2 := service.EditRole(param)
		if err2 != nil{
			this.Error(err2.Error())
		}
		this.Success("编辑成功", global.Data)
	}
}

/**
删除角色
*/
func (this *SystemController) DeleteRole()  {
	var param admin.DeleteRoleForm
	if err := this.ParseForm(&param); err != nil {
		this.Error(err.Error())
	}
	//验证数据
	if err1 := param.Validate(); err1 != nil {
		this.Error(err1.Error())
	}
	var service admin_services.RoleService
	err2 := service.DeleteRole(param)
	if err2 != nil{
		this.Error(err2.Error())
	}
	this.Success("删除成功", global.Data)
}

func (this *SystemController) Authority()  {
	if this.Ctx.Request.Method == "GET" {
		id := this.GetString("id")
		var role models.AdminRole
		global.DB.Model(&models.AdminRole{}).Where("id = ?",id).First(&role)
		power := strings.Split(role.Power,",")
		data := this.GetPowers()
		for i,v := range data{
			if utils.IsSlice(power,strconv.FormatInt(v.Id,10)){
				data[i].Select = 1
				for j,v1 := range v.Child{
					if utils.IsSlice(power,strconv.FormatInt(v1.Id,10)){
						data[i].Child[j].Select = 1
					}
				}
			}
		}
		this.Data["id"] = id
		this.Data["data"] = data
		this.TplName = "admin/system/authority.html"
	}
	if this.Ctx.Request.Method == "POST" {
		var param admin.AuthorityForm
		if err := this.ParseForm(&param); err != nil {
			this.Error(err.Error())
		}
		//验证数据
		if err1 := param.Validate(); err1 != nil {
			this.Error(err1.Error())
		}
		var service admin_services.RoleService
		err2 := service.Authority(param)
		if err2 != nil{
			this.Error(err2.Error())
		}
		this.Success("提交成功", global.Data)
	}
}

type AuthData struct {
	models.BaseModel
	Name string `json:"name"`
	Url string `json:"url"`
	Pid int64 `gorm:"column:pid" json:"pid"`
	Sort int `json:"sort"`
	Icon string `json:"icon"`
	Select int `json:"select"`
	Child []AuthData `json:"child" gorm:"foreignKey:Pid"`
}

func (this *SystemController) GetPowers() []AuthData {
	global.Bm.Delete(context.TODO(),"authority")
	data,_ := global.Bm.Get(context.TODO(),"authority")
	powers,ok := data.([]AuthData)
	if !ok{
		var powers_data []models.AdminPower
		global.DB.Model(&models.AdminPower{}).Where("pid = ?",0).Preload("Child").Order("id asc").Find(&powers_data)
		for _,v := range powers_data{
			var power AuthData
			StructAssign(&power,&v)
			powers = append(powers,power)
		}
		global.Bm.Put(context.TODO(),"authority",powers,3600*time.Second)
	}
	return powers
}

func StructAssign(bind interface{},value interface{})  {
	bval := reflect.ValueOf(bind).Elem()
	val := reflect.ValueOf(value).Elem()
	vtype := reflect.TypeOf(value).Elem()
	for i:=0;i<val.NumField();i++{
		name := vtype.Field(i).Name
		baseType := val.Field(i).Kind().String()
		if ok := bval.FieldByName(name).IsValid();ok{
			if baseType == "slice"{
				aa,_ := bval.FieldByName(name).Interface().([]AuthData)
				bb,_ := val.Field(i).Interface().([]models.AdminPower)
				for _,v := range bb{
					var power AuthData
					StructAssign(&power,&v)
					aa = append(aa,power)
				}
				bval.FieldByName(name).Set(reflect.ValueOf(aa))
				//bval.FieldByName(name).Set(reflect.ValueOf(val.Field(i).Interface()))
			}else{
				bval.FieldByName(name).Set(reflect.ValueOf(val.Field(i).Interface()))
			}
		}
	}
}
