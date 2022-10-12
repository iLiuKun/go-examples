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
	// 第三个参数代表key的过期时间，0代表不会过期。编码 int
	err = redisDb.Set("name1", 1, 0).Err()
	if err != nil {
		panic(err)
	}
	var oldVal string
	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	oldVal, err = redisDb.GetSet("name1", 3).Result()
	if err != nil {
		panic(err)
	}
	// 打印key的旧值
	fmt.Println("name1：", oldVal)

	//如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效，编码 emstr
	err = redisDb.SetNX("name2", "lisi", 0).Err()
	if err != nil {
		panic(err)
	}
}
