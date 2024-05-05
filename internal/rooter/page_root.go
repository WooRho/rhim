package rooter

import (
	"github.com/gin-gonic/gin"
	h "rhim/internal/handle"
)

func PageRoot(r *gin.RouterGroup) {
	r.GET("/", h.GetIndex)
	r.GET("index", h.GetIndex)

}
