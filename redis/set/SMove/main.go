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
	err = redisDb.SAdd("stuSet", 100, 200, 300, 300,400, 500, 600).Err()
	err = redisDb.SAdd("resSet", 900).Err()
	if err != nil {
		panic(err)
	}
	//把集合里的元素转换成map的key
	map1, err:= redisDb.SMembersMap("stuSet").Result()
	fmt.Println(map1) //map[100:{} 200:{} 300:{} 400:{} 500:{} 600:{}]

	//移动集合stuSet中的一个200元素到集合resSet中去
	ok, err:= redisDb.SMove("stuSet", "resSet", 200).Result()
	if ok{
		fmt.Println("移动数据成功")
	}
	//返回resSet集合中的所有元素
	resSetStr, err := redisDb.SMembers("resSet").Result()
	fmt.Println(resSetStr) //[200 900]
	stSetStr, err := redisDb.SMembers("stuSet").Result()
	fmt.Println(stSetStr) //[200 900]
}
