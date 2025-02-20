package common

import (
	"acl/service/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type ApiCommon struct {
	c      *gin.Context    `json:"-"`
	Logger *zap.Logger     `json:"-"`
	Reply  *CommonResponse `json:"-"`
}

type Page struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

func NewRequest(c *gin.Context) *ApiCommon {
	return &ApiCommon{
		c:      c,
		Logger: logger.Logger,
		Reply:  NewResponse(),
	}
}

func (r *ApiCommon) BindRequest(req interface{}) (err error) {
	// err = r.c.BindJSON(&req)
	err = r.c.ShouldBindBodyWith(req, binding.JSON)
	if err != nil {
		r.Logger.Error("读取API入参失败", zap.Error(err))
		return err
	}

	return
}
