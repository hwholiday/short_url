package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"short_url/base/config"
	"time"
)

func initMongoDb() {
	//链接mongo服务
	opt := options.Client().ApplyURI(config.GetMgoConfig().GetUrl())
	opt.SetLocalThreshold(3 * time.Second) //只使用与mongo操作耗时小于3秒的
	opt.SetAuth(options.Credential{
		Username: config.GetMgoConfig().GetName(),
		Password: config.GetMgoConfig().GetPass(),
	})
	opt.SetMaxConnIdleTime(5 * time.Second) //指定连接可以保持空闲的最大毫秒数
	opt.SetMaxPoolSize(200)                 //使用最大的连接数
	if mgo, err = mongo.Connect(context.TODO(), opt); err != nil {
		panic(err)
	}
}
