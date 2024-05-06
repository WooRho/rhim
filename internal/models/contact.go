package models

import (
	"context"
	"gorm.io/gorm"
	"rhim/internal/structure"
	"rhim/middleware"
	"rhim/tools"
)

// 人员关系
type (
	Contact struct {
		Basic
		OwnerId  uint   `gorm:"column:owner_id;type:bigint(20);not null;default:0;comment:谁的关系信息" ` //谁的关系信息
		TargetId uint   `gorm:"column:target_id;type:bigint(20);not null;default:0;comment:对应的谁" `  //对应的谁
		Type     int    `gorm:"column:type;type:int(11);not null;default:0;comment:对应的类型" `         //对应的类型  0  1  3
		Desc     string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述" `      // 描述
	}
	ContactList []*Contact
)

func (table *Contact) TableName() string {
	return "contact"
}

func (m *Contact) New(req *structure.AddContactInfo) {
	m.ID = middleware.Snowflake.GenerateID().UInt()
	m.OwnerId = req.OwnerId
	m.TargetId = req.TargetId
	m.Type = req.Type
	m.Desc = req.Desc
}

func (m *Contact) Update(req *structure.UpdateContactInfo) {
	m.OwnerId = req.OwnerId
	m.TargetId = req.TargetId
	m.Type = req.Type
	m.Desc = req.Desc
}

func (m *Contact) BuildResp() *structure.ContactInfo {
	resp := &structure.ContactInfo{}
	resp.OwnerId = m.OwnerId
	resp.TargetId = m.TargetId
	resp.Type = m.Type
	resp.Desc = m.Desc
	resp.Id = m.ID
	resp.CreatedAt = tools.Time2String(m.CreatedAt, tools.YMDHMS)
	resp.UpdatedAt = tools.Time2String(m.UpdatedAt, tools.YMDHMS)
	resp.CreatorId = m.CreatorId
	resp.UpdaterId = m.UpdaterId
	return resp
}

type (
	ContactDao struct {
		db *gorm.DB
	}
	ContactDaoInterface interface {
		Search(ctx context.Context, req *structure.SearchContactInfo) (
			list ContactList, count int64, err error)
	}
)

func NewContactDao(db *gorm.DB) ContactDaoInterface {
	return &ContactDao{
		db: db,
	}
}

func (d *ContactDao) Search(ctx context.Context, req *structure.SearchContactInfo) (
	list ContactList, count int64, err error) {
	var (
		_list = make(ContactList, 0)
	)
	db := d.db
	if req.OwnerId > 0 {
		db = db.Where("owner_id = ?", req.OwnerId)
	}
	if req.Type > 0 {
		db = db.Where("type = ?", req.Type)
	}
	if req.IsPage() {
		db = db.Limit(req.Limit).Offset(req.Offset)
	}
	err = db.Count(&count).Error
	if err != nil {
		return
	}
	err = db.Find(&_list).Error
	if err != nil {
		return
	}
	list = _list
	return
}
