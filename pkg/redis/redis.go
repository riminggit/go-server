package Redis

import (
	"context"
	"fmt"
	"go-server-template/config"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var rdb *redis.Client
var ctx = context.Background()

//redis连接
func InitRedis() *redis.Client {
	config := projectConfig.AppConfig
	rdb = redis.NewClient(&redis.Options{
		Addr:         config.RedisConfig.HOST + ":" + config.RedisConfig.PORT,
		Password:     config.RedisConfig.PASSWORD,
		DialTimeout:  time.Duration(config.RedisConfig.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(config.RedisConfig.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(config.RedisConfig.WriteTimeout) * time.Second,
		PoolSize:     config.RedisConfig.PoolSize,
		PoolTimeout:  time.Duration(config.RedisConfig.PoolTimeout) * time.Second,
		DB:           0, // use default DB
	})

	// ping一下检查是否连通
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Error(err)
	}
	// PONG
	fmt.Println("redis 连接成功")

	return rdb
}

func SetValue(key string, value interface{}, expiration time.Duration) error {
	// 0 意味着没有过期时间
	err := rdb.Set(ctx, key, value, expiration*time.Second).Err()
	if err != nil {
		log.Error(err)
	}
	return err
}

func GetValue(key string) interface{} {
	val, err := rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		log.Error("key对应的值不存在")
	} else if err != nil {
		log.Error(err)
	}

	return val
}

func DelValue(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		log.Error(err)
	}
	return err
}
