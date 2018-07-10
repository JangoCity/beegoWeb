package routers

import (
	"BeegoSolution/controllers"

	"github.com/astaxie/beego"
)

func init() {
	/*前台模块*/
	//首页
	beego.Router("/", &controllers.MainController{})

	//登录
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/logindo", &controllers.UserController{}, "post:Logindo")

	//注册
	beego.Router("/user/register", &controllers.UserController{}, "get:Register")
	beego.Router("/user/registerdo", &controllers.UserController{}, "Post:Registerdo")

	/*后台用户模块*/
	//登录
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login")
	beego.Router("/admin/logindo", &controllers.AdminController{}, "*:Logindo")
	//	后台首页
	beego.Router("/admin/index", &controllers.AdminController{}, "get:Index")
	//用户管理
	beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
	beego.Router("/admin/add", &controllers.AdminController{}, "post:Adddo")
	beego.Router("/admin/edit", &controllers.AdminController{}, "get:Edit")
	beego.Router("/admin/editdo", &controllers.AdminController{}, "post:Editdo")
	beego.Router("/admin/delete/?:id", &controllers.AdminController{}, "get:Delete")
	beego.Router("/admin/list", &controllers.AdminController{}, "get:List")

	//角色管理
	beego.Router("/role/add", &controllers.RoleController{}, "get:Add")
	beego.Router("/role/add", &controllers.RoleController{}, "post:Adddo")
	beego.Router("/role/edit", &controllers.RoleController{}, "get:Edit")
	beego.Router("/role/editdo", &controllers.RoleController{}, "post:Editdo")
	beego.Router("/role/delete/?:id", &controllers.RoleController{}, "get:Delete")
	beego.Router("/role/index", &controllers.RoleController{}, "get:Index")

	//资源管理
	beego.Router("/resource/add", &controllers.ResourceController{}, "get:Add")
	beego.Router("/resource/add", &controllers.ResourceController{}, "post:Adddo")
	beego.Router("/resource/edit", &controllers.ResourceController{}, "get:Edit")
	beego.Router("/resource/edit", &controllers.ResourceController{}, "post:Editdo")
	beego.Router("/resource/delete/?:id", &controllers.ResourceController{}, "get:Delete")
	beego.Router("/resource/index", &controllers.ResourceController{}, "get:Index")

	//权限管理
	beego.Router("/authority/user-role", &controllers.AuthorityController{}, "get:UserRole")
	beego.Router("/authority/user-role", &controllers.AuthorityController{}, "post:UserRoleDo")
	beego.Router("/authority/role-resource", &controllers.AuthorityController{}, "get:RoleResource")
	beego.Router("/authority/role-resource", &controllers.AuthorityController{}, "post:RoleResourceDo")
	beego.Router("/authority/addUsersInRole", &controllers.AuthorityController{}, "post:AddUsersInRole")
	beego.Router("/authority/deleteUsersFromRole", &controllers.AuthorityController{}, "post:DeleteUsersFromRole")
	beego.Router("/authority/addResourcesInRole", &controllers.AuthorityController{}, "post:AddResourcesInRole")
	initFilter()
}

func initFilter() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, FilterAdmin)
}
