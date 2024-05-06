package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"rhim/internal/models"
	"rhim/internal/structure"
	"rhim/tools"
)

type (
	UserBasicLogic struct {
		db  *gorm.DB
		sql models.UserBasicDaoInterface
	}
	UserBasicLogicInterface interface {
		GetUserList(ctx context.Context, req *structure.SearchUserBasicInfo) (data *structure.UserBasicInfoList, total int64, err error)
		Get(ctx context.Context, req *structure.Id) (data *structure.UserBasicInfo, err error)
		GetByIds(ctx context.Context, req *structure.Id) (data structure.UserBasicInfoList, err error)
		CreateUser(ctx context.Context, req *structure.AddUserBasicInfo) (data *structure.Id, err error)
		UpdateUser(ctx context.Context, req *structure.UpdateUserBasicInfo) (data *structure.Id, err error)
		DeleteUser(ctx context.Context, req *structure.Id) (err error)
		FindUserByNameAndPwd(ctx context.Context, req *structure.SearchUserBasicInfo) (resp *structure.UserBasicInfo, err error)
	}
)

func NewUserBasicLogic(db *gorm.DB) UserBasicLogicInterface {
	return &UserBasicLogic{
		db:  db,
		sql: models.NewUserBasicDao(db),
	}
}

func (l *UserBasicLogic) GetUserList(ctx context.Context, req *structure.SearchUserBasicInfo) (data *structure.UserBasicInfoList, total int64, err error) {
	var (
		list []*models.UserBasic
		resp = make(structure.UserBasicInfoList, 0)
	)
	list, total, err = l.sql.GetList(ctx, req)
	for _, basic := range list {
		info := &structure.UserBasicInfo{}
		info = basic.BuildResp()
		resp = append(resp, info)
	}
	data = &resp
	return
}

func (l *UserBasicLogic) Get(ctx context.Context, req *structure.Id) (data *structure.UserBasicInfo, err error) {
	var (
		model = &models.UserBasic{}
	)
	model, err = l.sql.Get(ctx, req.Id)
	if err != nil {
		return
	}
	data = model.BuildResp()
	return
}

func (l *UserBasicLogic) GetByIds(ctx context.Context, req *structure.Id) (data structure.UserBasicInfoList, err error) {
	var (
		list  = models.UserBasicList{}
		_data = make(structure.UserBasicInfoList, 0)
	)
	list, err = l.sql.GetByIds(ctx, req.IdsSlice)
	if err != nil {
		return
	}
	for _, user := range list {
		resp := user.BuildResp()
		_data = append(_data, resp)
	}
	data = _data
	return
}

func (l *UserBasicLogic) CreateUser(ctx context.Context, req *structure.AddUserBasicInfo) (data *structure.Id, err error) {
	var (
		model = &models.UserBasic{}
		resp  = &structure.Id{}
	)
	model.New(req)
	err = l.sql.Unique(ctx, model)
	if err != nil {
		return
	}
	err = l.sql.Create(ctx, model)
	if err != nil {
		return
	}
	resp.Id = model.ID
	data = resp
	return
}

func (l *UserBasicLogic) UpdateUser(ctx context.Context, req *structure.UpdateUserBasicInfo) (data *structure.Id, err error) {
	var (
		model = &models.UserBasic{}
		resp  = &structure.Id{}
	)
	model, err = l.sql.Get(ctx, req.Id)
	if err != nil {
		return
	}
	model.Update(req)
	err = l.sql.Unique(ctx, model)
	if err != nil {
		return
	}
	err = l.sql.Update(ctx, model)
	resp.Id = model.ID
	data = resp
	return
}

func (l *UserBasicLogic) DeleteUser(ctx context.Context, req *structure.Id) (err error) {
	return l.sql.Delete(ctx, req.Id)
}

func (l *UserBasicLogic) FindUserByNameAndPwd(ctx context.Context, req *structure.SearchUserBasicInfo) (resp *structure.UserBasicInfo, err error) {
	var (
		user = &models.UserBasic{}
	)
	user, err = l.sql.FindUserByName(ctx, req.Name)
	if user.Name == "" {
		return nil, errors.New("该用户不存在")
	}
	flag := tools.ValidPassword(req.Password, user.Salt, user.Password)
	if !flag {
		return nil, errors.New("密码不正确")
	}
	pwd := tools.MakePassword(req.Password, user.Salt)
	user, err = l.sql.FindUserByNameAndPwd(ctx, req.Name, pwd)
	if err != nil {
		return
	}
	user.GetToken()
	err = l.sql.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	resp = user.BuildResp()
	return resp, err
}
