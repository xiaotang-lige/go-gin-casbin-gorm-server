package linker

import (
	"fmt"
	"log"
	"net"
)

func tcpListen() {
	var err error
	Db.TcpConn, err = net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("TcpListen连接失败:", err)
		return
	}
	log.Println("tcpListen启动成功！")
}
