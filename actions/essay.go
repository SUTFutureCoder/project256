package actions

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"project256/util"
	"project256/models"
)

type Essay struct {
}

func (e *Essay) EssayList() (func(*gin.Context)) {
	essay := new(models.EssayStruct)
	wish  := new(models.WishStruct)
	return func(c *gin.Context) {
		// 获取用户id
		userId := c.Param("user_id")
		if userId == "" {
			// 一期必填，二期则改为推荐
			util.Exception(c, util.ERROR_PARAM_ERROR, "user_id不能为空")
			if c.IsAborted() {return}
		}
		curpage, _ := strconv.Atoi(c.DefaultQuery("curpage", "1"))
		perpage, _ := strconv.Atoi(c.DefaultQuery("perpage", "20"))
		if perpage > 20 {
			perpage = 20
		}
		if curpage < 1 {
			curpage = 1
		}
		var err error
		ret, err := essay.GetListByUser(userId, perpage, (curpage - 1) * perpage)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		// 附加心愿数据
		wishHash := make(map[string]bool)
		var wishIds []string
		for _, v := range ret {
			if wishHash[v.WishId] == false && v.WishId != "" {
				wishIds = append(wishIds, v.WishId)
				wishHash[v.WishId] = true
			}
		}
		// 回写
		wishInfoList, err := wish.GetWishByIds(wishIds)
		for i := 0; i < len(ret); i++ {
			ret[i].Ext = make(map[string]interface{})
			if wishInfoList[ret[i].WishId].WishId != "" {
				ret[i].Ext["wish_info"] = wishInfoList[ret[i].WishId]
			}
		}
		util.Output(c, ret)
	}
}

func (e *Essay) WriteEssay() (func(*gin.Context)) {
	essay := new(models.EssayStruct)
	feed  := new(models.FeedStruct)
	return func(c *gin.Context) {
		var err error
		data := make(map[string]string)
		data["essay_title"], _ = c.GetPostForm("essay_title")
		data["essay_content"], _ = c.GetPostForm("essay_content")
		data["wish_id"] = c.DefaultPostForm("wish_id", "")
		if data["essay_title"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "文章标题不能为空")
			if c.IsAborted() {return}
		}
		if len(data["essay_title"]) > 32 {
			util.Exception(c, util.ERROR_PARAM_ERROR, "文章标题需要小于32字符")
			if c.IsAborted() {return}
		}
		if data["essay_content"] == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "文章内容不能为空")
			if c.IsAborted() {return}
		}

		data["essay_id"], err = essay.InsertEssay(&data)
		if err != nil {
			util.Exception(c, util.ERROR_DB_INSERT, err.Error())
			if c.IsAborted() {return}
		}
		feed.AddFeed(&data, util.TYPE_ESSAY)
		util.Output(c)
	}
}


