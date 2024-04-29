package structure

type (
	BasicRecord struct {
		Id          uint64 `json:"id"`           // id
		CreatedAt   string `json:"created_at"`   // 创建时间
		UpdatedAt   string `json:"updated_at"`   // 更新时间
		CreatorId   uint64 `json:"creator_id"`   // 创建人id
		UpdaterId   uint64 `json:"updater_id"`   // 更新人id
		CreatorName string `json:"creator_name"` // 创建人名称
		UpdaterName string `json:"updater_name"` // 更新人名称
	}
	BasicRecordUnScope struct {
		Id          uint64 `json:"id"`           // id
		CreatedAt   string `json:"created_at"`   // 创建时间
		UpdatedAt   string `json:"updated_at"`   // 更新时间
		DeletedAt   string `json:"deleted_at"`   // 删除时间
		CreatorId   uint64 `json:"creator_id"`   // 创建人id
		UpdaterId   uint64 `json:"updater_id"`   // 更新人id
		DeletedId   uint64 `json:"deleted_id"`   // 删除人id
		CreatorName string `json:"creator_name"` // 创建人名称
		UpdaterName string `json:"updater_name"` // 更新人名称
		DeletedName string `json:"deleted_name"` // 删除人名称
	}
)

type IParam interface {
	AdjustParam()
	ValidateParam() error
}

type Param struct {
}

func (Param) ValidateParam() error {
	return nil
}

func (s Param) AdjustParam() {

}

type IResponseData interface {
	AdjustData()
}

type ResponseData struct {
}

func (ResponseData) AdjustData() {

}

// 通用响应结构体
type Response struct {
	baseResponse
	Data interface{} `json:"data,omitempty"`
}

// 列表响应结构体
type ListResponse struct {
	baseResponse
	Data struct {
		Summary interface{} `json:"summary,omitempty"`
		List    interface{} `json:"list,omitempty"`
		Total   int         `json:"total"`
	} `json:"data"`
}

type baseResponse struct {
	Code    int    `json:"code"`    // 响应码
	Msg     string `json:"msg"`     // 响应信息
	Version string `json:"version"` // 版本号
}

func (r *baseResponse) SetVersion(version string) {
	r.Version = version
}
