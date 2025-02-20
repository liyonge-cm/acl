package role

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetRoleListApi struct {
	*common.ApiCommon
	Data GetRoleListRequest
}
type GetRoleListRequest struct {
	common.Page
	Namespace string `json:"namespace" binding:"required"` // 所属项目组
	Role      string `json:"role"`                         // 角色标识
	Name      string `json:"name"`                         // 角色中文名
}
type GetRoleListResponse struct {
	Data  []*model.Role `json:"data"`
	Count int64         `json:"count"`
}

func GetRoleList(c *gin.Context) {
	req := &GetRoleListApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		req.Reply.Response(c)
		return
	}

	records, count := req.getRecords()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	res := GetRoleListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetRoleListApi) getRecords() (records []*model.Role, count int64) {
	records = make([]*model.Role, 0)
	tx := mysql.DB.Model(&model.Role{}).Where("status != ?", common.RecordStatusDeleted)
	if req.Data.Namespace != "" {
		tx = tx.Where("namespace = ?", req.Data.Namespace)
	}
	if req.Data.Role != "" {
		tx = tx.Where("role = ?", req.Data.Role)
	}
	if req.Data.Name != "" {
		tx = tx.Where("name = ?", req.Data.Name)
	}

	tx = tx.Count(&count)
	if count == 0 {
		return
	}

	if req.Data.Limit > 0 && req.Data.Offset > 0 {
		tx = tx.Limit(req.Data.Limit).Offset(req.Data.Limit * (req.Data.Offset - 1))
	}
	err := tx.Order("created_at desc").Find(&records).Error
	if err != nil {
		req.Logger.Error("get records failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}

	return records, count
}
