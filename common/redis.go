package common

import (
	"github.com/astaxie/beego"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func RedisInit() (err error) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis::host") + ":" + beego.AppConfig.String("redis::port"),
		Password: beego.AppConfig.String("redis::auth"),
		DB:       0, // use default DB
	})

	return
}
