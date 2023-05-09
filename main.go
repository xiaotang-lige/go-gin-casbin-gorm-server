package main

import (
	"messageServe/linker"
	"messageServe/server"
	_ "net/http/pprof"
)

func main() {

	linker.Main()
	server.Main()
}
