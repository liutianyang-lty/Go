package model

import(
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后，就初始化一个userDao实例
//把它做成全局的变量，在需要和redis操作时，就直接使用
var (
	MyUserDao *UserDao
)

//UserDao结构体：完成对User结构体的各种操作。
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

//根据用户id返回一个User实例 + err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error){

	//通过给定的id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil { //表示在 users 哈希表中，没有查找到给用户
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	user = &User{}
	//将res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	return
}

//完成登录的校验
//1. Login 完成对用户的验证
//2. 如果用户的id和pwd都正确，则返回一个user实例
//3. 如果用户的id或pwd有错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao 的连接池中取出连接
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