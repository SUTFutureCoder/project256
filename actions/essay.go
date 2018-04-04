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
	user  := new(models.UserStruct)
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
		ret := make(map[string]interface{})
		essayList, err := essay.GetListByUser(userId, perpage, (curpage - 1) * perpage)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		count, err := essay.GetListCountByUser(userId)

		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		// 附加心愿数据
		wishHash := make(map[string]bool)
		var wishIds []string
		// 附加用户数据
		userHash := make(map[string]bool)
		var userIds []string
		for _, v := range essayList {
			if wishHash[v.WishId] == false && v.WishId != "" {
				wishIds = append(wishIds, v.WishId)
				wishHash[v.WishId] = true
			}
			userIds = append(userIds, v.CreateUser)
			userHash[v.CreateUser] = true
		}

		// 回写
		wishInfoList, err := wish.GetWishByIds(wishIds)
		userInfoList, err := user.GetUserByIds(userIds)
		for i := 0; i < len(essayList); i++ {
			essayList[i].Ext = make(map[string]interface{})
			if wishInfoList[essayList[i].WishId].WishId != "" {
				essayList[i].Ext["wish_info"] = wishInfoList[essayList[i].WishId]
			}
			essayList[i].Ext["user_info"] = userInfoList[essayList[i].CreateUser]
		}

		ret["count"] = count
		ret["list"] = essayList
		util.Output(c, ret)
	}
}

func (e *Essay) EssayInfo() (func(*gin.Context)) {
	essay := new(models.EssayStruct)
	wish  := new(models.WishStruct)
	user  := new(models.UserStruct)
	return func(c *gin.Context) {
		// 获取文章id
		essayId := c.Param("essay_id")
		if essayId == "" {
			util.Exception(c, util.ERROR_PARAM_ERROR, "essayid不能为空")
			if c.IsAborted() {return}
		}
		var err error
		essayInfo, err := essay.GetInfo(essayId)
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		// 附加心愿数据
		var wishIds []string
		wishIds = append(wishIds, essayInfo.WishId)
		// 附加用户数据
		var userIds []string
		userIds = append(userIds, essayInfo.CreateUser)
		// 回写
		wishInfoList, err := wish.GetWishByIds(wishIds)
		userInfoList, err := user.GetUserByIds(userIds)
		essayInfo.Ext = make(map[string]interface{})
		if wishInfoList != nil && wishInfoList[essayInfo.WishId].WishId != "" {
			essayInfo.Ext["wish_info"] = wishInfoList[essayInfo.WishId]
		}
		essayInfo.Ext["user_info"] = userInfoList[essayInfo.CreateUser]
		util.Output(c, essayInfo)
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
		data["content_type"] = c.DefaultPostForm("content_type", "0")
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


