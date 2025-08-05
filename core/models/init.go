package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"xorm.io/xorm"
)

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8")
	if err != nil {
		slog.Error(fmt.Sprint("XormNewEngineErr ", err))
		return nil
	}
	return engine
}
