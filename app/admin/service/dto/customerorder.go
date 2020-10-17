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

type CustomerOrderSearch struct {
	dto.Pagination `search:"-"`
	CustomerId     int64 `form:"customerId" search:"type:contains;column:customer_id;table:customer_order" comment:"客户名"`

	Transaction string `form:"transaction" search:"type:exact;column:transaction;table:customer_order" comment:"交易单号"`

	TransactionType int64 `form:"transactionType" search:"type:exact;column:transaction_type;table:customer_order" comment:"交易渠道"`

	CreateBy string `form:"createBy" search:"type:exact;column:create_by;table:customer_order" comment:"创建人"`
}

func (m *CustomerOrderSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CustomerOrderSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *CustomerOrderSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CustomerOrderControl struct {
	ID uint `uri:"ID" comment:"记录ID"` // 记录ID

	CustomerId int64 `json:"customerId,string" comment:"客户名"`

	Name string `json:"name" comment:"客户名-冗余字段"`

	Transaction string `json:"transaction" comment:"交易单号"`

	TransactionType int64 `json:"transactionType,string" comment:"交易渠道"`

	Amount int64 `json:"amount,string" comment:"交易金额"`

	Remark string `json:"remark" comment:"备注"`
}

func (s *CustomerOrderControl) Bind(ctx *gin.Context) error {
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

func (s *CustomerOrderControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CustomerOrderControl) GenerateM() (common.ActiveRecord, error) {
	return &models.CustomerOrder{

		Model:           gorm.Model{ID: s.ID},
		CustomerId:      s.CustomerId,
		Name:            s.Name,
		Transaction:     s.Transaction,
		TransactionType: s.TransactionType,
		Amount:          s.Amount,
		Remark:          s.Remark,
	}, nil
}

func (s *CustomerOrderControl) GetId() interface{} {
	return s.ID
}

type CustomerOrderById struct {
	dto.ObjectById
}

func (s *CustomerOrderById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CustomerOrderById) GenerateM() (common.ActiveRecord, error) {
	return &models.CustomerOrder{}, nil
}
