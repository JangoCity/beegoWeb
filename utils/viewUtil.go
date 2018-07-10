//视图层自定义函数
package utils

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

//获取[]orm.ParamsList中的值，用于模版层，解决模版层使用中括号报错的问题
func GetParams(val orm.ParamsList, arg int) string {
	return val[arg].(string)
}

//获取[]orm.ParamsList中的值，用于模版层，解决模版层使用中括号报错的问题
func GetParamsInt(val orm.ParamsList, arg int) int {
	result, _ := strconv.Atoi(val[arg].(string))
	return result
}
