package service

import (
	"github.com/astaxie/beego"

	"BeegoSolution/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

func GetGanShowResrouces(adminName string) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	sql := " select DISTINCT * from tb_resource where parentId=-1 or id IN"
	sql += " (select resourceId as id from tb_role_resource"
	sql += " where roleId in "
	sql += " (select roleId FROM tb_user_role "
	sql += " where adminId=(select id as adminId from tb_admin where username=?)"
	sql += " )"
	sql += " )"
	sql += " and isShow=1"

	_, err := o.Raw(sql, adminName).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//根据角色id得到不在该角色的后台用户的集合
func GetNotInRoleAdmins(roleId int) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_admin where id not in(select adminId as id from tb_user_role where roleId=?)", roleId).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//根据角色id得到在该角色的后台用户的集合
func GetInRoleAdmins(roleId int) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_admin where id in(select adminId as id from tb_user_role where roleId=?)", roleId).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//将用户添加到角色
func AddUserToRole(roleId int, userId int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into tb_user_role(adminId,roleId) values(?,?)", userId, roleId).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//将用户集合添加到角色
func AddUsersToRole(roleId int, users []string) {
	for _, v := range users {
		userId, err := strconv.Atoi(v)
		utils.CheckError(err)
		AddUserToRole(roleId, userId)
	}
}

//将用户添加到角色
func DeleteUserFromRole(roleId int, userId int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_user_role where roleId=? and adminId=?", roleId, userId).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//将用户集合添加到角色
func DeleteUsersFromRole(roleId int, users []string) {
	for _, v := range users {
		userId, err := strconv.Atoi(v)
		utils.CheckError(err)
		beego.Info(v)
		beego.Info(userId)
		DeleteUserFromRole(roleId, userId)
	}
}

//根据后台用户名得到可以显示的资源集合
func GetCanShowResrouces(adminName string) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	sql := "select DISTINCT * from tb_resource where parentId=-1 or id IN"
	sql += "(select resourceId as id from tb_role_resource"
	sql += " where roleId in"
	sql += "(select roleId   FROM tb_user_role"
	sql += " where adminId=(select id as adminId from tb_admin where username=?)"
	sql += ")"
	sql += ")"
	sql += " and isShow=1"
	_, err := o.Raw(sql, adminName).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//根据角色id得到该角色可以访问的资源
func GetResroucesInRole(roleId int) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	sql := "select * from tb_resource where id in(select resourceId as id from tb_role_resource where roleId=?)"
	_, err := o.Raw(sql, roleId).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//根据角色id得到该角色可以访问的资源
func GetResroucesNotInRole(roleId int) (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	sql := "select * from tb_resource where id not in(select resourceId as id from tb_role_resource where roleId=?)"
	_, err := o.Raw(sql, roleId).ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//根据资源id判断该资源是否在角色中
func ISResourceInRole(resourceId int, roleId int) bool {
	result := false
	entitys := GetResroucesInRole(roleId)
	for _, val := range entitys {
		v, _ := strconv.Atoi(val[0].(string))
		if v == resourceId {
			result = true
			break
		}
	}
	return result
}

//根据角色id删除该角色可以访问的资源
func DeleteResrouceByRoleId(roleId int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_role_resource where roleId=?", roleId).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//将资源添加到角色
func AddReSourceToRole(roleId int, resourceId int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into tb_role_resource(resourceId,roleId) values(?,?)", resourceId, roleId).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//将资源集合添加到角色
func AddReSourcesToRole(roleId int, resourceIds []string) {
	for _, v := range resourceIds {
		resourceId, err := strconv.Atoi(v)
		utils.CheckError(err)
		AddReSourceToRole(roleId, resourceId)
	}
}
