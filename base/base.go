package base

import (
	"short_url/base/config"
	"short_url/base/db"
	"short_url/base/tool"
)

func Init(path string) {
	config.Init(path)
	tool.Init()
	db.Init()
}
