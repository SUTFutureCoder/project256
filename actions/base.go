package actions

import (
	"github.com/gin-gonic/gin"
	//"project256/util"
)

func ActionBase() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 此处进行权限校验等
		//util.Exception(c, util.ERROR_PARAM_ERROR, "文章标题不能为空")
	}
}