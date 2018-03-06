package feed

import (
	"project256/models"
	"project256/util"
	"fmt"
	"time"
	"log"
	"errors"
	"strings"
)

type FeedStruct struct {
	Id			int64
	FeedId		string
	FeedData 	string
	FeedType 	int
	Status 		int
	CreateUser	string
	CreateTime	int64
}

func AddFeed(data *map[string]interface{}, dataType int) {
	feedData := make(map[string]interface{})
	feedData["feed_type"] = dataType
	switch dataType {
	case util.TYPE_ESSAY:
		feedData["feed_data"] = (*data)["essay_title"]
	case util.TYPE_WISH:
		feedData["feed_data"] = (*data)["wish_content"]
	}
	InsertFeed(&feedData)
}

func InsertFeed(feedData *map[string]interface{}) (int64, error) {
	user := util.GetUserInfo()
	if user["status"] == util.STATUS_INVALID {
		return 0, errors.New(util.GetErrorMessage(util.ERROR_USER_UNAUTHORIZED))
	}
	feedId, err :=util.GenUUID32()
	db := models.GetDbConn()
	ret, err := db.Exec("INSERT INTO feed (feed_id, feed_data, feed_type, status, create_user, create_time) VALUES (?,?,?,?,?,?)",
		feedId,
		(*feedData)["feed_data"],
		(*feedData)["feed_type"],
		util.STATUS_VALID,
		user["user_id"],
		time.Now().Unix(),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("Insert Feed Error: %s", err))
		return 0, err
	}
	row, _ := ret.RowsAffected()
	return row, err
}

func GetFeed(limit, offset string) (*[]FeedStruct, error){
	// 获取用户信息
	_ = util.GetUserInfo()
	// 获取用户关注的账号 1期不做
	var userSubscribe []string
	userSubscribe = append(userSubscribe, "1", "2")
	// 转为逗号分隔 注意为空的情况
	db := models.GetDbConn()
	ret, err := db.Query(fmt.Sprintf("SELECT * FROM feed WHERE create_user IN (%s) ORDER BY id DESC LIMIT ? OFFSET ?",
			"'" + strings.Join(userSubscribe, "','") + "'",
		),
		limit,
		offset,
	)
	defer ret.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Get Feed Error: %s", err))
		return nil, err
	}
	_, err = ret.Columns()
	if err != nil {
		log.Fatal(fmt.Sprintf("Get Feed Error: %s", err))
		return nil, err
	}

	// 初始化结构
	var feedData FeedStruct
	var feedDataList []FeedStruct
	for ret.Next() {
		err = ret.Scan(&feedData.Id, &feedData.FeedId, &feedData.FeedData, &feedData.FeedType, &feedData.Status, &feedData.CreateUser, &feedData.CreateTime)
		if err != nil {
			log.Fatal(fmt.Sprintf("Scan Data Error: %s", err))
			return nil, err
		}
		feedDataList = append(feedDataList, feedData)
	}
	return &feedDataList, err
}