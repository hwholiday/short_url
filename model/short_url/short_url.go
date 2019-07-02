package short_url

import (
	"fmt"
	"short_url/base/db"
	"short_url/model/sequence"
	"sync"

	"github.com/go-redis/redis"

	"github.com/go-xorm/xorm"
)

var (
	s *service
	m sync.Once
)

type service struct {
	m        *xorm.Engine
	r        *redis.Client
	sequence sequence.Service
}

type Service interface {
	CreateLinks(url, key string, status int) (string, error)
	GetLinksByUrl(url string) (*Links, error)
	GetLinksByKeyword(keyword string) (*Links, error)
}

func Init() {
	var err error
	m.Do(func() {
		s = &service{}
		s.m = db.GetMySqlDb()
		s.r = db.GetRedisDb()
		s.sequence, err = sequence.GetServer()
		if err != nil {
			panic(err)
		}
	})
}

func GetServer() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[short_url] no init")
	}
	return s, nil
}
