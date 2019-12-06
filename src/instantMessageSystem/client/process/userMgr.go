package process

import (
	"fmt"
	_"fmt"
	"instantMessageSystem/client/model"
	"instantMessageSystem/common/message"
)

//客户端要维护的map
var onlineUsers map[int]*common.User = make(map[int]*common.User, 10)
var CurUser model.CurUser //在用户登录成功后，完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户ID：\t", id)
	}
}

//编写一个函数，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *common.NotifyUserStatusMes) {

	//适当优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &common.User{
			UserId : notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}