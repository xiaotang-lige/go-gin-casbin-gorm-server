package linker

import (
	"github.com/redis/go-redis/v9"
	"net"
)

var RedisDb *redis.Client
var TcpConn net.Listener

func Main() {
	redisDb()
	tcpListen()
}
