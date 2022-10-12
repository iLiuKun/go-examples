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

	//编码 quicklist
	//仅当列表存在的时候才插入数据,此时列表不存在，无法插入
	redisDb.LPushX("studentList", "tom")

	//此时列表不存在，依然可以插入
	redisDb.LPush("studentList", "jack")

	//此时列表存在的时候才能插入数据
	redisDb.LPushX("studentList", "tom")

	// LPush支持一次插入任意个数据
	err = redisDb.LPush("studentList", "lily","lilei","zhangsan","lisi").Err()
	if err != nil {
		panic(err)
	}
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err := redisDb.LRange("studentList",0,-1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals) //注意列表是有序的，输出结果是[lisi zhangsan lilei lily tom jack]
}
