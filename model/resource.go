package model

type Resource struct {
	Id         uint   `json:"id"`         //
	Namespace  string `json:"namespace"`  // 所属项目组
	Category   string `json:"category"`   // 资源分类
	Resource   string `json:"resource"`   // 资源标识
	Properties string `json:"properties"` // 资源属性 (JSON Schema)
	Name       string `json:"name"`       // 资源中文名
	Describe   string `json:"describe"`   // 描述
	Operator   string `json:"operator"`   // 创建人
	Status     int    `json:"status"`     // 状态: 0:正常，-1:删除
	CreatedAt  int64  `json:"created_at"` // 创建时间
	UpdatedAt  int64  `json:"updated_at"` // 更新时间
}

func (Resource) TableName() string {
	return "resource"
}
