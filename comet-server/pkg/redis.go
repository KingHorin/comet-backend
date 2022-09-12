package pkg

import (
	"github.com/go-redis/redis"
)

var rd *redis.Client

func init() {

	rd = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0, // redis一共16个库，指定其中一个库即可
	})
	_, err := rd.Ping().Result()
	if err != nil {
		panic("redis连接失败, " + err.Error())
	}

	rd.FlushDB() //每次清空所有token和黑名单记录

}

func GetRD() *redis.Client {
	return rd
}
