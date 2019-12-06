package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个全局pool
var pool *redis.Pool

//初始化方法
func init() {

	pool = &redis.Pool{
		MaxIdle: 8, //最大空闲连接数
		MaxActive: 0, //与数据库的最大连接数，0表示没有限制
		IdleTimeout: 100, //空闲连接等待时间
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	//Redis连接池
	//说明:
	//1. 事先初始化一定数量的连接，放入到连接池
	//2. 当Go需要Redis时，直接从Redis连接池取出连接即可
	//3. 这样可以节省临时获取Redis的时间，从而提高效率

	//Redis连接池案例

	//从连接池取出连接
	coon := pool.Get()
	defer coon.Close()

	//写数据
	_, err := coon.Do("Set", "name", "Golang")
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

	fmt.Println("get res=", r)

	//如果我们要从pool中取出连接，一定要保证连接池没有关闭

}
