package controllers

import (
	"BeegoSolution/service"
	"BeegoSolution/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type RoleController struct {
	beego.Controller
}

func (c *RoleController) Add() {
	c.TplName = "role/add.html"
}

func (c *RoleController) Adddo() {
	name := c.GetString("name")
	isBase, err := strconv.ParseBool(c.GetString("isBase"))
	utils.CheckError(err)
	isExist := service.IsExistRole(name)
	if isExist == true {
		beego.Info("用户已存在")
	} else {
		service.AddRole(name, isBase)
	}

	c.TplName = "role/add.html"
	c.Redirect("/role/index", 301)
}

func (c *RoleController) Edit() {
	c.TplName = "role/edit.html"
}

func (c *RoleController) Editdo() {
	c.TplName = "role/edit.html"
}

func (c *RoleController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	utils.CheckError(err)
	service.DeleteRole(id)
	c.TplName = "role/index.html"
	c.Redirect("/role/index", 301)
}

func (c *RoleController) Index() {
	pageIndex, err := c.GetInt("pageIndex")
	utils.CheckError(err)
	pageSize := 10
	if pageIndex == 0 {
		pageIndex = 1
	}
	users, totalCount := service.GetRoles(pageIndex, pageSize)
	c.Data["PageIndex"] = pageIndex
	c.Data["PageSize"] = pageSize
	c.Data["TotalCount"] = totalCount
	c.Data["Entitys"] = users
	c.TplName = "role/index.html"
}
