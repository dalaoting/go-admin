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

type CommentDemandSearch struct {
	dto.Pagination `search:"-"`
	SerialNumber   string `form:"serialNumber" search:"type:exact;column:serial_number;table:comment_demand" comment:"编号"`

	CustomerId string `form:"customerId" search:"type:exact;column:customer_id;table:comment_demand" comment:"客户"`

	Status string `form:"status" search:"type:exact;column:status;table:comment_demand" comment:"状态"`
}

func (m *CommentDemandSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CommentDemandSearch) Bind(ctx *gin.Context) error {
	msgID := tools.GenerateMsgIDFromContext(ctx)
	err := ctx.ShouldBind(m)
	if err != nil {
		log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
	}
	return err
}

func (m *CommentDemandSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CommentDemandControl struct {
	ID uint `uri:"ID" comment:""` //

	DeptId string `json:"deptId" comment:"企业"`

	SerialNumber string `json:"serialNumber" comment:"编号"`

	ShopName string `json:"shopName" comment:"店铺名"`

	CustomerId string `json:"customerId" comment:"客户"`

	Title string `json:"title" comment:"标题"`

	ProductCode string `json:"productCode" comment:"商品码"`

	Reward string `json:"reward" comment:"佣金(分)"`

	ProductPrice string `json:"productPrice" comment:"价格"`

	CommentNum string `json:"commentNum" comment:"测评数"`

	Desc string `json:"desc" comment:"说明"`

	Remark string `json:"remark" comment:"备注"`

	Status string `json:"status" comment:"状态"`
}

func (s *CommentDemandControl) Bind(ctx *gin.Context) error {
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

func (s *CommentDemandControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentDemandControl) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentDemand{

		Model:        gorm.Model{ID: s.ID},
		DeptId:       s.DeptId,
		SerialNumber: s.SerialNumber,
		CustomerId:   s.CustomerId,
		Title:        s.Title,
		ShopName:     s.ShopName,
		ProductCode:  s.ProductCode,
		Reward:       s.Reward,
		ProductPrice: s.ProductPrice,
		CommentNum:   s.CommentNum,
		Desc:         s.Desc,
		Remark:       s.Remark,
		Status:       s.Status,
	}, nil
}

func (s *CommentDemandControl) GetId() interface{} {
	return s.ID
}

type CommentDemandById struct {
	dto.ObjectById
}

func (s *CommentDemandById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CommentDemandById) GenerateM() (common.ActiveRecord, error) {
	return &models.CommentDemand{}, nil
}
