package linker

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net"
)

type db struct {
	RedisDb *redis.Client
	TcpConn net.Listener
	MysqlDb *gorm.DB
}

var Api = new(db)

func Main() {
	mysqlDb()
	redisDb()
	casbinDb()
	tcpListen()
}
