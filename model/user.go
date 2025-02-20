package model

type User struct {
	Id          uint   `json:"id"`           // 自有id, 应该是不会使用的
	OId         int    `json:"o_id"`         // oa那边的Id
	User        string `json:"user"`         // 用户英文名
	Name        string `json:"name"`         // 用户中文名
	Job         string `json:"job"`          // 职位名称
	PhoneNumber string `json:"phone_number"` // 手机号码
	Email       string `json:"email"`        // 邮箱
	Status      int    `json:"status"`       // 状态: 0:正常，-1:删除
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

func (User) TableName() string {
	return "user"
}
