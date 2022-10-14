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

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func main() {
	err := initClient()
	if err != nil {
		//redis连接错误
		panic(err)
	}
	// 将100,200,300批量添加到集合中 集合的元素不能重复
	err = redisDb.SAdd("stuSet", 100, 200, 300, 400, 500, 600).Err()
	if err != nil {
		panic(err)
	}
	//随机返回集合stuSet中的一个元素
	member1, _ := redisDb.SRandMember("stuSet").Result()
	fmt.Println(member1) //600
	//随机返回集合stuSet中的3个元素
	member2, _ := redisDb.SRandMemberN("stuSet", 3).Result()
	fmt.Println(member2) //[600 400 500]
}
