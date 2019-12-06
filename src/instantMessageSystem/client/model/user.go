package model

//定义一个用户的结构体
type User struct {
	//为了序列化和反序列化成功，必须加上tag
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName`
}
