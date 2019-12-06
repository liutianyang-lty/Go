package common

//消息类型
const (
	LoginMesType			= "LoginMes"
	LoginResMesType     	= "LoginResMes"
	RegisterMesType     	= "RegisterMes"
	RegisterResMesType  	= "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType				= "SmsMes"
)

//定义用户状态的常量
const (
	UserOnline		= iota  //用户在线
	UserOffline				//用户离线
	UserBusyStatus			//用户繁忙
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"`//消息内容
}

//定义两个消息... 后面需要再增加

//登录消息
type LoginMes struct {
	UserId int `json:"userId"`//用户id
	UserPwd string `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}

//登录结果消息
type LoginResMes struct {
	Code int `json:"code"`//返回状态码 500表示该用户未注册 200表示登录成功
	UsersId []int        //增加字段，保存用户id的切片
	Error string `json:"error"`//返回错误信息
}

//注册消息
type RegisterMes struct {
	User User `json:"user"` //类型就是user结构体
}

//注册结果消息
type RegisterResMes struct {
	Code int `json:"code"` //状态码400表示用户已经占用，200表示注册成功
	Error string `json:"error"`
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户状态
}

//增加一个SmsMes 用于群发消息
type SmsMes struct {
	Content string `json:"content"`//消息内容
	User  //标识是哪个用户 (继承)
}