package essay

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project256/util"
	//"project256/models/essay"
)

func EssayList() (func(*gin.Context)) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "asdfsf",
		})
	}
}

func AddEssay() (func(*gin.Context)) {
	return func(c *gin.Context) {
		data := make(map[string]interface{})
		data["essay_title"], _ = c.GetPostForm("essay_title")
		data["essay_content"], _ = c.GetPostForm("essay_content")
		data["wish_id"] = c.DefaultPostForm("wish_id", "0")
		data["little_wish_id"] = c.DefaultPostForm("little_wish_id", "0")
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

		util.Output(c, nil)
	}
}


