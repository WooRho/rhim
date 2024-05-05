package models

import (
	"gorm.io/gorm"
	"rhim/internal/structure"
	"rhim/middleware"
	"rhim/tools"
)

type (
	Message struct {
		Basic
		FromId   uint   `gorm:"column:from_id;type:bigint(20) unsigned;not null;default:0;comment:发送者"`   // 发送者
		TargetId uint   `gorm:"column:target_id;type:bigint(20) unsigned;not null;default:0;comment:接受者"` // 接受者
		Type     string `gorm:"column:type;type:varchar(255);not null;default:'私聊';comment:消息类型"`         // 聊天类型  群聊 私聊 广播
		Media    int    `gorm:"column:media;type:tinyint(2) unsigned;not null;default:1;comment:消息类型"`    // 消息类型  文字 图片 音频
		Content  string `gorm:"column:content;type:text;not null;comment:消息内容"`                           //消息内容
		Pic      string `gorm:"column:pic;type:text ;not null;comment:图片"`                                // 图片
		Url      string `gorm:"column:url;type:text ;not null;comment:url"`                               // url
		Desc     string `gorm:"column:desc;type:text ;not null;comment:描述"`                               // 描述
		Amount   int    `gorm:"column:amount;type:int(11) ;not null;default:0;comment:其他数字统计"`            //其他数字统计
	}
)

func (table *Message) TableName() string {
	return "message"
}

func (m *Message) New(req *structure.AddMessageInfo) {
	m.ID = middleware.Snowflake.GenerateID().UInt()
	m.FromId = req.FromId
	m.TargetId = req.TargetId
	m.Type = req.Type
	m.Media = req.Media
	m.Content = req.Content
	m.Pic = req.Pic
	m.Url = req.Url
	m.Desc = req.Desc
	m.Amount = req.Amount
}

func (m *Message) Update(req *structure.UpdateMessageInfo) {
	m.FromId = req.FromId
	m.TargetId = req.TargetId
	m.Type = req.Type
	m.Media = req.Media
	m.Content = req.Content
	m.Pic = req.Pic
	m.Url = req.Url
	m.Desc = req.Desc
	m.Amount = req.Amount
}

func (m *Message) BuildResp() *structure.MessageInfo {
	resp := &structure.MessageInfo{}
	resp.FromId = m.FromId
	resp.TargetId = m.TargetId
	resp.Type = m.Type
	resp.Media = m.Media
	resp.Content = m.Content
	resp.Pic = m.Pic
	resp.Url = m.Url
	resp.Desc = m.Desc
	resp.Amount = m.Amount
	resp.Id = m.ID
	resp.CreatedAt = tools.Time2String(m.CreatedAt, tools.YMDHMS)
	resp.UpdatedAt = tools.Time2String(m.UpdatedAt, tools.YMDHMS)
	resp.CreatorId = m.CreatorId
	resp.UpdaterId = m.UpdaterId
	return resp
}

type (
	MessageDao struct {
		db *gorm.DB
	}
	MessageDaoInterface interface {
	}
)

func NewMessageDao(db gorm.DB) MessageDaoInterface {
	return &MessageDao{
		db: &db,
	}
}
