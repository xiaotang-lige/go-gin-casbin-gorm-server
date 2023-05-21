package linker

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"net"
)

var RedisDb *redis.Client
var TcpConn net.Listener
var MysqlDb *gorm.DB

func Main() {
	redisDb()
	tcpListen()
	mysqlDb()
}
