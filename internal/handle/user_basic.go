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
//
//	@Summary	所有用户
//	@Tags		用户模块
//
//	@Produce	json
//	@Param		page	query		int		false	"页码"
//	@Param		size	query		int		false	"页大小"
//	@Param		code	query		string	false	"名字"
//	@Param		code	query		string	false	"密码"
//	@Param		code	query		string	false	"手机号"
//	@Param		code	query		string	false	"邮箱"
//	@Param		code	query		string	false	"加盐"
//	@Param		code	query		string	false	"身份"
//	@Param		code	query		string	false	"客户端ip"
//	@Param		code	query		string	false	"客户端端口"
//	@Param		code	query		string	false	"登录时间"
//	@Param		code	query		string	false	"心跳时间"
//	@Param		code	query		string	false	"登出时间"
//	@Param		code	query		string	false	"是否登出"
//	@Param		code	query		string	false	"设备信息"
//
//	@Success	200		{object}	structure.UserBasicInfoList{}
//	@Router		/user/getUserList [get]
func GetUserList(c *gin.Context) {
	var (
		err   error
		db    = middleware.GetDb()
		p     = &structure.SearchUserBasicInfo{}
		data  = &structure.UserBasicInfoList{}
		total int64
		ctx   = context.TODO()
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	defer func() {
		tools.BuildListResponse(c, err, data, total)
	}()

	data, total, err = logic.NewUserBasicLogic(db).GetUserList(ctx, p)
	if err != nil {
		return
	}

}

// CreateUser
//
//	@Summary	新增用户
//	@Tags		用户模块
//
//	@Param		body	body		structure.AddUserBasicInfo{}	true	"创建"
//
//	@Success	200		{object}	structure.Id{}
//	@Router		/user/createUser [get]
func CreateUser(c *gin.Context) {

	var (
		err  error
		db   = middleware.GetDb()
		p    = &structure.AddUserBasicInfo{}
		data = &structure.Id{}
		ctx  = context.TODO()
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
		tools.BuildResponse(c, err, data)
	}()
	data, err = logic.NewUserBasicLogic(db).CreateUser(ctx, p)
	return
}

// DeleteUser
//
//	@Summary	删除用户
//	@Tags		用户模块
//	@Param		body	body		structure.Id{}	true	"刪除"
//	@Success	200		{object}	structure.Id{}
//	@Router		/user/deleteUser [get]
func DeleteUser(c *gin.Context) {

	var (
		err  error
		db   = middleware.GetDb()
		p    = &structure.Id{}
		data = &structure.Id{}
		ctx  = context.TODO()
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
		tools.BuildResponse(c, err, data)
	}()
	err = logic.NewUserBasicLogic(db).DeleteUser(ctx, p)
	return
}

// UpdateUser
//
//	@Summary	修改用户
//	@Tags		用户模块
//	@Param		body	body		structure.UpdateUserBasicInfo{}	true	"更新"
//	@Success	200		{object}	structure.Id{}
//	@Router		/user/updateUser [post]
func UpdateUser(c *gin.Context) {
	var (
		err  error
		db   = middleware.GetDb()
		p    = &structure.UpdateUserBasicInfo{}
		data = &structure.Id{}
		ctx  = context.TODO()
	)

	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
		tools.BuildResponse(c, err, data)
	}()
	data, err = logic.NewUserBasicLogic(db).UpdateUser(ctx, p)
	return

}

// GetUserList
// @Summary 登录
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [get]
func Login(c *gin.Context) {

	var (
		err  error
		db   = middleware.GetDb()
		p    = &structure.SearchUserBasicInfo{}
		data = &structure.UserBasicInfo{}
		ctx  = context.TODO()
	)
	err = tools.ShouldBind(c, p)
	if err != nil {
		return
	}
	db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
		tools.BuildResponse(c, err, data)
	}()

	data, err = logic.NewUserBasicLogic(db).FindUserByNameAndPwd(ctx, p)
	return
}
