package util

var userInfo map[string]interface{}

func CheckUser(token, cookie string) (bool) {
	userInfo = make(map[string]interface{})
	// 这里验证cookies，并且核对token
	cfg, _ := GetConfig("secret")
	if cfg["token"] == token {
		// 这里写入用户信息
		userInfo["user_id"] = "1"
		userInfo["user_name"] = "*Chen"
		userInfo["user_sign"] = ""
		userInfo["status"] = STATUS_VALID
	} else {
		userInfo["user_id"] = ""
		userInfo["user_name"] = ""
		userInfo["user_sign"] = ""
		userInfo["status"] = STATUS_INVALID
	}
	return true
}

func GetUserInfo() (map[string]interface{}) {
	return userInfo
}