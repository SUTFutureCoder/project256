package actions

import (
	"github.com/gin-gonic/gin"
	//"project256/util"
	"project256/util"
)

func ActionBase() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 此处进行权限校验等
		if false == util.CheckUser("", "") {
			util.Exception(c, util.ERROR_USER_NOT_LOGIN)
			if c.IsAborted() {return}
		}

		// 返回用户信息
		//c.GetHeader("token")
		//c.Cookie("project256")
	}
}