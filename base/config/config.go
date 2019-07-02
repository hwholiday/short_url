package config

import (
	"path/filepath"
	"sync"
)
import "github.com/go-ini/ini"

var (
	mysqlConfig defaultMysqlConfig
	utilsConfig defaultLogToolConfig
	redisConfig defaultRedisConfig
	m           sync.Mutex
)

func Init(path string) {
	var (
		err error
		cfg *ini.File
	)

	m.Lock()
	defer m.Unlock()
	if cfg, err = ini.Load(filepath.Join(path, "db.ini")); err != nil {
		panic(err)
	}
	if err = cfg.Section("mysql").MapTo(&mysqlConfig); err != nil {
		panic(err)
	}
	if err = cfg.Section("redis").MapTo(&redisConfig); err != nil {
		panic(err)
	}
	if cfg, err = ini.Load(filepath.Join(path, "tool.ini")); err != nil {
		panic(err)
	}
	if err = cfg.Section("zap").MapTo(&utilsConfig); err != nil {
		panic(err)
	}
}

func GetMysqlConfig() (fig sqlConfig) {
	return mysqlConfig
}
func GetToolLogConfig() (fig toolLogConfig) {
	return utilsConfig
}

func GetRedisConfig() (fig rdsConfig) {
	return redisConfig
}
