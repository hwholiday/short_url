package main

import (
	"flag"
	"seqsvr/base"
	"seqsvr/controller"
	"seqsvr/server"
)

var path = flag.String("p", "./conf/", "")
var addr = flag.String("s", "127.0.0.1:8099", "server addr")

func main() {
	flag.Parse()
	base.Init(*path)
	controller.Init()
	server.Run(*addr)
}
