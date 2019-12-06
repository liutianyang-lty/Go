package goOperateString

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//Go连接到redis

	//操作String
	coon, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	defer coon.Close()

	//写数据
	_, err = coon.Do("Set", "name", "hello world")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	//读数据
	r, err := redis.String(coon.Do("Get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}

	fmt.Println("read result ", r)
}