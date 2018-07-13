package controllers

import (
	"beegoWeb/service"
	"beegoWeb/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Login() {
	c.TplName = "admin/login.html"
}

func (c *AdminController) Logindo() {
	username := c.GetString("username")
	password := c.GetString("password")

	if service.ValidateAdminLogin(username, password) {

		c.SetSession("Adminname", username)
		c.Redirect("index", 301)
	} else {
		beego.Info("密码错误")
		c.TplName = "admin/login.html"
	}
}

func (c *AdminController) Logout() {
	c.DelSession("Adminname")
	c.Redirect("login", 301)
}

func (c *AdminController) Index() {
	adminName := service.GetAdminName(c.Ctx)
	canShowResrouces := service.GetCanShowResrouces(adminName)
	c.Data["CanShowResrouces"] = canShowResrouces
	c.TplName = "admin/index.html"
}

func (c *AdminController) Add() {
	c.TplName = "admin/add.html"
}

func (c *AdminController) Adddo() {
	adminUsername := c.GetString("username")
	adminPassword := c.GetString("password")
	isEnable, err := strconv.ParseBool(c.GetString("isEnable"))
	utils.CheckError(err)
	isExist := service.IsExistAdmin(adminUsername)
	if isExist == true {
		beego.Info("用户已存在")
	} else {
		service.AddAdmin(adminUsername, adminPassword, isEnable)
	}

	c.TplName = "admin/add.html"
	c.Redirect("/admin/list", 301)
}

func (c *AdminController) Edit() {
	c.TplName = "admin/edit.html"
}

func (c *AdminController) Editdo() {
	c.TplName = "admin/edit.html"
}

func (c *AdminController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	utils.CheckError(err)
	service.DeleteAdmin(id)
	c.TplName = "admin/list.html"
	c.Redirect("/admin/list", 301)
}

func (c *AdminController) List() {
	pageIndex := 1
	ci := c.Input().Get("pageIndex")
	if ci != "" {
		pi, err := strconv.Atoi(ci)
		utils.CheckError(err)
		pageIndex = pi
	}
	pageSize := 10
	users, totalCount := service.GetAdmins(pageIndex, pageSize)
	c.Data["PageIndex"] = pageIndex
	c.Data["PageSize"] = pageSize
	c.Data["TotalCount"] = totalCount
	c.Data["Entitys"] = users

	c.TplName = "admin/list.html"
}
