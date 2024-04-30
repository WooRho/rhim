package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"rhim/config"
)

var SvcCtx = &ServiceContext{}

type ServiceContext struct {
	DB    *gorm.DB
	Redis *redis.Client
	ctx   context.Context
}

func NewServiceContext() {
	SvcCtx = &ServiceContext{
		DB:    GetDb(config.GetMysql()),
		Redis: InitRedis(config.GetRedis()),
		ctx:   context.Background(),
	}
	return
}

func InitRedis(redisConfig *config.Redis) *redis.Client {
	if redisConfig == nil {
		redisConfig = config.GetRedis()
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port), // Redis地址
		Password: "",                                                       // Redis密码，如果没有则为空字符串
		DB:       0,                                                        // 使用默认DB
	})
	pong, err := rdb.Ping(context.TODO()).Result()
	if err != nil {
		fmt.Println("init redis  。。。。", err)
	} else {
		fmt.Println(" Redis inited 。。。。", pong)
	}
	return rdb
}

// 全局变量
var (
	Snowflake NodeIface
)

func InitSnowflake() {
	Snowflake = NewCustomNode()
}
