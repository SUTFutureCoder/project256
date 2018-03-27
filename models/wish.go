package models

import (
	"project256/util"
	"time"
	"errors"
	"log"
	"fmt"
	"strings"
)

type WishStruct struct {
	Id			int64
	WishId		string
	ParentWishId 	string
	WishContent string
	Status 		int
	CreateUser	string
	CreateTime	int64
	UpdateTime	int64
	Ext 		map[string]interface{}
}

func (w *WishStruct) InsertWish(wishData *map[string]string) (string, error) {
	user := util.GetUserInfo()
	if user["status"] == util.STATUS_INVALID {
		return "", errors.New(util.GetErrorMessage(util.ERROR_USER_UNAUTHORIZED))
	}
	wishId, err :=util.GenUUID32()
	db := GetDbConn()
	ret, err := db.Exec("INSERT INTO wish (wish_id, parent_wish_id, wish_content, status, create_user, create_time) VALUES (?,?,?,?,?,?)",
		wishId,
		(*wishData)["parent_wish_id"],
		(*wishData)["wish_content"],
		util.STATUS_VALID,
		user["user_id"],
		time.Now().Unix(),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("Insert Wish Error: %s", err))
		return "", err
	}
	_, err = ret.RowsAffected()
	return wishId, err
}

func (w *WishStruct) GetListByUser(userId string) (*[]WishStruct, error){
	db := GetDbConn()
	ret, err := db.Query("SELECT * FROM wish WHERE create_user=?",
		userId,
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
	var wishData WishStruct
	var wishDataList []WishStruct
	for ret.Next() {
		err = ret.Scan(&wishData.Id, &wishData.WishId, &wishData.ParentWishId, &wishData.WishContent, &wishData.Status, &wishData.CreateUser, &wishData.CreateTime, &wishData.UpdateTime)
		if err != nil {
			log.Fatal(fmt.Sprintf("Scan Data Error: %s", err))
			return nil, err
		}
		wishDataList = append(wishDataList, wishData)
	}

	return &wishDataList, err
}

func (w *WishStruct) GetWishByIds(wishIds []string) (map[string]WishStruct, error) {
	var err error
	db := GetDbConn()
	ret, err := db.Query(fmt.Sprintf("SELECT * FROM wish WHERE wish_id IN (%s)",
		"'" + strings.Join(wishIds, "','") + "'",
	),
	)
	defer ret.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Get Wish Error: %s", err))
		return nil, err
	}
	// 初始化结构
	var wishData WishStruct
	wishDataList := make(map[string]WishStruct)
	for ret.Next() {
		err = ret.Scan(&wishData.Id, &wishData.WishId, &wishData.ParentWishId, &wishData.WishContent, &wishData.Status, &wishData.CreateUser, &wishData.CreateTime, &wishData.UpdateTime)
		if err != nil {
			log.Fatal(fmt.Sprintf("Scan Data Error: %s", err))
			return nil, err
		}
		wishDataList[wishData.WishId] = wishData
	}
	return wishDataList, err
}