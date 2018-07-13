package service

import (
	"beegoWeb/utils"
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

func GetDb() *sql.DB {
	db, _ := sql.Open("mysql", "root:root@@tcp(127.0.0.1:3306)/ibs")
	return db
}

//得到数据库中用户名为adminname密码为password的后台用户的数量
func getCount(adminname string, password string) int {
	o := orm.NewOrm()
	var count int
	o.Raw("select count(*) from tb_admin where username=? and password=?", adminname, password).QueryRow(&count)

	return count
}

//验证后台用户登录
func ValidateAdminLogin(adminname string, password string) bool {
	count := getCount(adminname, password)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//管理员注销登录
func LogoutAdmin() {

}

//判断是否存在同名后台用户
func IsExistAdmin(adminname string) bool {
	o := orm.NewOrm()
	var count int
	err := o.Raw("select count(*) from tb_admin where username=?", adminname).QueryRow(&count)
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//添加后台用户
func AddAdmin(adminName string, adminPassword string, isEnable bool) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into tb_admin(username,password,isEnable,createTime) values(?,?,?,?)", adminName, adminPassword, isEnable, time.Now()).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//获得后台用户集合
func GetAdmins(pageIndex int, pageSize int) (users []orm.ParamsList, totalCount int) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_admin order by createTime desc limit ? offset ?", pageSize, pageSize*(pageIndex-1)).ValuesList(&users)
	utils.CheckError(err)
	err = o.Raw("select count(*) from tb_admin").QueryRow(&totalCount)
	utils.CheckError(err)
	return
}

//删除后台用户
func DeleteAdmin(id int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_admin where id=?", id).Exec()
	utils.CheckError(err)
	fmt.Println(err)
	count, _ := res.RowsAffected()
	if count > 0 {
		return true
	} else {
		return false
	}
}

func GetAdminName(ctx *context.Context) string {
	adminName := ctx.Input.CruSession.Get("Adminname")
	if adminName == nil {
		return ""
	} else {
		return adminName.(string)
	}

}
