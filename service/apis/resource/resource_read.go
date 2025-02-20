package resource

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetResourceApi struct {
	*common.ApiCommon
	Data GetResourceRequest
}
type GetResourceRequest struct {
	Id	int	`json:"id"`
}
type GetResourceResponse struct {
	Data	*model.Resource	`json:"data"`
}

func GetResource(c *gin.Context) {
	req := &GetResourceApi{
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
	res := GetResourceResponse{
		Data:  record,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetResourceApi) checkParams() {
	if req.Data.Id <= 0 {
		req.Reply.MsgSet(common.ReplyStatusMissingParam, common.ReplyMessageMissingParam)
		return
	}
}

func (req *GetResourceApi) getRecord() (record *model.Resource) {
	err := mysql.DB.Model(&model.Resource{}).
		Where("status != ?", common.RecordStatusDeleted).
		Where("id = ?", req.Data.Id).
		First(&record).Error
	if err != nil {
		req.Logger.Error("get record failed", zap.Error(err))
		req.Reply.ReadFailed()
		return
	}
	return record
}
