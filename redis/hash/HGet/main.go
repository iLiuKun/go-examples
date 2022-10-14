package main
import (
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的redisDb变量
var redisDb *redis.Client

// 根据redis配置初始化一个客户端
func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	//通过 *redis.Client.Ping() 来检查是否成功连接到了redis服务器
	_, err = redisDb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	//根据key和field字段设置,field字段的值。 user_1 是hash key，username 是字段名, admin是字段值
	err = redisDb.HSet("user_1", "username", "admin").Err()
	if err != nil {
		panic(err)
	}
	//根据key和field字段,查询field字段的值。user_1 是hash key，username是字段名
	username, err := redisDb.HGet("user_1", "username").Result()

	if err != nil {
		panic(err)
	}
	fmt.Println(username) //admin
	//继续往user_1中添加字段password
	_ = redisDb.HSet("user_1", "password", "abc123").Err()

	// HGetAll 一次性返回key=user_1的所有hash字段和值
	data, err := redisDb.HGetAll("user_1").Result()
	if err != nil {
		panic(err)
	}
	// data是一个map类型，这里使用使用循环迭代输出
	for field, val := range data {
		fmt.Println(field, val)
	}

	// 初始化hash数据的多个字段值
	batchData := make(map[string]interface{})
	batchData["username"] = "test"
	batchData["password"] = 123456
	// 一次性保存多个hash字段值
	err = redisDb.HMSet("user_2", batchData).Err()
	if err != nil {
		panic(err)
	}

	//如果email字段不存在，则设置hash字段值
	redisDb.HSetNX("user_2", "email", "ourlang@foxmail.com")
	// HMGet支持多个field字段名，意思是一次返回多个字段值
	values, err := redisDb.HMGet("user_2", "username", "password", "email").Result()
	if err != nil {
		panic(err)
	}
	// values是一个数组
	fmt.Println("user_2=", values) //user_2= [test 123456 ourlang@foxmail.com]
}
