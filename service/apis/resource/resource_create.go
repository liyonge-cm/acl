package resource

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateResourceApi struct {
	*common.ApiCommon
	Data *model.Resource
}

func CreateResource(c *gin.Context) {
	req := &CreateResourceApi{
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

func (req *CreateResourceApi) checkParams() {
}

func (req *CreateResourceApi) createRecord() {
	now := time.Now().Unix()
	record := &model.Resource{
		Namespace:	req.Data.Namespace,
		Category:	req.Data.Category,
		Resource:	req.Data.Resource,
		Properties:	req.Data.Properties,
		Name:	req.Data.Name,
		Describe:	req.Data.Describe,
		Operator:	req.Data.Operator,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Resource{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
