package db

import (
	"fmt"
	"seqsvr/base/config"
	"seqsvr/base/tool"

	"github.com/go-xorm/xorm"
)

func initMysql() {
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8",
		config.GetMysqlConfig().GetName(), config.GetMysqlConfig().GetPass(), config.GetMysqlConfig().GetIP(), config.GetMysqlConfig().GetPort(), config.GetMysqlConfig().GetDb())
	if mysqlDb, err = xorm.NewEngine("mysql", url); err != nil {
		panic(err)
	}
	mysqlDb.SetMaxOpenConns(config.GetMysqlConfig().GetMaxIdle())
	mysqlDb.SetMaxIdleConns(config.GetMysqlConfig().GetMaxOpen())
	_ = mysqlDb.Ping()
	if config.GetToolLogConfig().GetDevelopment() {
		mysqlDb.ShowSQL(true)
		mysqlDb.ShowExecTime(true)
	}
	tool.GetLogger().Debug("mysql : " + url)

	//tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "im_")
	//mysqlDb.SetTableMapper(tbMapper)
}

func closeMysql() {
	if mysqlDb != nil {
		_ = mysqlDb.Close()
	}
}
