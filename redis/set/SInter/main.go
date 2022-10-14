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
	redisDb.SAdd("blacklist", "Obama")     // 向 blacklist 中添加元素
	redisDb.SAdd("blacklist", "Hillary")   // 再次添加
	redisDb.SAdd("blacklist", "the Elder") // 添加新元素

	redisDb.SAdd("whitelist", "the Elder") // 向 whitelist 添加元素

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	names, err := redisDb.SInter("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("交集结果是: ", names) // [the Elder]

	//求交集并将交集保存到 destSet 的集合
	res, err := redisDb.SInterStore("destSet", "blacklist", "whitelist").Result()
	fmt.Println(res)
	//获取交集的值[the Elder]
	destStr, _ := redisDb.SMembers("destSet").Result()
	fmt.Println(destStr) //[the Elder]

	// 求差集
	diffStr, err := redisDb.SDiff("blacklist", "whitelist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("差集结果是: ", diffStr) //[Hillary Obama]
}
