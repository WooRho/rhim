package server

import "github.com/gin-gonic/gin"

func UserRoot(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	userGroup.GET("", TestRootFunc)
	userGroup.PUT("")
	userGroup.POST("")
	userGroup.DELETE("")
}
