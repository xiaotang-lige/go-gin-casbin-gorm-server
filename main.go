package main

import (
	"messageServe/dao"
	"messageServe/linker"
	"messageServe/route"
	"messageServe/server"
	_ "net/http/pprof"
)

func main() {
	linker.Main()
	dao.Main()
	server.Main()
	route.Main()
}
