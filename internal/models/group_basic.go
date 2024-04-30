package models

import "gorm.io/gorm"

type GroupBasic struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名字" `     // 名字
	OwnerId uint   `gorm:"column:owner_id;type:varchar(255);not null;default:'';comment:群主" ` // 群主
	Icon    string `gorm:"column:icon;type:text;not null;default:'';comment:图标" `             // 图标
	Type    int    `gorm:"column:type;type:int(11);not null;default:0;comment:群类型" `          // 群类型
	Desc    string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述" `     // 描述
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}
