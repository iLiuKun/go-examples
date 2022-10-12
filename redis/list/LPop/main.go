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
	err = redisDb.LPush("studentList", "lily", "lilei", "zhangsan", "lisi", "tom", "lisi", "laowang").Err()
	if err != nil {
		panic(err)
	}
	vals, err := redisDb.LRange("studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("还未进行任何删除时列表中的值为:", vals) //[laowang lisi tom lisi zhangsan lilei lily]
	//从列表左边删除第一个数据，并返回删除的数据
	redisDb.LPop("studentList")
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err = redisDb.LRange("studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("进行了LPop删除操作后列表的值：", vals) //[lisi tom lisi zhangsan lilei lily]

	//删除列表中的数据。删除count个key的list中值为value 的元素
	//https://redis.io/commands/lrem/  For example, LREM list -2 "hello" will remove the last two occurrences of "hello" in the list stored at list.
	redisDb.LRem("studentList", 10, "lisi")
	vals, err = redisDb.LRange("studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("进行了LRem删除操作后列表的值：", vals) // [tom zhangsan lilei lily]
}


