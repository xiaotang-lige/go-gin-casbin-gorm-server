package linker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func redisDb() {
	Api.RedisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if err := Api.RedisDb.Ping(context.Background()).Err(); err != nil {
		log.Println("redis连接失败!", err)
		return
	}
	log.Println("redis连接成功！")
}
func mysqlDb() {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/messageServer?charset=utf8mb4&parseTime=True&loc=Local"
	Api.MysqlDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("mysql連接出錯：", err)
	}
	log.Println("mysql连接成功！")
}
