package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"rhim/config"
)

type ServiceContext struct {
	DB    *gorm.DB
	Redis *redis.Client
	ctx   context.Context
}

func NewServiceContext(config config.Config) *ServiceContext {
	return &ServiceContext{
		DB:    GetDb(config.Mysql),
		Redis: InitRedis(config.Redis),
		ctx:   context.Background(),
	}
}

func InitRedis(config config.Redis) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port), // Redis地址
		Password: "",                                             // Redis密码，如果没有则为空字符串
		DB:       0,                                              // 使用默认DB
	})
	return rdb
}
