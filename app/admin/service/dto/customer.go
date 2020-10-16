package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-admin/pkg/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type CustomerSearch struct {
	dto.Pagination     `search:"-"`
    ID uint `form:"ID" search:"type:exact;column:id;table:customer" comment:"客户ID"`

    Status string `form:"status" search:"type:exact;column:status;table:customer" comment:"状态"`

    CreateBy string `form:"createBy" search:"type:exact;column:create_by;table:customer" comment:"创建人"`

    UpdateBy string `form:"updateBy" search:"type:exact;column:update_by;table:customer" comment:"更新人"`

    CreatedAt time.Time `form:"createdAt" search:"type:gte;column:created_at;table:customer" comment:""`

    
}

func (m *CustomerSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *CustomerSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *CustomerSearch) Generate() dto.Index {
	o := *m
	return &o
}

type CustomerControl struct {
    
    ID uint `uri:"ID" comment:"客户ID"` // 客户ID

    DeptId string `json:"deptId" comment:"企业ID"`
    

    Name string `json:"name" comment:"客户名称"`
    

    AvailAmount string `json:"availAmount" comment:"当前余额，包含预扣冻结"`
    

    TotalAmount string `json:"totalAmount" comment:"历史入账总数"`
    

    PrepayAmount string `json:"prepayAmount" comment:"预扣冻结金额，avail_amount减去prepay_amount等于实际可用余额"`
    

    Status string `json:"status" comment:"状态"`
    
}

func (s *CustomerControl) Bind(ctx *gin.Context) error {
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

func (s *CustomerControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CustomerControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Customer{
	
        Model:     gorm.Model{ID: s.ID},
        DeptId:  s.DeptId,
        Name:  s.Name,
        AvailAmount:  s.AvailAmount,
        TotalAmount:  s.TotalAmount,
        PrepayAmount:  s.PrepayAmount,
        Status:  s.Status,
	}, nil
}

func (s *CustomerControl) GetId() interface{} {
	return s.ID
}

type CustomerById struct {
	dto.ObjectById
}

func (s *CustomerById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *CustomerById) GenerateM() (common.ActiveRecord, error) {
	return &models.Customer{}, nil
}
