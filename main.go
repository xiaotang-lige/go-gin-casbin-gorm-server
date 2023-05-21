package main

import (
	"messageServe/linker"
	"messageServe/route"
	"messageServe/server"
	_ "net/http/pprof"
)

func main() {
	linker.Main()
	server.Main()
	route.Main()
}
