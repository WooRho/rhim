package rooter

import (
	"github.com/gin-gonic/gin"
	h "rhim/internal/handle"
)

func UserRoot(r *gin.RouterGroup) {
	userGroup := r.Group("user")
	userGroup.GET("/getUserList", h.GetUserList)
	userGroup.POST("/createUser", h.CreateUser)
	userGroup.PUT("/deleteUser", h.DeleteUser)
	userGroup.PUT("/updateUser", h.UpdateUser)
}
