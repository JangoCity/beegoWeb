package main

import (
	_ "BeegoSolution/routers"
	"BeegoSolution/service"
	"BeegoSolution/utils"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver(beego.AppConfig.String("db"), orm.DRMySQL)
	orm.RegisterDataBase("default", beego.AppConfig.String("db"),
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			beego.AppConfig.String("user"),
			beego.AppConfig.String("pass"),
			beego.AppConfig.String("host"),
			beego.AppConfig.String("port"),
			beego.AppConfig.String("name")))
	// orm.RegisterModel(new(models.User))
	// orm.RunSyncdb("default", false, true)
}

func main() {
	//开启Session
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.AddFuncMap("GetParams", utils.GetParams)
	beego.AddFuncMap("GetParamsInt", utils.GetParamsInt)
	beego.AddFuncMap("ISResourceInRole", service.ISResourceInRole)
	beego.AddFuncMap("GetSubResourcesCount", service.GetSubResourcesCount)

	beego.Run()
}
