package util

var userInfo map[string]interface{}

func CheckUser(token, cookie string) (bool) {
	// 这里验证cookies，并且核对token

	// 这里写入用户信息
	userInfo = make(map[string]interface{})
	userInfo["user_id"] = "1"
	userInfo["user_name"] = "FutureCoder"
	userInfo["user_sign"] = ""
	return true
}

func GetUserInfo() (map[string]interface{}) {
	return userInfo
}