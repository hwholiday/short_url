package base

import (
	"seqsvr/base/config"
	"seqsvr/base/db"
	"seqsvr/base/tool"
)

func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}
