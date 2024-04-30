package models

type (
	Message struct {
		Basic
		FromId   uint   `gorm:"column:from_id;type:bigint(20) unsigned;not null;default:0;comment:发送者"`   // 发送者
		TargetId uint   `gorm:"column:target_id;type:bigint(20) unsigned;not null;default:0;comment:接受者"` // 接受者
		Type     string `gorm:"column:type;type:bigint(20) unsigned;not null;default:私聊;comment:消息类型"`    // 聊天类型  群聊 私聊 广播
		Media    int    `gorm:"column:media;type:bigint(20) unsigned;not null;default:1;comment:消息类型"`    // 消息类型  文字 图片 音频
		Content  string `gorm:"column:content;type:text unsigned;not null;default:'';comment:消息内容"`       //消息内容
		Pic      string `gorm:"column:pic;type:text unsigned;not null;default:'';comment:图片"`             // 图片
		Url      string `gorm:"column:url;type:text unsigned;not null;default:'';comment:url"`            // url
		Desc     string `gorm:"column:desc;type:text unsigned;not null;default:'';comment:描述"`            // 描述
		Amount   int    `gorm:"column:amount;type:int(11) unsigned;not null;default:0;comment:其他数字统计"`    //其他数字统计
	}
)

func (table *Message) TableName() string {
	return "message"
}
