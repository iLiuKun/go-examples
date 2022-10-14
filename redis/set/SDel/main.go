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
	res1, _ := redisDb.SMembers("stuSet").Result()
	fmt.Println("res1=", res1) //res1= [100 200 300 400 500 600]

	//随机返回集合中的一个元素，并且删除这个元素,这里删除的是400
	member1, err := redisDb.SPop("stuSet").Result()
	fmt.Println(member1) //400

	res2, _ := redisDb.SMembers("stuSet").Result()
	fmt.Println("res2=", res2) //res2= [100 200 300 500 600]

	// 随机返回集合中的4个元素，并且删除这些元素
	member2, err := redisDb.SPopN("stuSet", 4).Result()
	fmt.Println(member2) //[]

	res3, _ := redisDb.SMembers("stuSet").Result()
	fmt.Println("res3=", res3) // [100 200 300 500 600]

	//删除集合stuSet名称为300,400的元素,并返回删除的元素个数
	member3, err := redisDb.SRem("stuSet", 500, 600).Result()
	fmt.Println(member3) //2
	res4, _ := redisDb.SMembers("stuSet").Result()
	fmt.Println("res4=", res4) //res4= [100 200 300]
}
