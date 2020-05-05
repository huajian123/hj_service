package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

var engine *xorm.Engine

func GetDb() *xorm.Engine {
	return engine
}

func InitDb() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@/hjdb?charset=utf8")
	if err != nil {
		println(err.Error())
		return
	}
	engine.ShowSQL(true)
}
