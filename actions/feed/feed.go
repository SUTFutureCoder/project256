package feed

import (
	"github.com/gin-gonic/gin"
	"project256/models/feed"
	"project256/util"
)

func FeedList() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 读取关注动态feed流
		// 默认读取1feed流，之后用户和关注系统出来后再改
		limit := c.DefaultQuery("limit", "20")
		offset := c.DefaultQuery("offset", "0")
		ret, err := feed.GetFeed(limit, offset)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		util.Output(c, ret)
	}
}
