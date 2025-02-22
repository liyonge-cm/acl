package user

import (
	"time"

	"acl/model"
	"acl/service/apis/common"
	"acl/service/mysql"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CreateUserApi struct {
	*common.ApiCommon
	Data *model.User
}

func CreateUser(c *gin.Context) {
	req := &CreateUserApi{
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

func (req *CreateUserApi) checkParams() {
}

func (req *CreateUserApi) createRecord() {
	now := time.Now().Unix()
	record := &model.User{
		OId:	req.Data.OId,
		User:	req.Data.User,
		Name:	req.Data.Name,
		Job:	req.Data.Job,
		PhoneNumber:	req.Data.PhoneNumber,
		Email:	req.Data.Email,
		Status:    common.RecordStatusInit,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := mysql.DB.Model(&model.User{}).Create(&record).Error
	if err != nil {
		req.Logger.Error("create record failed", zap.Any("Data", record), zap.Error(err))
		req.Reply.CreateFailed()
		return
	}
}
