package actions

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"project256/util"
	"project256/models"
)

type Feed struct {}

func (f *Feed) FeedList() (func(*gin.Context)) {
	feed := new(models.FeedStruct)
	user := new(models.UserStruct)
	return func(c *gin.Context) {
		// 读取关注动态feed流
		// 默认读取1feed流，之后用户和关注系统出来后再改
		curpage, _ := strconv.Atoi(c.DefaultQuery("curpage", "1"))
		perpage, _ := strconv.Atoi(c.DefaultQuery("perpage", "20"))
		if perpage > 20 {
			perpage = 20
		}
		ret, err 		:= feed.GetFeed(perpage, (curpage - 1) * perpage)
		// 关联查询
		userIdsHash := make(map[string]bool)
		var userIds []string
		// 输出
		for _, v := range *ret {
			if userIdsHash[v.CreateUser] == false {
				userIds = append(userIds, v.CreateUser)
				userIdsHash[v.CreateUser] = true
			}
		}
		// 回写
		userInfoList, err := user.GetUserByIds(userIds)
		for i := 0; i < len(*ret); i++ {
			(*ret)[i].Ext = make(map[string]interface{})
			if userInfoList[(*ret)[i].CreateUser] != false {
				(*ret)[i].Ext["user_info"] = userInfoList[(*ret)[i].CreateUser]
			}
		}

		if err != nil {
			util.Exception(c, util.ERROR_DB_SELECT, err.Error())
			if c.IsAborted() {return}
		}
		util.Output(c, ret)
	}
}
