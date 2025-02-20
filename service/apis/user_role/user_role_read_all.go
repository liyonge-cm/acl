package user_role

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserRoleListApi struct {
	*common.ApiCommon
	Data GetUserRoleListRequest
}
type GetUserRoleListRequest struct {
	common.Page
	Namespace string `json:"namespace" binding:"required"` // 所属项目组
}
type GetUserRoleListResponse struct {
	Data  []*model.UserRole `json:"data"`
	Count int64             `json:"count"`
}

func GetUserRoleList(c *gin.Context) {
	req := &GetUserRoleListApi{
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

	res := GetUserRoleListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetUserRoleListApi) getRecords() (records []*model.UserRole, count int64) {
	records = make([]*model.UserRole, 0)
	tx := mysql.DB.Model(&model.UserRole{}).Where("status != ?", common.RecordStatusDeleted)
	if req.Data.Namespace != "" {
		tx = tx.Where("namespace = ?", req.Data.Namespace)
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
