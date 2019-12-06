package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	common "instantMessageSystem/common/message"
)

//我们希望在服务器启动后，就初始化一个UserDao实例
//把它做成全局变量，在需要和redis操作时，直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool:pool,
	}
	return
}

//1.根据用户id 返回一个User实例+err
func (this *UserDao) getUserById(coon redis.Conn, id int) (user *User, err error) {

	//通过给定id 去 Redis查询这个用户
	res, err := redis.String(coon.Do("HGet", "users", id)) //注意这里一定要把从Redis里取出的数据转换为string类型
	if err != nil {
		//错误！
		if err == redis.ErrNil { //表示从Redis没有查询到数据
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//由于从Redis取出的数据是字符串，所以先反序列化
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登录的校验Login
//如果用户的id和pwd都正确，则返回一个user实例
//如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//从连接池中取出连接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//验证密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

//完成注册的校验Register
func (this *UserDao) Register(user *common.User) (err error) {

	//先从UserDao的连接池中取出连接
	coon := this.pool.Get()
	defer coon.Close()
	_, err = this.getUserById(coon, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//这时，说明用户还没注册过，则可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//入库
	_, err = coon.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}