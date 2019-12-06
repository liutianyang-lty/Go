package process

import (
	"encoding/json"
	"fmt"
	"instantMessageSystem/server/model"
	common "instantMessageSystem/common/message"
	"instantMessageSystem/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Coon net.Conn
	//增加一个字段，表示该Coon属于哪个用户
	UserId int
}

//编写通知所有在线的用户的方法
//这个userId的人 通知其他的在线用户，他上线了
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历onlineUsers,然后挨个发送通知 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤自己
		if id == userId {
			continue
		}

		//开始通知(封装)
		up.NotifyMeOnline(userId)
	}
}

func (this *UserProcess) NotifyMeOnline(userId int) {

	//组装我们的NotifyUserStatusMes
	var mes common.Message
	mes.Type = common.NotifyUserStatusMesType

	var notifyUserStatusMes common.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = common.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMes) err=", err)
		return
	}
	mes.Data = string(data)

	//对mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}

	//发送消息
	tf := &utils.Transfer{
		Coon : this.Coon,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

//编写函数ServerProcessLogin 函数，处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {

	//1.先从mes中取出mes.Data 并直接序列化成LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//声明一个响应信息
	var resMes common.Message
	resMes.Type = common.LoginResMesType

	//再声明一个 LoginResMes
	var loginResMes common.LoginResMes

	//到redis数据库完成验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_PWD {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		//登录成功后，把登录成功的用户放入userMgr中
		//将登录成功的用户userId 赋给 this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其他用户我上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)

		//将当前在线用户的id，放入到LoginResMes中的UsersId切片中
		//遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}

		fmt.Println(user, "登录成功")
	}

	//序列化loginResMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) fail err =", err)
		return
	}

	resMes.Data = string(data)

	//5. 将resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err =", err)
		return
	}

	//6. 发送data，调用封装函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Coon:this.Coon,
	}
	err = tf.WritePkg(data)
	return
}

//编写函数ServerProcessRegister， 处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *common.Message) (err error) {

	//1.先从mes中取出mes.Data 并直接序列化成RegisterMes
	var registerMes common.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//声明一个响应信息
	var resMes common.Message
	resMes.Type = common.RegisterResMesType

	//再声明一个 LoginResMes
	var registerResMes common.RegisterResMes

	//到redis数据库完成验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
		fmt.Println("注册成功")
	}

	//序列化loginResMes
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal(registerResMes) fail err =", err)
		return
	}

	resMes.Data = string(data)

	//5. 将resMes序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) fail err =", err)
		return
	}

	//6. 发送data，调用封装函数
	//因为使用分层模式，我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Coon:this.Coon,
	}
	err = tf.WritePkg(data)
	return
}