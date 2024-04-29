package structure

type (
	UserBasicInfo struct {
		BasicRecord
		Name          string `json:"name"`                                     // 名字
		PassWord      string `json:"pass_word"`                                // 密码
		Phone         string `json:"phone" valid:"matches(^1[3-9]{1}\\d{9}$)"` // 手机号
		Email         string `json:"email" valid:"email"`                      // 邮箱
		Identity      string `json:"identity"`                                 // 身份
		ClientIp      string `json:"client_ip"`                                // 客户端ip
		ClientPort    string `json:"client_port"`                              // 客户端端口
		LoginTime     string `json:"login_time"`                               // 登录时间
		HeartbeatTime string `json:"heartbeat_time"`                           // 心跳时间
		LoginOutTime  string `json:"login_out_time"`                           // 登出时间
		IsLogout      uint8  `json:"is_logout"`                                // 是否登出
		DeviceInfo    string `json:"device_info"`                              // 设备信息
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
