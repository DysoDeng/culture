package redis

import (
	"culture/internal/config"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/goava/di"
	"sync"
)

// container Redis服务容器
var container *di.Container

var mutex sync.Mutex

func init() {
	if config.Config.Env != config.Release {
		di.SetTracer(&di.StdTracer{})
	}
	var err error
	container, err = di.New(
		di.Provide(initRedis),
	)
	if err != nil {
		panic(err)
	}
}

// 初始化redis连接
func initRedis() *redis.Client {
	mutex.Lock()
	defer func() {
		mutex.Unlock()
	}()
	addr := config.Config.Redis.Host + ":" + config.Config.Redis.Port
	redisPoolClient := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     config.Config.Redis.Password,
		DB:           config.Config.Redis.DataBase,
		MinIdleConns: 2,
	})

	pong, err := redisPoolClient.Ping().Result()
	fmt.Println(pong, err)
	return redisPoolClient
}

// 获取redis实例
func Client() *redis.Client {
	var client *redis.Client
	_ = container.Resolve(&client)
	return client
}

// 构建安全key
func Key(key string) string {
	prefix := config.Config.Redis.KeyPrefix
	if prefix != "" {
		key = config.Config.Redis.KeyPrefix + ":" + key
	}
	return key
}
