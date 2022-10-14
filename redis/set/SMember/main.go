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
	// 添加100到集合中
	err = redisDb.SAdd("stuSet", 100).Err()
	if err != nil {
		panic(err)
	}
	// 将100,200,300批量添加到集合中 集合的元素不能重复
	redisDb.SAdd("stuSet", 100, 200, 300)

	//返回名称为集合中的所有元素
	es, err := redisDb.SMembers("stuSet").Result()
	fmt.Println(es) //[100 200 300]
	fmt.Println(typeof(es))

	//此处flag1=true
	flag1, err := redisDb.SIsMember("stuSet", "200").Result()
	if flag1 {
		fmt.Println("集合stuSet中包含指定元素200") //输出
	} else {
		fmt.Println("集合stuSet不包含元素200")
	}
	//此处flag2=false
	flag2, err := redisDb.SIsMember("stuSet", "400").Result()
	if flag2 {
		fmt.Println("集合stuSet中包含指定元素400")
	} else {
		fmt.Println("集合stuSet不包含元素400") //输出
	}
}

