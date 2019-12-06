package main

import(
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {

	//Go连接到redis

	//操作Hash
	coon, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	defer coon.Close()

	//写数据
	//_, err = coon.Do("HSet", "user01", "name", "hahha")
	//if err != nil {
	//	fmt.Println("set err=", err)
	//	return
	//}
	//
	////读数据
	//r, err := redis.String(coon.Do("HGet", "user01", "name"))
	//if err != nil {
	//	fmt.Println("get err=", err)
	//	return
	//}

	//批量写数据
	_, err = coon.Do("HMSet", "user02", "name", "john", "age", 39)
	if err != nil {
		fmt.Println("HMSet err=", err)
		return
	}

	//返回的结果相当于切片
	r, err := redis.Strings(coon.Do("HMGet", "user02", "name", "age"))
	if err != nil {
		fmt.Println("HMGet err=", err)
		return
	}
	for i, v :=  range r {
		fmt.Printf("r[%d] = %s\n", i, v)
	}
	//fmt.Println("read result ", r)
}