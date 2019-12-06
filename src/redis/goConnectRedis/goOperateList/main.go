package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//Go操作List

	//连接
	coon, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	//写数据
	_, err = coon.Do("lpush", "heroList", "no1:宋江", 30, "no2:卢俊义", 28)
	if err != nil {
		fmt.Println("lpush err=", err)
		return
	}

	//读数据
	r, err := redis.String(coon.Do("rpop", "heroList"))
	if err != nil {
		fmt.Println("rpop err=", err)
		return
	}
	fmt.Println("rpop res=", r)
}
