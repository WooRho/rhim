package structure

type (
	Message struct {
		FromId   uint   `json:"from_id"`   // 发送者
		TargetId uint   `json:"target_id"` // 接受者
		Type     string `json:"type"`      // 聊天类型  群聊 私聊 广播
		Media    int    `json:"media"`     // 消息类型  文字 图片 音频
		Content  string `json:"content"`   //消息内容
		Pic      string `json:"pic"`       // 图片
		Url      string `json:"url"`       // url
		Desc     string `json:"desc"`      // 描述
		Amount   int    `json:"amount"`    //其他数字统计
	}
)
