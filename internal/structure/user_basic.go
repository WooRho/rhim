package structure

type (
	common struct {
		Name          string `json:"name" form:"name"`                                      // 名字
		PassWord      string `json:"pass_word" form:"pass_word"`                            // 密码
		Phone         string `json:"phone" form:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"` // 手机号
		Email         string `json:"email" form:"email"  valid:"email"`                     // 邮箱
		Identity      string `json:"identity" form:"identity"`                              // 身份
		ClientIp      string `json:"client_ip" form:"client_ip"`                            // 客户端ip
		ClientPort    string `json:"client_port" form:"client_port"`                        // 客户端端口
		LoginTime     string `json:"login_time" form:"login_time"`                          // 登录时间
		HeartbeatTime string `json:"heartbeat_time" form:"heartbeat_time"`                  // 心跳时间
		LoginOutTime  string `json:"login_out_time" form:"login_out_time"`                  // 登出时间
		IsLogout      uint8  `json:"is_logout" form:"is_logout"`                            // 是否登出
		DeviceInfo    string `json:"device_info" form:"device_info"`                        // 设备信息
	}
	UserBasicInfo struct {
		BasicRecord
	}
	AddUserBasicInfo struct {
		common
	}
	UpdateUserBasicInfo struct {
		common
		Id uint64 `json:"id"`
	}
	SearchUserBasicInfo struct {
		common
	}
	UserBasicInfoList []*UserBasicInfo
)

func (u UserBasicInfoList) AdjustData() {
}

func (u UserBasicInfo) AdjustParam() {

}

func (u UserBasicInfo) ValidateParam() error {
	return nil
}
