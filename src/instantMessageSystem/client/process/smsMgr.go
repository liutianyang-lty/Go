package process

import (
	"encoding/json"
	"fmt"
	"instantMessageSystem/common/message"
)

func outputGroupMes(mes *common.Message) {
	//显示即可
	//1. 先发序列化
	var smsMes common.SmsMes
	err :=json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t %d 对大家说\t %s", smsMes.UserId, smsMes.Content)
	fmt.Println(info)
}