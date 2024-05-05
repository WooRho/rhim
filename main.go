package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"rhim/config"
	"rhim/internal/models"
	"rhim/internal/rooter"
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

	err = r.Run(":" + systemInfo.Port)
	if err != nil {
		return
	}
}

func Init(config config.Config, r *gin.Engine) {
	middleware.NewDatabase(&config.Mysql)
	middleware.NewServiceContext()
	middleware.InitSnowflake()
	models.InitSql()
	InitRoot(config, r)
}

func InitRoot(c config.Config, r *gin.Engine) {
	baseGroup := r.Group("")
	rooter.UserRoot(baseGroup)
	rooter.SwagRoot(baseGroup)
	rooter.PageRoot(baseGroup)

	// 静态资源
	r.Static("front/asset", "front/asset")
	r.LoadHTMLGlob("front/view/**/*")
}
