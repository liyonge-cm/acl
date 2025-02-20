package role

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateRoleApi struct {
	*common.ApiCommon
	Data *model.Role
}

func CreateRole(c *gin.Context) {
	req := &CreateRoleApi{
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

func (req *CreateRoleApi) checkParams() {
}

func (req *CreateRoleApi) createRecord() {
	now := time.Now().Unix()
	record := &model.Role{
		Namespace:	req.Data.Namespace,
		Role:	req.Data.Role,
		Name:	req.Data.Name,
		Describe:	req.Data.Describe,
		Operator:	req.Data.Operator,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Role{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
