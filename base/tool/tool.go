package tool

import (
	"short_url/base/config"
	"sync"
)

var (
	one sync.Once
	err error
)

func Init() {
	one.Do(func() {
		initLogger(getLoggerOptions())
		//NewWorker ID 填入 分布式的服务唯一ID从1到1024
		if err = NewWorker(1); err != nil {
			panic(err)
		}
	})
}
func getLoggerOptions() *Options {
	op := &Options{}
	op.Development = config.GetToolLogConfig().GetDevelopment()
	op.LogFileDir = config.GetToolLogConfig().GetLogFileDir()
	op.AppName = config.GetToolLogConfig().GetAppName()
	op.MaxSize = config.GetToolLogConfig().GetMaxSize()
	op.MaxBackups = config.GetToolLogConfig().GetMaxBackups()
	op.MaxAge = config.GetToolLogConfig().GetMaxAge()
	return op
}
