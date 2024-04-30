package models

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"rhim/internal/structure"
	"rhim/middleware"
	"rhim/tools"
	"time"
)

type (
	UserBasic struct {
		Basic
		Name          string    `gorm:"column:name;type:varchar(255) ;not null;default:'';comment:名字" `                            // 名字
		Password      string    `gorm:"column:pass_word;type:varchar(255) ;not null;default:'';comment:密码"`                        // 密码
		Phone         string    `gorm:"column:phone; type:varchar(255) ;not null;default:'';comment:手机号"`                          // 手机号
		Email         string    `gorm:"column:email; type:varchar(255) ;not null;default:'';comment:邮箱"`                           // 邮箱
		Identity      string    `gorm:"column:identity; type:varchar(255) ;not null;default:'';comment:身份"`                        // 身份
		Salt          string    `gorm:"column:salt; type:varchar(255) ;not null;default:'';comment:加盐"`                            // 加盐
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

func (m *UserBasic) New(req *structure.AddUserBasicInfo) {
	m.ID = middleware.Snowflake.GenerateID().UInt()
	m.Name = req.Name
	m.Salt = tools.GenerateSalt()
	m.Password = tools.MakePassword(req.Password, "salt")
	m.Phone = req.Phone
	m.Email = req.Email
	m.Identity = req.Identity
	m.ClientIp = req.ClientIp
	m.ClientPort = req.ClientPort
	m.DeviceInfo = req.DeviceInfo
}

func (m *UserBasic) Update(req *structure.UpdateUserBasicInfo) {
	m.Name = req.Name
	m.Password = req.Password
	m.Phone = req.Phone
	m.Email = req.Email
	m.Identity = req.Identity
	m.ClientIp = req.ClientIp
	m.ClientPort = req.ClientPort
	m.DeviceInfo = req.DeviceInfo
}

func (m *UserBasic) BuildResp() *structure.UserBasicInfo {
	resp := &structure.UserBasicInfo{}
	resp.Name = m.Name
	resp.Password = m.Password
	resp.Phone = m.Phone
	resp.Email = m.Email
	resp.Identity = m.Identity
	resp.ClientIp = m.ClientIp
	resp.ClientPort = m.ClientPort
	resp.DeviceInfo = m.DeviceInfo
	resp.Id = m.ID
	resp.CreatedAt = tools.Time2String(m.CreatedAt, tools.YMDHMS)
	resp.UpdatedAt = tools.Time2String(m.UpdatedAt, tools.YMDHMS)
	resp.CreatorId = m.CreatorId
	resp.UpdaterId = m.UpdaterId
	return resp
}

func (m *UserBasic) GetToken() {
	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	m.Identity = tools.MD5Encode(str)
}

type (
	UserBasicDao struct {
		db *gorm.DB
	}
	UserBasicDaoInterface interface {
		Create(ctx context.Context, user *UserBasic) error
		Update(ctx context.Context, user *UserBasic) error
		Delete(ctx context.Context, id uint) error
		Get(ctx context.Context, id uint) (*UserBasic, error)
		GetList(ctx context.Context, req *structure.SearchUserBasicInfo) (UserBasicList, int64, error)
		Unique(ctx context.Context, req *UserBasic) error
		FindUserByNameAndPwd(ctx context.Context, name string, password string) (*UserBasic, error)
		FindUserByName(ctx context.Context, name string) (*UserBasic, error)
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

func (d *UserBasicDao) Create(ctx context.Context, user *UserBasic) error {
	return d.db.Create(user).Error
}

func (d *UserBasicDao) Update(ctx context.Context, user *UserBasic) error {
	modifyMap := tools.Structure2ModifyMap(user)
	return d.db.Updates(modifyMap).Error
}

func (d *UserBasicDao) Delete(ctx context.Context, id uint) error {
	return d.db.Where("id = ?", id).Delete(&UserBasic{}).Error
}
func (d *UserBasicDao) Get(ctx context.Context, id uint) (*UserBasic, error) {
	var (
		m = &UserBasic{}
	)
	err := d.db.Where("id = ?", id).First(m).Error
	return m, err
}
func (d *UserBasicDao) GetList(ctx context.Context, req *structure.SearchUserBasicInfo) (UserBasicList, int64, error) {
	var (
		count int64
		model = UserBasicList{}
	)

	db := d.db
	if req.IsPage() {
		db = db.Limit(req.Limit).Offset(req.Offset)
	}
	db.Order("id desc").Count(&count).Find(&model)
	if db.Error != nil {
		panic(db.Error.Error())
	}
	return model, count, nil
}

func (d *UserBasicDao) Unique(ctx context.Context, req *UserBasic) error {
	d.db.Where("name = ? and id != ?", req.Name, req.ID).First(&req)
	if req.ID > 0 {
		return errors.New("名称已被使用")
	}
	d.db.Where("phone = ? and id != ?", req.Phone, req.ID).First(&req)
	if req.ID > 0 {
		return errors.New("号码已被使用")
	}
	d.db.Where("email = ? and id != ?", req.Email, req.ID).First(&req)
	if req.ID > 0 {
		return errors.New("号码已被使用")
	}
	return nil
}

func (d *UserBasicDao) FindUserByNameAndPwd(ctx context.Context, name string, password string) (*UserBasic, error) {
	user := &UserBasic{}
	err := d.db.Where("name = ? and pass_word=?", name, password).First(user).Error
	return user, err
}

func (d *UserBasicDao) FindUserByName(ctx context.Context, name string) (*UserBasic, error) {
	user := &UserBasic{}
	err := d.db.Where("name = ?", name).First(user).Error
	return user, err
}
