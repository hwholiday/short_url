package db

import (
	"sync"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	err     error
	mysqlDb *xorm.Engine
	redisDb *redis.Client
	m       sync.Mutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	initMysql()
	initRedis()
}

func GetMySqlDb() *xorm.Engine {
	return mysqlDb
}

func CloseMySqlDb() {
	closeMysql()
}

func GetRedisDb() *redis.Client {
	return redisDb
}

func CloseRedisDb() {
	closeRedis()
}
