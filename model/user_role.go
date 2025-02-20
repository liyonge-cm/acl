package model

type UserRole struct {
	Id        uint   `json:"id"`         //
	Namespace string `json:"namespace"`  // 所属项目组
	User      string `json:"user"`       // 用户标识
	RoleId    uint   `json:"role_id"`    // 角色ID
	Operator  string `json:"operator"`   // 创建人
	Status    int    `json:"status"`     // 状态: 0:正常，-1:删除
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

func (UserRole) TableName() string {
	return "user_role"
}
