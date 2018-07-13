package service

import (
	"beegoWeb/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

//判断是否存在
func IsExistResource(controller string, action string) bool {
	o := orm.NewOrm()
	var count int
	err := o.Raw("select count(*) from tb_resource where controller=? and action=?", controller, action).QueryRow(&count)
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//添加资源
func AddResource(parentId int, name string, controller string, action string, sort int, url string, isShow int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("insert into tb_resource(parentId,name,controller,action,sort,url,isShow) values(?,?,?,?,?,?,?)", parentId, name, controller, action, sort, url, isShow).Exec()
	utils.CheckError(err)
	count, err := res.RowsAffected()
	utils.CheckError(err)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//获得资源集合
func GetResource(pageIndex int, pageSize int) (entitys []orm.ParamsList, totalCount int) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_resource limit ? offset ?", pageSize, pageSize*(pageIndex-1)).ValuesList(&entitys)
	utils.CheckError(err)
	err = o.Raw("select count(*) from tb_resource").QueryRow(&totalCount)
	utils.CheckError(err)
	return
}

//删除资源
func DeleteResource(id int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_resource where id=?", id).Exec()
	utils.CheckError(err)
	fmt.Println(err)
	count, _ := res.RowsAffected()
	if count > 0 {
		return true
	} else {
		return false
	}
}

//获得所有资源
func GetAllResources() (entitys []orm.ParamsList) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_resource").ValuesList(&entitys)
	utils.CheckError(err)
	return
}

//获得分页的资源集合
func GetResources(pageIndex int, pageSize int) (entitys []orm.ParamsList, totalCount int) {
	o := orm.NewOrm()
	_, err := o.Raw("select * from tb_resource limit ? offset ?", pageSize, pageSize*(pageIndex-1)).ValuesList(&entitys)
	utils.CheckError(err)
	err = o.Raw("select count(*) from tb_resource").QueryRow(&totalCount)
	utils.CheckError(err)
	return
}

//删除后台用户
func DeleteResrouce(id int) bool {
	o := orm.NewOrm()
	res, err := o.Raw("delete from tb_resource where id=?", id).Exec()
	utils.CheckError(err)
	fmt.Println(err)
	count, _ := res.RowsAffected()
	if count > 0 {
		return true
	} else {
		return false
	}
}

//根据父资源id得到子资源id的数量
//func GetSubResourcesCount(resourceId int) (count int) {
//	o := orm.NewOrm()
//	beego.Info(resourceId)
//	err := o.Raw("select count(*) from tb_resource where parentId=?", resourceId).QueryRow(&count)
//	utils.CheckError(err)
//	return count
//}

//根据父资源id得到子资源id的数量
func GetSubResourcesCount(resourceId int, entitys []orm.ParamsList) int {
	count := 0
	for _, val := range entitys {
		subResourceId, _ := strconv.Atoi(val[1].(string))
		if subResourceId == resourceId {
			count++
		}
	}

	return count
}
