package actions

import (
	"github.com/gin-gonic/gin"
	"project256/util"
)

func ActionBase() (func(*gin.Context)) {
	return func(c *gin.Context) {
		token, _ := c.GetQuery("token")
		cookie := ""
		tmpCookie, err := c.Cookie("project256")
		if err != nil {
			cookie = tmpCookie
		}
		// 此处进行权限校验等
		if false == util.CheckUser(token, cookie) {
			util.Exception(c, util.ERROR_USER_NOT_LOGIN)
			if c.IsAborted() {return}
		}
	}
}