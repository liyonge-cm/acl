package user

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetUserPermissionApi struct {
	*common.ApiCommon
	Data GetUserPermissionRequest
}
type GetUserPermissionRequest struct {
	User      string `json:"user" binding:"required"`
	Namespace string `json:"namespace" binding:"required"` // 所属项目组
}
type GetUserPermissionResponse struct {
	Data []*model.Resource `json:"data"`
}

func GetUserPermission(c *gin.Context) {
	req := &GetUserPermissionApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed()
		req.Reply.Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	record := req.getRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	res := GetUserPermissionResponse{
		Data: record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetUserPermissionApi) checkParams() {
	if req.Data.User == "" {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetUserPermissionApi) getRecord() (record []*model.Resource) {
	err := mysql.DB.Table("user").
		Joins("left join user_role ON user_role.user = user.user").
		Joins("left join role_permission ON user_role.role_id = role_permission.role_id").
		Joins("left join resource ON resource.id = role_permission.resource_id").
		Where("user.status != ?", common.RecordStatusDeleted).
		Where("user_role.status != ?", common.RecordStatusDeleted).
		Where("role_permission.status != ?", common.RecordStatusDeleted).
		Where("resource.status != ?", common.RecordStatusDeleted).
		Where("user.user = ?", req.Data.User).
		Where("resource.namespace = ?", req.Data.Namespace).
		Select("resource.*").
		Find(&record).Error
	if err != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}
