package essay

import (
	"github.com/gin-gonic/gin"
	"project256/util"
	"project256/models/essay"
	"project256/models/feed"
)

func EssayList() (func(*gin.Context)) {
	return func(c *gin.Context) {
		// 获取用户id
		userId := c.Param("user_id")
		if userId == "" {
			// 一期必填，二期则改为推荐
			util.Exception(c, util.ERROR_PARAM_ERROR, "user_id不能为空")
			if c.IsAborted() {return}
		}
		limit := c.DefaultQuery("limit", "20")
		offset := c.DefaultQuery("offset", "0")
		ret := make(map[string]interface{})
		var err error
		ret["essay_list"], err = essay.GetListByUser(userId, limit, offset)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		util.Output(c, ret)
	}
}

func WriteEssay() (func(*gin.Context)) {
	return func(c *gin.Context) {
		data := make(map[string]interface{})
		data["essay_title"], _ = c.GetPostForm("essay_title")
		data["essay_content"], _ = c.GetPostForm("essay_content")
		data["wish_id"] = c.DefaultPostForm("wish_id", "")
		if data["essay_title"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "文章标题不能为空")
			if c.IsAborted() {return}
		}
		if tmpEssayTitle, ok := data["essay_title"].(string); ok {
			if len(tmpEssayTitle) > 32 {
				util.Exception(c, util.ERROR_PARAM_ERROR, "文章标题需要小于32字符")
				if c.IsAborted() {return}
			}
		}
		if data["essay_content"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "文章内容不能为空")
			if c.IsAborted() {return}
		}

		_, err := essay.InsertEssay(&data)
		if err != nil {
			util.Exception(c, util.ERROR_DB_INSERT, err.Error())
			if c.IsAborted() {return}
		}
		feed.AddFeed(&data, util.TYPE_ESSAY)
		util.Output(c)
	}
}


