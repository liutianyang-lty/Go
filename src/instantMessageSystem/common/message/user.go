package common

//定义一个用户的结构体
type User struct {
	//确定字段信息
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"` //表示用户状态
}
