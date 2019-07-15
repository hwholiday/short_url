package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"sync"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	err     error
	mysqlDb *xorm.Engine
	redisDb *redis.Client
	mgo     *mongo.Client
	m       sync.Mutex
)

func Init() {
	m.Lock()
	defer m.Unlock()
	initMysql()
	initRedis()
	initMongoDb()
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
