package logic

import (
	"context"
	"gorm.io/gorm"
	"rhim/internal/models"
	"rhim/internal/structure"
)

type (
	ContactLogic struct {
		db  *gorm.DB
		sql models.ContactDaoInterface
	}
	ContactLogicInterface interface {
		GetContactList(ctx context.Context, req structure.SearchContactInfo) (
			data *structure.ContactInfoList, total int64, err error)
		//Get(ctx context.Context, req *structure.Id) (data *structure.ContactInfo, err error)
		//CreateContact(ctx context.Context, req *structure.AddContactInfo) (data *structure.Id, err error)
		//UpdateContact(ctx context.Context, req *structure.UpdateContactInfo) (data *structure.Id, err error)
		//DeleteContact(ctx context.Context, req *structure.Id) (err error)
	}
)

func NewContactLogic(db *gorm.DB) *ContactLogic {
	return &ContactLogic{
		db:  db,
		sql: models.NewContactDao(db),
	}
}

func (l *ContactLogic) GetContactList(ctx context.Context, req *structure.SearchContactInfo) (
	data structure.ContactInfoList, total int64, err error) {
	var (
		list  = make(models.ContactList, 0)
		_data = make(structure.ContactInfoList, 0)
	)
	list, total, err = l.sql.Search(ctx, req)
	if err != nil {
		return
	}
	for _, contact := range list {
		resp := contact.BuildResp()
		_data = append(_data, resp)
	}
	data = _data
	return
}
