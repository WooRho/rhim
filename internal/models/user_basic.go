package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	UserBasic struct {
		Basic
		Name          string    `gorm:"column:name;type:varchar(255) ;not null;default:'';comment:名字" `                            // 名字
		PassWord      string    `gorm:"column:pass_word;type:varchar(255) ;not null;default:'';comment:密码"`                        // 密码
		Phone         string    `gorm:"column:phone; type:varchar(255) ;not null;default:'';comment:手机号"`                          // 手机号
		Email         string    `gorm:"column:email; type:varchar(255) ;not null;default:'';comment:邮箱"`                           // 邮箱
		Identity      string    `gorm:"column:identity; type:varchar(255) ;not null;default:'';comment:身份"`                        // 身份
		ClientIp      string    `gorm:"column:client_ip; type:varchar(255) ;not null;default:'';comment:客户端ip"`                    // 客户端ip
		ClientPort    string    `gorm:"column:client_port; type:varchar(255) ;not null;default:'';comment:客户端端口"`                  // 客户端端口
		LoginTime     time.Time `gorm:"column:login_time; type:datetime; not null;default:'0000-00-00 00:00:00';comment:登录时间"`     // 登录时间
		HeartbeatTime time.Time `gorm:"column:heartbeat_time; type:datetime; not null;default:'0000-00-00 00:00:00';comment:心跳时间"` // 心跳时间
		LoginOutTime  time.Time `gorm:"column:login_out_time; type:datetime; not null;default:'0000-00-00 00:00:00';comment:登出时间"` // 登出时间
		IsLogout      bool      `gorm:"column:is_logout;type:tinyint(1) unsigned;not null;default:0;comment:是否登出"`                 // 是否登出
		DeviceInfo    string    `gorm:"column:device_info;type:varchar(255);not null;default:'';comment:设备信息"`                     // 设备信息
	}
	UserBasicList []*UserBasic
)

func (table *UserBasic) TableName() string {
	return "user_basic"
}

type (
	UserBasicDao struct {
		db *gorm.DB
	}
	UserBasicDaoInterface interface {
		TestCreate() error
	}
)

func NewUserBasicDao(db *gorm.DB) UserBasicDaoInterface {
	err := db.AutoMigrate(&UserBasic{})
	if err != nil {
		fmt.Println(err)
	}
	return &UserBasicDao{
		db: db,
	}
}

func (d *UserBasicDao) TestCreate() error {

	user := &UserBasic{}
	user.Name = "申专"
	return d.db.Create(user).Error
}
