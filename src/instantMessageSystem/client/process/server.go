package process

import (
	"encoding/json"
	"fmt"
	"instantMessageSystem/client/utils"
	"instantMessageSystem/common/message"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("----------恭喜xxx登录成功----------")
	fmt.Println("----------1. 显示在线用户列表----------")
	fmt.Println("----------2. 发送消息----------")
	fmt.Println("----------3. 信息列表----------")
	fmt.Println("----------4. 退出系统----------")
	fmt.Println("请选择(1-4):")

	var key int
	var content string
	var smsProcess = &SmsProcess{}

	fmt.Scanf("%d\n", &key)
	switch key {
		case 1:
			//fmt.Println("显示在线用户列表")
			outputOnlineUser()
		case 2:
			fmt.Println("请输入你想对大家说的话:)")
			fmt.Scanf("%s\n", &content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出了系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不对...")
	}

}

//和服务器端保持通讯
func serverProcessMes(Coon net.Conn) {

	//创建一个Transfer实例, 不停读取服务器发送的消息
	tf := &utils.Transfer{
		Coon:Coon,
	}

	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}

		//如果读取到消息，执行下一步逻辑
		switch mes.Type {
			case common.NotifyUserStatusMesType: //有人上线

				//1.取出notifyUserStatusMes
				var notifyUserStatusMes common.NotifyUserStatusMes
				json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				//2.把这个用户的信息、状态保存到客户端的map[int]User中
				updateUserStatus(&notifyUserStatusMes)
			case common.SmsMesType: //有人群发消息
				outputGroupMes(&mes)
			default:
				fmt.Println("服务器端返回一个未知类型消息")
		}

		//fmt.Printf("mes=%v", mes)
	}
}