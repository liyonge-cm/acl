package role_permission

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateRolePermissionApi struct {
	*common.ApiCommon
	Data UpdateRolePermissionRequest
}
type UpdateRolePermissionRequest struct {
	model.RolePermission
}

func UpdateRolePermission(c *gin.Context) {
	req := &UpdateRolePermissionApi{
		ApiCommon: common.NewRequest(c),
	}
	if err := req.BindRequest(&req.Data); err != nil {
		req.Reply.BindRequestFailed().Response(c)
		return
	}
	req.checkParams()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}
	req.updateRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *UpdateRolePermissionApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}

	permission := model.RolePermission{}
	err := mysql.DB.Model(&model.RolePermission{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		First(&permission).Error
	if err != nil {
		req.Logger.Error("get role permission failed", zap.Error(err))
		req.Reply.CheckParamFailed("Id")
		return
	}
	resource := model.Resource{}
	err = mysql.DB.Model(&model.Resource{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("namespace = ?", permission.Namespace).
		Where("resource_id = ?", req.Data.ResourceId).
		First(&resource).Error
	if err != nil {
		req.Logger.Error("get resource failed", zap.Error(err))
		req.Reply.CheckParamFailed("ResourceId")
		return
	}
}

func (req *UpdateRolePermissionApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.RolePermission{
		ResourceId: req.Data.ResourceId,
		Describe:   req.Data.Describe,
		Operator:   req.Data.Operator,
		UpdatedAt:  now,
	}
	err := mysql.DB.Model(&model.RolePermission{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
}
