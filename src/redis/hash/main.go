package hash

func main() {

	//Redis--hash
	//hash类似于golang里的Map
	//Redis--hash 是一个键值对集合，是一个string类型的field和value的映射表
	//hash特别适合用于存储对象

	//使用细节
	//hash的CRUD操作：hget/hset/hdel
	//1. hgetall：一次获取所有的字段
	//2. hmset：一次设置多个field值
	//3. hexists：判断hash中是否存在某个field
	//4. hmget：一次获取多个field值
	//5. hlen：统计一个hash有多少个键值对
}
