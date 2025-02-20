package namespace

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetNamespaceListApi struct {
	*common.ApiCommon
	Data GetNamespaceListRequest
}
type GetNamespaceListRequest struct {
	common.Page
}
type GetNamespaceListResponse struct {
	Data	[]*model.Namespace	`json:"data"`
	Count int64	`json:"count"`
}

func GetNamespaceList(c *gin.Context) {
	req := &GetNamespaceListApi{
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

	res := GetNamespaceListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetNamespaceListApi) getRecords() (records []*model.Namespace, count int64) {
	records = make([]*model.Namespace, 0)
	tx := mysql.DB.Model(&model.Namespace{}).Where("status != ?", common.RecordStatusDeleted)

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
