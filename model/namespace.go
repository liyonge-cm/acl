package model

type Namespace struct {
	Id        uint   `json:"id"`         //
	Namespace string `json:"namespace"`  // 项目组标识
	Parent    string `json:"parent"`     // 所属父项目组
	Name      string `json:"name"`       // 项目组中文名
	Describe  string `json:"describe"`   // 描述
	Operator  string `json:"operator"`   // 创建人
	Status    int    `json:"status"`     // 状态: 0:正常，-1:删除
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

func (Namespace) TableName() string {
	return "namespace"
}
