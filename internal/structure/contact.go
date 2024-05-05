package structure

type (
	contact struct {
		OwnerId  uint   `json:"owner_id"`  //谁的关系信息
		TargetId uint   `json:"target_id"` //对应的谁
		Type     int    `json:"type"`      //对应的类型  0  1  3
		Desc     string `json:"desc"`      // 描述
	}
	ContactInfo struct {
		BasicRecord
		contact
	}
	AddContactInfo struct {
		contact
	}
	UpdateContactInfo struct {
		contact
		Id uint `json:"id"`
	}
	SearchContactInfo struct {
		ListQuery
		contact
	}
	ContactInfoList []*ContactInfo
)
