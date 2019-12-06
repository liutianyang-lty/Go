package process

import(
	"encoding/json"
	"fmt"
	"instantMessageSystem/client/utils"
	"instantMessageSystem/common/message"
)

type SmsProcess struct {

}

//发送群聊的消息
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	//1.创建一个Mes
	var mes common.Message
	mes.Type = common.SmsMesType

	//2. 创建一个SmsMes实例
	var smsMes common.SmsMes
	smsMes.Content = content //内容
	smsMes.UserId = CurUser.UserId //用户id
	smsMes.UserStatus = CurUser.UserStatus //用户状态

	//3. 序列化
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(smsMes) fail=", err.Error())
		return
	}

	mes.Data = string(data)

	//4. 对mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal(mes) fail=", err.Error())
		return
	}

	//5. 发送mes
	tf := &utils.Transfer{
		Coon:CurUser.Coon,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}