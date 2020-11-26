package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"

	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/pkg/models"
	"go-admin/tools"
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

	Identity string `json:"identity" comment:"用户id"`

	Account string `json:"account" comment:"账号"`

	Password string `json:"password" comment:"密码"`

	NickName string `json:"nickName" comment:"昵称"`

	Avatar string `json:"avatar" comment:"头像"`

	Profession string `json:"profession" comment:"职业"`

	Sex string `json:"sex" comment:"性别"`

	Remark string `json:"remark" comment:"备注"`

	Status string `json:"status" comment:"状态"`

	ForbidDeadline time.Time `json:"forbidDeadline" comment:"禁用时间"`

	CreateTime time.Time `json:"createTime" comment:"创建时间"`

	UpdateTime time.Time `json:"updateTime" comment:"更新时间"`
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

		Model:          gorm.Model{ID: s.ID},
		Identity:       s.Identity,
		Account:        s.Account,
		Password:       s.Password,
		NickName:       s.NickName,
		Avatar:         s.Avatar,
		Profession:     s.Profession,
		Sex:            s.Sex,
		Remark:         s.Remark,
		Status:         s.Status,
		ForbidDeadline: s.ForbidDeadline,
		CreateTime:     s.CreateTime,
		UpdateTime:     s.UpdateTime,
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
