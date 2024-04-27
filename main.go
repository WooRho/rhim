package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rhim/config"
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
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	err = r.Run(":" + systemInfo.Port)
	if err != nil {
		return
	}
}
