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
	// LPush支持一次插入任意个数据
	err = redisDb.LPush("studentList", "lily", "lilei", "zhangsan", "lisi", "tom").Err()
	if err != nil {
		panic(err)
	}
	//给名称为key的list中index位置的元素赋值，把原来的数据覆盖
	redisDb.LSet("studentList", 2, "beer")
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err := redisDb.LRange("studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals) //注意列表是[tom lisi beer lilei lily]

	//在list列表studentList中值为lilei前面添加元素hello
	redisDb.LInsert("studentList", "before", "lilei", "hello")
	//redisDb.LInsertBefore("studentList","lilei","hello") 执行效果同22行

	//在list列表studentList中值为tom后面添加元素world
	redisDb.LInsert("studentList", "after", "tom", "world")
	//redisDb.LInsertAfter("studentList","tom","world") 执行效果同26行

	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err = redisDb.LRange("studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals) //[tom world lisi beer hello lilei lily]
}


