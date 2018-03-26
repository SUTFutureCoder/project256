package actions

import (
	"html"
	"github.com/gin-gonic/gin"
	"project256/util"
	"project256/models"
)

type Wish struct{}

func (w *Wish) WishList() (func(*gin.Context)) {
	wish := new(models.WishStruct)
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

func (w *Wish) MakeAWish() (func(*gin.Context)) {
	wish := new(models.WishStruct)
	feed := new(models.FeedStruct)
	return func(c *gin.Context) {
		var err error
		data := make(map[string]string)
		data["parent_wish_id"] = c.DefaultPostForm("parent_wish_id", "")
		data["wish_content"], _ = c.GetPostForm("wish_content")
		if data["wish_content"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "心愿内容不能为空")
			if c.IsAborted() {return}
		}

		data["wish_content"] = html.EscapeString(data["wish_content"])

		data["wish_id"], err = wish.InsertWish(&data)
		if err != nil {
			util.Exception(c, util.ERROR_DB_INSERT, err.Error())
			if c.IsAborted() {return}
		}
		feed.AddFeed(&data, util.TYPE_WISH)
		util.Output(c)
	}
}


