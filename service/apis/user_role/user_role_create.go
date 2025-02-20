package user_role

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateUserRoleApi struct {
	*common.ApiCommon
	Data *model.UserRole
}

func CreateUserRole(c *gin.Context) {
	req := &CreateUserRoleApi{
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

func (req *CreateUserRoleApi) checkParams() {
	user := model.User{}
	err := mysql.DB.Model(&model.User{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("user = ?", req.Data.User).
		First(&user).Error
	if err != nil {
		req.Logger.Error("get role failed", zap.Error(err))
		req.Reply.CheckParamFailed("user")
		return
	}
	role := model.Role{}
	err = mysql.DB.Model(&model.Role{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("namespace = ?", req.Data.Namespace).
		Where("id = ?", req.Data.RoleId).
		First(&role).Error
	if err != nil {
		req.Logger.Error("get role failed", zap.Error(err))
		req.Reply.CheckParamFailed("role_id")
		return
	}
}

func (req *CreateUserRoleApi) createRecord() {
	now := time.Now().Unix()
	record := &model.UserRole{
		Namespace: req.Data.Namespace,
		User:      req.Data.User,
		RoleId:    req.Data.RoleId,
		Operator:  req.Data.Operator,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.UserRole{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
