package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	common "instantMessageSystem/common/message"
	"net"
)

/**
 * 这个包的作用是进行数据传输(数据包解析)
 */
type Transfer struct {
	//分析它应该有哪些字段
	Coon net.Conn
	Buf [8096]byte  //这是传输时的缓冲
}

//读取数据包的函数
func (this *Transfer) ReadPkg() (mes common.Message, err error) {

	fmt.Println("等待读取客户端发送的数据...")

	//coon.Read 在coon没有被关闭的时候，才会阻塞
	//如果客户端关闭了coon, 则就不会阻塞了

	_, err = this.Coon.Read(this.Buf[:4])

	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//将buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据 pkgLen 读取消息内容
	n, err := this.Coon.Read(this.Buf[:pkgLen]) //从coon里读取pkgLen长度的字节，放到buf里
	if uint32(n) != pkgLen || err != nil {
		//err = errors.New("read pkg body error")
		return
	}

	//把pkgLen 反序列化成 -> common.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes) //符号&一定要带上
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

//写数据包
func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送长度给对方
	pkgLen := uint32(len(data))
	//var buf [4]byte

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Coon.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("coon.Write(buf) err=", err)
		return
	}

	//发送data本身
	n, err = this.Coon.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("coon.Write(data) err=", err)
		return
	}
	return
}