package main

import (
	"fmt"
	common "instantMessageSystem/common/message"
	process2 "instantMessageSystem/server/process"
	"instantMessageSystem/server/utils"
	"io"
	"net"
)

//先创建一个Processor结构体
type Processor struct{
	//保存由main包中的协程传递过来的TCP连接
	Coon net.Conn
}

//编写函数ServerProcessMes函数：根据客户端发送的不同种类消息，决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *common.Message) (err error) {

	//看看是否能接受到客户端发送的群发消息
	fmt.Println("mes=", mes)

	//将消息分发给相应的控制器去处理
	switch  mes.Type {
		case common.LoginMesType:

			//处理登录
			//创建一个UserProcess实例
			up := &process2.UserProcess{
				Coon : this.Coon,
			}
			err = up.ServerProcessLogin(mes)

		case common.RegisterMesType:

			//处理注册
			//创建一个UserProcess实例
			up := &process2.UserProcess{
				Coon : this.Coon,
			}
			err = up.ServerProcessRegister(mes)
		case common.SmsMesType:
			//处理客户端过来的群发消息
			smsProcess := &process2.SmsProcess{}
			smsProcess.SendGroupMes(mes)
		default :
			fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环接收客户端发送的消息
	for {

		//创建一个Transfer 实例
		tf := &utils.Transfer{
			Coon: this.Coon,
		}

		//接收到客户端发送来的消息并解析出来(看属于哪种类型的消息) 交给调度中心处理
		mes, err := tf.ReadPkg()

		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器端也退出...")
				return err
			} else {
				fmt.Println("readPkg err=", err)
				return err
			}

		}

		//fmt.Println("mes=", mes)
		//将消息交给业务调度中心去处理
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}