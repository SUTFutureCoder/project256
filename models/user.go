package models

type UserStruct struct {
}

func (u *UserStruct) GetUserByIds(userIds []string) (map[string]interface{}, error) {
	// 目前只有一个1号
	userInfoList := make(map[string]interface{})
	userInfo := make(map[string]interface{})
	userInfo["user_id"] = "1"
	userInfo["user_name"] = "FutureCoder"
	userInfo["user_sign"] = "Progressive Developer"
	userInfo["user_avatar"] = "https://img.waimai.baidu.com/pc/0c65e42b51064007812b4988a0074438b9"
	userInfo["status"]  = "1"
	userInfoList["1"] = make(map[string]interface{})
	userInfoList["1"] = userInfo
	return userInfoList, nil
}