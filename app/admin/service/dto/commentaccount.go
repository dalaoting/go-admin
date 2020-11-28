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

type CommentAccountSearch struct {
	dto.Pagination `search:"-"`
	Identity       string `form:"identity" search:"type:exact;column:identity;table:comment_account" comment:"用户id"`

	Type string `form:"type" search:"type:exact;column:type;table:comment_account" comment:"类型"`

	Account string `form:"account" search:"type:exact;column:account;table:comment_account" comment:"账号"`
}

func (m *CommentAccountSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CommentAccountSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *CommentAccountSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CommentAccountControl struct {
	ID uint `uri:"ID" comment:"id"` // id

	Identity string `json:"identity" comment:"用户id"`

	Type string `json:"type" comment:"类型"`

	Account string `json:"account" comment:"账号"`

	AccountName string `json:"accountName" comment:"用户名"`

	Status string `json:"status" comment:"状态"`

	IsDelete string `json:"isDelete" comment:"已删除"`

	CreateTime time.Time `json:"createTime" comment:"创建时间"`

	UpdateTime time.Time `json:"updateTime" comment:"更新时间"`
}

func (s *CommentAccountControl) Bind(ctx *gin.Context) error {
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

func (s *CommentAccountControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentAccountControl) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentAccount{

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

func (s *CommentAccountControl) GetId() interface{} {
	return s.ID
}

type CommentAccountById struct {
	dto.ObjectById
}

func (s *CommentAccountById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentAccountById) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentAccount{}, nil
}
