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
	}
)

func NewUserBasicLogic(db *gorm.DB) UserBasicLogicInterface {
	return &UserBasicLogic{
		db: db,
	}
}

func (l *UserBasicLogic) GetUserList(ctx context.Context, req *structure.UserBasicInfo) {

}
