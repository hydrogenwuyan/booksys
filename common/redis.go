package common

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit() (err error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis::host") + ":" + beego.AppConfig.String("redis::port"),
		Password: beego.AppConfig.String("redis::auth"),
		DB:       0, // use default DB
	})

	return
}
