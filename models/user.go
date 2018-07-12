package models

import (
	"time"
)

type User struct {
	Id         int64
	Username   string
	Password   string
	Createtime time.Time
}

// func init() {
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
// 	orm.RegisterDataBase("default", beego.AppConfig.String("DB::db"),
// 		fmt.Sprintf("%s:%s@tcp(beego.AppConfig.String("DB::host"):beego.AppConfig.String("DB::port"))/%s?charset=utf8",
// 			beego.AppConfig.String("DB::user"),
// 			beego.AppConfig.String("DB::pass"),
// 			beego.AppConfig.String("DB::name")))
// 	orm.RegisterModel(new(User))
// 	orm.RunSyncdb("default", false, true)
// }
