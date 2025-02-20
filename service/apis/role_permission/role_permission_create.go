package role_permission

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateRolePermissionApi struct {
	*common.ApiCommon
	Data *model.RolePermission
}

func CreateRolePermission(c *gin.Context) {
	req := &CreateRolePermissionApi{
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
	req.createRecord()
	if req.Reply.IsStatusFailed() {
		req.Reply.Response(c)
		return
	}

	req.Reply.Response(c)
}

func (req *CreateRolePermissionApi) checkParams() {
	role := model.Role{}
	err := mysql.DB.Model(&model.Role{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("namespace = ?", req.Data.Namespace).
		Where("id = ?", req.Data.RoleId).
		First(&role).Error
	if err != nil {
		req.Logger.Error("get role failed", zap.Error(err))
		req.Reply.CheckParamFailed("role_id")
		return
	}
	resource := model.Resource{}
	err = mysql.DB.Model(&model.Resource{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("namespace = ?", req.Data.Namespace).
		Where("id = ?", req.Data.ResourceId).
		First(&resource).Error
	if err != nil {
		req.Logger.Error("get resource failed", zap.Error(err))
		req.Reply.CheckParamFailed("resource_id")
		return
	}
}

func (req *CreateRolePermissionApi) createRecord() {
	now := time.Now().Unix()
	record := &model.RolePermission{
		Namespace:  req.Data.Namespace,
		RoleId:     req.Data.RoleId,
		ResourceId: req.Data.ResourceId,
		Describe:   req.Data.Describe,
		Operator:   req.Data.Operator,
		Status:     common.RecordStatusInit,
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	err := mysql.DB.Model(&model.RolePermission{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
