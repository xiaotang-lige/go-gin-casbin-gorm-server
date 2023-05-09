package linker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func redisDb() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := RedisDb.Ping(context.Background()).Err(); err != nil {
		log.Println("redis连接失败!", err)
		return
	}
	log.Println("redis连接成功！")
}
