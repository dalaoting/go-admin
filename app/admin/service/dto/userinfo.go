package dto

import (
	"github.com/gin-gonic/gin"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/pkg/models"
	"go-admin/tools"
	"gorm.io/gorm"
)

type UserInfoSearch struct {
	dto.Pagination `search:"-"`
}

func (m *UserInfoSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *UserInfoSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *UserInfoSearch) Generate() dto.Index {
	o := *m
	return &o
}

type UserInfoControl struct {
	ID uint `uri:"ID" comment:"主键id"` // 主键id

	Status string `json:"status" comment:"状态"`
}

func (s *UserInfoControl) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBindUri(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
		return err
	}
	err = ctx.ShouldBind(s)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
	}
	return err
}

func (s *UserInfoControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserInfoControl) GenerateM() (common.ActiveRecord, error) {
	return &models.UserInfo{

		Model:  gorm.Model{ID: s.ID},
		Status: s.Status,
	}, nil
}

func (s *UserInfoControl) GetId() interface{} {
	return s.ID
}

type UserInfoById struct {
	dto.ObjectById
}

func (s *UserInfoById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *UserInfoById) GenerateM() (common.ActiveRecord, error) {
	return &models.UserInfo{}, nil
}
