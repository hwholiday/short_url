package model

import (
	"short_url/model/sequence"
	"short_url/model/short_url"
)

func Init() {
	sequence.Init()
	short_url.Init()
}
