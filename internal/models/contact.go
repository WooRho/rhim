package models

// 人员关系
type Contact struct {
	Basic
	OwnerId  uint   `gorm:"column:owner_id;type:bigint(20);not null;default:0;comment:谁的关系信息" ` //谁的关系信息
	TargetId uint   `gorm:"column:target_id;type:bigint(20);not null;default:0;comment:对应的谁" `  //对应的谁
	Type     int    `gorm:"column:type;type:int(11);not null;default:0;comment:对应的类型" `         //对应的类型  0  1  3
	Desc     string `gorm:"column:desc;type:varchar(255);not null;default:'';comment:描述" `      // 描述
}

func (table *Contact) TableName() string {
	return "contact"
}
