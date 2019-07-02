package model

import (
	"seqsvr/model/sequence"
	"seqsvr/model/short_url"
)

func Init() {
	sequence.Init()
	short_url.Init()
}
