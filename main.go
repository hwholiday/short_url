package main

import (
	"flag"
	"short_url/base"
	"short_url/controller"
	"short_url/server"
)

var path = flag.String("p", "./conf/", "")
var addr = flag.String("s", "127.0.0.1:8099", "server addr")

func main() {
	flag.Parse()
	base.Init(*path)
	controller.Init()
	server.Run(*addr)
}
