package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// 启动redis客户端
func NewClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // DB selection
	})

	// 测试连接
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return rdb
}

// 存储短码和长URL的映射关系
func ShortCodeimageUrl(rdb *redis.Client, shortCode string, longURL string) error {
	err := rdb.Set(context.Background(), shortCode, longURL, 120*time.Second).Err()
	return err
}
