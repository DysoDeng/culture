package redis

import (
	"culture/internal/config"
	"fmt"
	"github.com/go-redis/redis/v7"
)

var redisPoolClient *redis.Client

// 初始化redis连接
func init() {
	addr := config.Config.Redis.Host + ":" + config.Config.Redis.Port
	redisPoolClient = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.DataBase,
		MinIdleConns: 2,
	})

	pong, err := redisPoolClient.Ping().Result()
	fmt.Println(pong, err)
}

// 获取redis实例
func Client() *redis.Client {
	return redisPoolClient
}

// 构建安全key
func Key(key string) string {
	prefix := config.Config.Redis.KeyPrefix
	if prefix != "" {
		key = config.Config.Redis.KeyPrefix + ":" + key
	}
	return key
}
