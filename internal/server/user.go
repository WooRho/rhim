package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rhim/config"
	"rhim/middleware"
)

func TestRootFunc(ctx *gin.Context) {
	db := middleware.GetDb(config.GetMysql())
	db.Begin()
	defer db.Commit()

	ctx.String(http.StatusOK, "hello World!")
}
