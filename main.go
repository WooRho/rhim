package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"rhim/config"
	"rhim/internal/server"
	"rhim/middleware"
)

func main() {
	// init
	var (
		initConfig, err = config.InitConfig()
		systemInfo      = initConfig.System
	)
	if err != nil {
		err = errors.New("yaml init error")
		return
	}

	// 1.创建路由
	r := gin.Default()
	err = r.SetTrustedProxies([]string{systemInfo.Host})
	if err != nil {
		return
	}
	Init(initConfig, r)
	// gin.Context，封装了request和response

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err = r.Run(":" + systemInfo.Port)
	if err != nil {
		return
	}
}

func Init(config config.Config, r *gin.Engine) {
	middleware.NewDatabase(config.Mysql)
	middleware.NewServiceContext(config)
	InitRoot(config, r)
}

func InitRoot(c config.Config, r *gin.Engine) {
	baseGroup := r.Group(c.System.Name)
	server.UserRoot(baseGroup)
}
