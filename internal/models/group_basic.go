package models

import (
	"gorm.io/gorm"
	"rhim/internal/structure"
	"rhim/middleware"
	"rhim/tools"
)

type GroupBasic struct {
	Basic
	Name    string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名字" `    // 名字
	OwnerId uint   `gorm:"column:owner_id;type:varchar(255);not null;default:0;comment:群主" ` // 群主
	Icon    string `gorm:"column:icon;type:text;not null;default:'';comment:图标" `            // 图标
	Type    int    `gorm:"column:type;type:int(11);not null;default:0;comment:群类型" `         // 群类型
	Desc    string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述" `    // 描述
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}

func (m *GroupBasic) New(req *structure.AddGroupBasicInfo) {
	m.ID = middleware.Snowflake.GenerateID().UInt()
	m.Name = req.Name
	m.OwnerId = req.OwnerId
	m.Icon = req.Icon
	m.Type = req.Type
	m.Desc = req.Desc
}

func (m *GroupBasic) Update(req *structure.UpdateGroupBasicInfo) {
	m.Name = req.Name
	m.OwnerId = req.OwnerId
	m.Icon = req.Icon
	m.Type = req.Type
	m.Desc = req.Desc
}

func (m *GroupBasic) BuildResp() *structure.GroupBasicInfo {
	resp := &structure.GroupBasicInfo{}
	resp.Name = m.Name
	resp.OwnerId = m.OwnerId
	resp.Icon = m.Icon
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
	GroupBasicDao struct {
		db *gorm.DB
	}
	GroupBasicDaoInterface interface {
	}
)

func NewGroupBasicDao(db gorm.DB) GroupBasicDaoInterface {
	return &GroupBasicDao{
		db: &db,
	}
}
