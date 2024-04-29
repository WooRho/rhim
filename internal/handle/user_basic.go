package handle

import (
	"context"
	"github.com/gin-gonic/gin"
	"rhim/internal/logic"
	"rhim/internal/structure"
	"rhim/middleware"
	"rhim/tools"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	var (
		err       error
		db        = middleware.GetDb()
		p         = &structure.SearchUserBasicInfo{}
		data      = &structure.UserBasicInfoList{}
		total     int64
		ctx       = context.TODO()
		userLogic = logic.NewUserBasicLogic(db)
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	defer func() {
		tools.BuildListResponse(c, err, data, total)
	}()

	data, total, err = userLogic.GetUserList(ctx, p)
	if err != nil {
		return
	}

}

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {

	var (
		err       error
		db        = middleware.GetDb()
		p         = &structure.AddUserBasicInfo{}
		data      = &structure.Id{}
		ctx       = context.TODO()
		userLogic = logic.NewUserBasicLogic(db)
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	defer func() {
		tools.BuildResponse(c, err, data)
	}()
	data, err = userLogic.CreateUser(ctx, p)
	return
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {

	var (
		err       error
		db        = middleware.GetDb()
		p         = &structure.Id{}
		data      = &structure.Id{}
		ctx       = context.TODO()
		userLogic = logic.NewUserBasicLogic(db)
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	defer func() {
		tools.BuildResponse(c, err, data)
	}()
	err = userLogic.DeleteUser(ctx, p)
	return
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	var (
		err       error
		db        = middleware.GetDb()
		p         = &structure.UpdateUserBasicInfo{}
		data      = &structure.Id{}
		ctx       = context.TODO()
		userLogic = logic.NewUserBasicLogic(db)
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	defer func() {
		tools.BuildResponse(c, err, data)
	}()
	data, err = userLogic.UpdateUser(ctx, p)
	return

}
