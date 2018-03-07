package feed

import (
"github.com/gin-gonic/gin"
)

func FeedList() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 读取关注动态feed流

		// 默认读取1feed流，之后用户和关注系统出来后再改



		//ret := make(map[string]interface{})
		//var err error
		//ret["essay_list"], err = essay.GetListByUser(userId, 0, 20)
		//if err != nil {
		//	util.Exception(c, util.ERROR_DB_SELECT, err.Error())
		//	if c.IsAborted() {return}
		//}
		//util.Output(c, ret)
	}
}
