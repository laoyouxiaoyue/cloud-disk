package models

import (
	"cloud-disk/core/internal/config"
	"fmt"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"xorm.io/xorm"
)

func Init(datasource string) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", datasource)
	if err != nil {
		slog.Error(fmt.Sprint("XormNewEngineErr ", err))
		return nil
	}
	return engine
}

func InitRedis(c config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Redis.Addr,
		Password: "",
		DB:       0,
	})
}
