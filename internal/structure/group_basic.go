package structure

type (
	groupBasic struct {
		Name    string `json:"name"`     // 名字
		OwnerId uint   `json:"owner_id"` // 群主
		Icon    string `json:"icon"`     // 图标
		Type    int    `json:"type"`     // 群类型
		Desc    string `json:"desc"`     // 描述
	}
	GroupBasicInfo struct {
		BasicRecord
		groupBasic
	}
	AddGroupBasicInfo struct {
		groupBasic
	}
	UpdateGroupBasicInfo struct {
		groupBasic
		Id uint `json:"id"`
	}
	SearchGroupBasicInfo struct {
		ListQuery
		groupBasic
	}
	GroupBasicInfoList []*GroupBasicInfo
)
