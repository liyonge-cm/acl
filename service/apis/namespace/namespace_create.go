package namespace

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateNamespaceApi struct {
	*common.ApiCommon
	Data *model.Namespace
}

func CreateNamespace(c *gin.Context) {
	req := &CreateNamespaceApi{
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

func (req *CreateNamespaceApi) checkParams() {
}

func (req *CreateNamespaceApi) createRecord() {
	now := time.Now().Unix()
	record := &model.Namespace{
		Namespace: req.Data.Namespace,
		Parent:    req.Data.Parent,
		Name:      req.Data.Name,
		Describe:  req.Data.Describe,
		Operator:  req.Data.Operator,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Namespace{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
