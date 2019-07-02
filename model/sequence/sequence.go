package sequence

import (
	"fmt"
	"seqsvr/base/db"
	"sync"

	"github.com/go-redis/redis"

	"github.com/go-xorm/xorm"
)

var (
	s *service
	m sync.Once
)

type service struct {
	m *xorm.Engine
	r *redis.Client
}

type Service interface {
	GetBorrowOrder() (int64, error)
}

func Init() {
	m.Do(func() {
		s = &service{}
		s.m = db.GetMySqlDb()
		s.r = db.GetRedisDb()
	})
}

func GetServer() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[sequence] no init")
	}
	return s, nil
}
