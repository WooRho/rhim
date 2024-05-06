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
	userGroup.POST("/login", h.Login)
	r.POST("/searchFriends", h.SearchFriends)
	//发送消息
	r.GET("/user/sendMsg", h.SendMsg)
	r.GET("/user/senUserMsg", h.SenUserMsg)
	r.GET("/toChat", h.ToChat)
}
