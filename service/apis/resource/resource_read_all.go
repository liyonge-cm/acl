package resource

import (
	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GetResourceListApi struct {
	*common.ApiCommon
	Data GetResourceListRequest
}
type GetResourceListRequest struct {
	common.Page
	Namespace string `json:"namespace" binding:"required"` // 所属项目组
	Category  string `json:"category"`                     // 资源分类
	Resource  string `json:"resource"`                     // 资源标识
	Name      string `json:"name"`                         // 资源中文名
}
type GetResourceListResponse struct {
	Data  []*model.Resource `json:"data"`
	Count int64             `json:"count"`
}

func GetResourceList(c *gin.Context) {
	req := &GetResourceListApi{
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

	res := GetResourceListResponse{
		Data:  records,
		Count: count,
	}
	req.Reply.DataSet(res)
	req.Reply.Response(c)
}

func (req *GetResourceListApi) getRecords() (records []*model.Resource, count int64) {
	records = make([]*model.Resource, 0)
	tx := mysql.DB.Model(&model.Resource{}).Where("status != ?", common.RecordStatusDeleted)
	if req.Data.Namespace != "" {
		tx = tx.Where("namespace = ?", req.Data.Namespace)
	}
	if req.Data.Category != "" {
		tx = tx.Where("category = ?", req.Data.Category)
	}
	if req.Data.Resource != "" {
		tx = tx.Where("resource = ?", req.Data.Resource)
	}
	if req.Data.Name != "" {
		tx = tx.Where("name = ?", req.Data.Name)
	}

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
