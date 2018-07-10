package service

import (
	"BeegoSolution/utils"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

//判断是否存在
func IsExistRole(name string) bool {
	o := orm.NewOrm()
	var count int
	err := o.Raw("select count(*) from tb_role where name=?", name).QueryRow(&count)
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//添加角色
func AddRole(name string, isBase bool) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into tb_role(name,createTime,isBase) values(?,?,?)", name, time.Now(), isBase).Exec()
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
func GetRoles(pageIndex int, pageSize int) (entitys []orm.ParamsList, totalCount int) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_role order by createTime desc limit ? offset ?", pageSize, pageSize*(pageIndex-1)).ValuesList(&entitys)
	utils.CheckError(err)
	err = o.Raw("select count(*) from tb_role").QueryRow(&totalCount)
	utils.CheckError(err)
	return
}

//删除后台用户
func DeleteRole(id int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_role where id=?", id).Exec()
	utils.CheckError(err)
	fmt.Println(err)
	count, _ := res.RowsAffected()
	if count > 0 {
		return true
	} else {
		return false
	}
}

//得到所有的角色集合
func GetAllRoles() (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_role order by createTime desc").ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//得到用户角色管理页面默认下拉列表中的角色id
func GetFirstRoleId() (firstId int) {
	o := orm.NewOrm()
	err := o.Raw("select id from tb_role order by createTime desc LIMIT 1").QueryRow(&firstId)
	utils.CheckError(err)
	return
}
