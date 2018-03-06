package wish

import (
	"github.com/gin-gonic/gin"
	"project256/util"
	"project256/models/wish"
	"project256/models/feed"
)

func WishList() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 获取用户id
		userId := c.Param("user_id")
		if userId == "" {
			// 一期必填，二期则改为推荐
			util.Exception(c, util.ERROR_PARAM_ERROR, "user_id不能为空")
			if c.IsAborted() {return}
		}
		ret := make(map[string]interface{})
		var err error
		ret["wish_list"], err = wish.GetListByUser(userId)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		util.Output(c, ret)
	}
}

func MakeAWish() (func(*gin.Context)) {
	return func(c *gin.Context) {
		data := make(map[string]interface{})
		data["parent_wish_id"] = c.DefaultPostForm("parent_wish_id", "")
		data["wish_content"], _ = c.GetPostForm("wish_content")
		if data["wish_content"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "心愿内容不能为空")
			if c.IsAborted() {return}
		}

		_, err := wish.InsertWish(&data)
		if err != nil {
			util.Exception(c, util.ERROR_DB_INSERT, err.Error())
			if c.IsAborted() {return}
		}
		feed.AddFeed(&data, util.TYPE_WISH)
		util.Output(c)
	}
}


