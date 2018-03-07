package essay

import (
	"project256/models"
	"project256/util"
	"fmt"
	"time"
)

func InsertEssay(essayData *map[string]interface{}) {
	user := util.GetUserInfo()
	essayId, err :=util.GenUUID32()
	db := models.GetDbConn()
	ret, err := db.Exec("INSERT INTO essay (essay_id, essay_title, essay_content, status, create_user, create_time) VALUES (?,?,?,?,?,?)",
			essayId,
			(*essayData)["essay_title"],
			(*essayData)["essay_content"],
			util.STATUS_VALID,
			user["user_id"],
			time.Now().Unix(),
		)
	fmt.Println(ret)
	fmt.Println(err)
}