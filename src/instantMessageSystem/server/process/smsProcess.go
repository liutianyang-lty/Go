package process

import(
	"encoding/json"
	"fmt"
	"instantMessageSystem/common/message"
	"instantMessageSystem/server/utils"
	"net"
)

type SmsProcess struct {
	//暂时没有字段
}

//转发消息
func (this *SmsProcess) SendGroupMes(msg *common.Message) {

	//遍历服务器端的onlineUsers map[int]*UserProcess
	//转发消息

	//1.取出mes中的内容
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(msg.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//为调用SendMesToEachOnlineUser()函数做准备
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		//将自己过滤掉
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Coon)
	}
}

//给每个在线用户发送消息
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, coon net.Conn) {

	//创建一个Transfer实例，发送data
	tf := &utils.Transfer{
		Coon : coon,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败 err=", err)
		return
	}
}