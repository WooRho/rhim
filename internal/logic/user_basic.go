package logic

import (
	"context"
	"gorm.io/gorm"
	"rhim/internal/models"
	"rhim/internal/structure"
)

type (
	UserBasicLogic struct {
		db  *gorm.DB
		sql models.UserBasicDaoInterface
	}
	UserBasicLogicInterface interface {
		GetUserList(ctx context.Context, req *structure.SearchUserBasicInfo) (data *structure.UserBasicInfoList, total int64, err error)
		Get(ctx context.Context, req *structure.Id) (data *structure.UserBasicInfo, err error)
		CreateUser(ctx context.Context, req *structure.AddUserBasicInfo) (data *structure.Id, err error)
		UpdateUser(ctx context.Context, req *structure.UpdateUserBasicInfo) (data *structure.Id, err error)
		DeleteUser(ctx context.Context, req *structure.Id) (err error)
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
func (l *UserBasicLogic) CreateUser(ctx context.Context, req *structure.AddUserBasicInfo) (data *structure.Id, err error) {
	var (
		model = &models.UserBasic{}
		resp  = &structure.Id{}
	)
	model.New(req)
	err = l.sql.Create(ctx, model)
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
	err = l.sql.Update(ctx, model)
	resp.Id = model.ID
	data = resp
	return
}

func (l *UserBasicLogic) DeleteUser(ctx context.Context, req *structure.Id) (err error) {
	return l.sql.Delete(ctx, req.Id)
}
