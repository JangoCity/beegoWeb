package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	c.TplName = "user/login.html"
}

func (c *UserController) Logindo() {
	c.TplName = "user/login.html"
}

func (c *UserController) Register() {
	c.TplName = "user/register.html"
}

func (c *UserController) Registerdo() {
	c.TplName = "user/register.html"
}
