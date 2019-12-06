package model

import(
	"instantMessageSystem/common/message"
	"net"
)

//因为在客户端，我们很多地方会使用到curUser,我们将其作为一个全局变量
type CurUser struct {
	Coon net.Conn
	common.User
}
