package routers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterAdmin = func(ctx *context.Context) {
	sess := ctx.Input.CruSession.Get("Adminname")
	var lowerUrl string = strings.ToLower(ctx.Request.RequestURI)
	beego.Info(sess)
	if sess == nil && !NoCheckUrl(lowerUrl) {
		ctx.Redirect(302, "/admin/login")
	}
}

/*排除的URL */
func NoCheckUrl(lowerUrl string) bool {
	var flag bool = false
	flag = flag || lowerUrl == "/admin/login"
	flag = flag || lowerUrl == "/admin/logindo"
	flag = flag || lowerUrl == "/user/register"
	return flag
}
