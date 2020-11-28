package dto

import (
	"github.com/gin-gonic/gin"
	"time"

	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/pkg/models"
	"go-admin/tools"
)

type WithdrawAccountSearch struct {
	dto.Pagination `search:"-"`
	Identity       string `form:"identity" search:"type:exact;column:identity;table:withdraw_account" comment:"用户id"`

	Type string `form:"type" search:"type:exact;column:type;table:withdraw_account" comment:"类型"`

	Account string `form:"account" search:"type:exact;column:account;table:withdraw_account" comment:"账号"`

	Status string `form:"status" search:"type:exact;column:status;table:withdraw_account" comment:"状态"`

	IsDelete string `form:"isDelete" search:"type:exact;column:is_delete;table:withdraw_account" comment:"是否删除"`
}

func (m *WithdrawAccountSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *WithdrawAccountSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *WithdrawAccountSearch) Generate() dto.Index {
	o := *m
	return &o
}

type WithdrawAccountControl struct {
	ID uint `uri:"ID" comment:"id"` // id

	Identity string `json:"identity" comment:"用户id"`

	Type string `json:"type" comment:"类型"`

	Account string `json:"account" comment:"账号"`

	AccountName string `json:"accountName" comment:"账号用户名"`

	Status string `json:"status" comment:"状态"`

	IsDelete string `json:"isDelete" comment:"是否删除"`

	CreateTime time.Time `json:"createTime" comment:"创建时间"`

	UpdateTime time.Time `json:"updateTime" comment:"更新时间"`
}

func (s *WithdrawAccountControl) Bind(ctx *gin.Context) error {
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

func (s *WithdrawAccountControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *WithdrawAccountControl) GenerateM() (common.ActiveRecord, error) {
	return &models.WithdrawAccount{

		ID:          s.ID,
		Identity:    s.Identity,
		Type:        s.Type,
		Account:     s.Account,
		AccountName: s.AccountName,
		Status:      s.Status,
		IsDelete:    s.IsDelete,
		CreateTime:  s.CreateTime,
		UpdateTime:  s.UpdateTime,
	}, nil
}

func (s *WithdrawAccountControl) GetId() interface{} {
	return s.ID
}

type WithdrawAccountById struct {
	dto.ObjectById
}

func (s *WithdrawAccountById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *WithdrawAccountById) GenerateM() (common.ActiveRecord, error) {
	return &models.WithdrawAccount{}, nil
}
