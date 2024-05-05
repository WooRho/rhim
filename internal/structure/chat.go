package structure

type (
	ChatParam struct {
		UserId   uint   `json:"user_id" form:"user_id"`
		Type     string `json:"type" form:"type"`
		TargetId int    `json:"target_id" form:"target_id"`
		Context  string `json:"context" form:"context"`
	}
)
