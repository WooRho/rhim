package logic

import (
	"context"
	"gorm.io/gorm"
	"rhim/internal/structure"
)

type (
	UserBasicLogic struct {
		db *gorm.DB
	}
	UserBasicLogicInterface interface {
		GetUserList(ctx context.Context, req *structure.SearchUserBasicInfo) (data *structure.UserBasicInfoList, total int, err error)
		Get(ctx context.Context, req *structure.Id) (data *structure.UserBasicInfoList, err error)
		CreateUser(ctx context.Context, req *structure.AddUserBasicInfo) (data *structure.Id, err error)
		UpdateUser(ctx context.Context, req *structure.UpdateUserBasicInfo) (data *structure.Id, err error)
		DeleteUser(ctx context.Context, req *structure.Id) (err error)
	}
)

func NewUserBasicLogic(db *gorm.DB) UserBasicLogicInterface {
	return &UserBasicLogic{
		db: db,
	}
}
