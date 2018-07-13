package controllers

import (
	"beegoWeb/service"
	"beegoWeb/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type AuthorityController struct {
	beego.Controller
}

func (c *AuthorityController) UserRole() {
	firstId := service.GetFirstRoleId()
	c.Data["Roles"] = service.GetAllRoles()
	c.Data["FirstId"] = firstId
	c.Data["AdminsNotInRole"] = service.GetNotInRoleAdmins(firstId)
	c.Data["AdminsInRole"] = service.GetInRoleAdmins(firstId)
	c.TplName = "authority/userRole.html"
}

func (c *AuthorityController) UserRoleDo() {
	firstId, err := strconv.Atoi(c.GetString("listRole"))
	utils.CheckError(err)
	c.Data["Roles"] = service.GetAllRoles()
	c.Data["FirstId"] = firstId

	c.Data["AdminsNotInRole"] = service.GetNotInRoleAdmins(firstId)
	c.Data["AdminsInRole"] = service.GetInRoleAdmins(firstId)
	c.TplName = "authority/userRole.html"
}

func (c *AuthorityController) AddUsersInRole() {
	userIds := strings.Split(c.GetString("array"), ",")
	roleId, _ := strconv.Atoi(c.GetString("roleId"))

	service.AddUsersToRole(roleId, userIds)
	c.TplName = "authority/userRole.html"
}

func (c *AuthorityController) DeleteUsersFromRole() {
	userIds := strings.Split(c.GetString("array"), ",")
	roleId, _ := strconv.Atoi(c.GetString("roleId"))
	service.DeleteUsersFromRole(roleId, userIds)
	c.TplName = "authority/userRole.html"
}

func (c *AuthorityController) RoleResource() {
	firstId := service.GetFirstRoleId()
	c.Data["Roles"] = service.GetAllRoles()
	c.Data["FirstId"] = firstId
	c.Data["AllResources"] = service.GetAllResources()
	c.Data["ResourcesInRole"] = service.GetResroucesInRole(firstId)
	c.TplName = "authority/roleResource.html"
}

func (c *AuthorityController) RoleResourceDo() {
	firstId, _ := strconv.Atoi(c.GetString("listRole"))
	c.Data["Roles"] = service.GetAllRoles()
	c.Data["FirstId"] = firstId
	c.Data["AllResources"] = service.GetAllResources()
	c.Data["ResourcesInRole"] = service.GetResroucesInRole(firstId)
	c.TplName = "authority/roleResource.html"
}

func (c *AuthorityController) AddResourcesInRole() {
	resourceIds := strings.Split(c.GetString("array"), ",")
	beego.Info(resourceIds)
	roleId, _ := strconv.Atoi(c.GetString("roleId"))
	service.DeleteResrouceByRoleId(roleId)
	service.AddReSourcesToRole(roleId, resourceIds)
	c.TplName = "authority/roleResource.html"
}
