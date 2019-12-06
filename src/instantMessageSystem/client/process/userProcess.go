package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"instantMessageSystem/client/utils"
	common "instantMessageSystem/common/message"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要任何字段
}

//用户登录功能
func (this *UserProcess) Login(userId int, userPwd string) (err error) { //最好用error来描述返回信息

	//1. 连接到服务器
	coon, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//延时关闭
	defer coon.Close()

	//2. 通过coon发送消息给服务
	var mes common.Message
	mes.Type = common.LoginMesType

	//3. 创建一个LoginMes 结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4. 将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	//5. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//6. 发送消息
	//6.1 先把data的长度发送给服务器
	//先获取到data的长度，然后将长度转为byte数组
	pkgLen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := coon.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("coon.Write(buf) err=", err)
		return
	}

	fmt.Printf("客户端, 发送消息的长度=%d 内容=%s\n", len(data), string(data))

	//发送消息本身
	_, err = coon.Write(data)
	if err != nil {
		fmt.Println("coon.Write(data) fail", err)
		return
	}

	//这里还需要处理服务器端返回的消息
	tf := &utils.Transfer{
		Coon:coon,
	}
	mes, err = tf.ReadPkg()

	if err != nil {
		fmt.Println("readPkg(coon) err=", err)
		return
	}
	//将mes的Data部分反序列化成 LoginResMes
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Coon = coon
		CurUser.UserId = userId
		CurUser.UserStatus = common.UserOnline

		//显示当前在线用户列表，遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {

			//如果要求不显示自己在线
			if v == userId {
				continue
			}
			fmt.Printf("用户id：%d\t", v)

			//完成客户端onlineUsers的初始化
			user := &common.User{
				UserId : v,
				UserStatus : common.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里我们需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接受并显示在客户端的终端
		go serverProcessMes(coon)

		//1. 显示我们登录成功后的菜单[循环]
		for {
			ShowMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}

//用户注册功能
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1. 连接到服务器
	coon, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	//延时关闭
	defer coon.Close()

	//2. 通过coon发送消息给服务
	var mes common.Message
	mes.Type = common.RegisterMesType

	//3. 创建一个RegisterMes 结构体
	var registerMes common.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将registerMes 序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)

	//5. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer实例
	tf := &utils.Transfer{
		Coon:coon,
	}

	//发送data给服务器
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(coon) err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，你重新登录一把")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}
