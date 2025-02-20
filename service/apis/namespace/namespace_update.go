package namespace

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UpdateNamespaceApi struct {
	*common.ApiCommon
	Data UpdateNamespaceRequest
}
type UpdateNamespaceRequest struct {
	model.Namespace
}

func UpdateNamespace(c *gin.Context) {
	req := &UpdateNamespaceApi{
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

func (req *UpdateNamespaceApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
	var count int64
	err := mysql.DB.Model(&model.Namespace{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		Count(&count).Error
	if err != nil {
		req.Logger.Error("check record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	if count <= 0 {
		req.Reply.MsgSet(common.ReplyStatusBindRequestFailed, common.ReplyMessageBindRequestFailed)
		return
	}
}

func (req *UpdateNamespaceApi) updateRecord() {
	now := time.Now().Unix()
	record := &model.Namespace{
		Namespace: req.Data.Namespace.Namespace,
		Parent:    req.Data.Parent,
		Name:      req.Data.Name,
		Describe:  req.Data.Describe,
		Operator:  req.Data.Operator,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.Namespace{}).
		Where("id = ?", req.Data.Id).
		Updates(&record).Error
	if err != nil {
		req.Logger.Error("update record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.UpdateFailed()
		return
	}
}
