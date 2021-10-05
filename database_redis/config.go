package database_redis

import (
	"gopkg.in/redis.v5"
	"os"
)

func Setup() (*redis.Client,error) {
	client:=redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: os.Getenv("REDISPASS"),
		DB: 0,
	})

	_,err:=client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client,err
}
