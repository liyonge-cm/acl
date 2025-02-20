package model

type RolePermission struct {
	Id         uint   `json:"id"`          //
	Namespace  string `json:"namespace"`   // 所属项目组
	RoleId     uint   `json:"role_id"`     // 角色ID
	ResourceId uint   `json:"resource_id"` // 资源ID
	Describe   string `json:"describe"`    // 描述
	Operator   string `json:"operator"`    // 创建人
	Status     int    `json:"status"`      // 状态: 0:正常，-1:删除
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
}

func (RolePermission) TableName() string {
	return "role_permission"
}
