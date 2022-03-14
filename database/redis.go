package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func StartRedis(ctx context.Context) error {
	pong, err := Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("連線redis出錯，錯誤資訊：%v", err)
	} else {
		fmt.Println("成功連線redis", pong)
	}
	return err
}
