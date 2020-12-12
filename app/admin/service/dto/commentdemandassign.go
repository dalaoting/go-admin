package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/pkg/models"
	"go-admin/tools"
)

type CommentDemandAssignSearch struct {
	dto.Pagination `search:"-"`
	Serial         string `form:"serial" search:"type:exact;column:serial;table:comment_demand_assign" comment:"任务编号"`

	DemandSerial string `form:"demandSerial" search:"type:exact;column:demand_serial;table:comment_demand_assign" comment:"需求编号"`

	Commentator string `form:"commentator" search:"type:exact;column:commentator;table:comment_demand_assign" comment:"测评人"`

	Status string `form:"status" search:"type:exact;column:status;table:comment_demand_assign" comment:"状态"`
}

func (m *CommentDemandAssignSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CommentDemandAssignSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *CommentDemandAssignSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CommentDemandAssignControl struct {
	ID uint `uri:"ID" comment:""` //

	Serial string `json:"serial" comment:"任务编号"`

	DemandSerial string `json:"demandSerial" comment:"需求编号"`

	DemandSnapshotCode string `json:"demandSnapshotCode" comment:"快照ID"`

	DeptId string `json:"deptId" comment:"企业"`

	Commentator string `json:"commentator" comment:"测评人"`

	CommentAccount string `json:"commentAccount" comment:"测评账号"`

	CommentName string `json:"commentName" comment:"测评账号名称"`

	Status string `json:"status" comment:"状态"`

	TipsStatus string `json:"tipsStatus" comment:"提示"`

	IncomeAccount string `json:"incomeAccount" comment:"收款账号"`

	IncomeType string `json:"incomeType" comment:"收款类型"`

	IncomeName string `json:"incomeName" comment:"收款名称"`

	OrderMedias string `json:"orderMedias" comment:"下单截图"`

	CommentMedias string `json:"commentMedias" comment:"留评截图"`

	SettleMedias string `json:"settleMedias" comment:"结算截图"`
}

func (s *CommentDemandAssignControl) Bind(ctx *gin.Context) error {
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

func (s *CommentDemandAssignControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentDemandAssignControl) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentDemandAssign{

		Model:              gorm.Model{ID: s.ID},
		Serial:             s.Serial,
		DemandSerial:       s.DemandSerial,
		DemandSnapshotCode: s.DemandSnapshotCode,
		DeptId:             s.DeptId,
		Commentator:        s.Commentator,
		CommentAccount:     s.CommentAccount,
		CommentName:        s.CommentName,
		Status:             s.Status,
		TipsStatus:         s.TipsStatus,
		IncomeAccount:      s.IncomeAccount,
		IncomeType:         s.IncomeType,
		IncomeName:         s.IncomeName,
		OrderMedias:        s.OrderMedias,
		CommentMedias:      s.CommentMedias,
		SettleMedias:       s.SettleMedias,
	}, nil
}

func (s *CommentDemandAssignControl) GetId() interface{} {
	return s.ID
}

type CommentDemandAssignById struct {
	dto.ObjectById
	SerialNumber string `uri:"serial"`
}

func (s *CommentDemandAssignById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentDemandAssignById) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentDemandAssign{}, nil
}

func (s *CommentDemandAssignById) GetId() interface{} {
	return s.SerialNumber
}
