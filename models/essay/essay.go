package essay

import (
	"project256/models"
	"project256/util"
	"fmt"
	"time"
	"log"
	"errors"
)

type EssayStruct struct {
	Id			int64
	EssayId		string
	EssayTitle 	string
	EssayContent string
	WishId		string
	Status 		int
	CreateUser	string
	CreateTime	int64
	UpdateTime	int64
}

func InsertEssay(essayData *map[string]interface{}) (int64, error) {
	user := util.GetUserInfo()
	if user["status"] == util.STATUS_INVALID {
		return 0, errors.New(util.GetErrorMessage(util.ERROR_USER_UNAUTHORIZED))
	}
	essayId, err :=util.GenUUID32()
	db := models.GetDbConn()
	ret, err := db.Exec("INSERT INTO essay (essay_id, essay_title, essay_content, wish_id, status, create_user, create_time) VALUES (?,?,?,?,?,?,?)",
			essayId,
			(*essayData)["essay_title"],
			(*essayData)["essay_content"],
			(*essayData)["wish_id"],
			util.STATUS_VALID,
			user["user_id"],
			time.Now().Unix(),
		)
	if err != nil {
		log.Fatal(fmt.Sprintf("Insert Essay Error: %s", err))
		return 0, err
	}
	row, _ := ret.RowsAffected()
	return row, err
}

func GetListByUser(userId string, offset, limit int) (*[]EssayStruct, error){
	db := models.GetDbConn()
	ret, err := db.Query("SELECT * FROM essay WHERE create_user=? ORDER BY id DESC LIMIT ? OFFSET ?",
			userId,
			limit,
			offset,
		)
	defer ret.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Get List By User Error: %s", err))
		return nil, err
	}
	_, err = ret.Columns()
	if err != nil {
		log.Fatal(fmt.Sprintf("Get List By User Error: %s", err))
		return nil, err
	}

	// 初始化结构
	var essayData EssayStruct
	var essayDataList []EssayStruct
	if ret.Next() == false {
		return &essayDataList, nil
	}
	for ret.Next() {
		err = ret.Scan(&essayData.Id, &essayData.EssayId, &essayData.EssayTitle, &essayData.EssayContent, &essayData.WishId, &essayData.Status, &essayData.CreateUser, &essayData.CreateTime, &essayData.UpdateTime)
		if err != nil {
			log.Fatal(fmt.Sprintf("Scan Data Error: %s", err))
			return nil, err
		}
		essayDataList = append(essayDataList, essayData)
	}

	return &essayDataList, err
}