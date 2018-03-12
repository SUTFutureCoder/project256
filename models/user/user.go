package user

func GetUserByIds(userIds []string) (map[string]interface{}, error) {
	// 目前只有一个1号
	userInfoList := make(map[string]interface{})
	userInfo := make(map[string]interface{})
	userInfo["user_id"] = "1"
	userInfo["user_name"] = "*Chen"
	userInfo["user_sign"] = ""
	userInfo["status"]  = "1"
	userInfoList["1"] = make(map[string]interface{})
	userInfoList["1"] = userInfo
	return userInfoList, nil
}