package redis

import "github.com/mikeqiao/newant/db/redis"

var R redis.CRedis

func InitRedis() {
	R.InitDB()
}

func OnClose() {
	R.OnClose()
}
