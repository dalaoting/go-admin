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

type BlockListSearch struct {
	dto.Pagination     `search:"-"`
    Account string `form:"account" search:"type:exact;column:account;table:block_list" comment:"账号"`

    Origin string `form:"origin" search:"type:exact;column:origin;table:block_list" comment:"来源"`

    
}

func (m *BlockListSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *BlockListSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *BlockListSearch) Generate() dto.Index {
	o := *m
	return &o
}

type BlockListControl struct {
    
    ID uint `uri:"ID" comment:""` // 

    Account string `json:"account" comment:"账号"`
    

    Type string `json:"type" comment:"类型"`
    

    Remark string `json:"remark" comment:"备注，封禁原因"`
    

    Origin string `json:"origin" comment:"来源"`
    
}

func (s *BlockListControl) Bind(ctx *gin.Context) error {
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

func (s *BlockListControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *BlockListControl) GenerateM() (common.ActiveRecord, error) {
	return &models.BlockList{
	
        Model:     gorm.Model{ID: s.ID},
        Account:  s.Account,
        Type:  s.Type,
        Remark:  s.Remark,
        Origin:  s.Origin,
	}, nil
}

func (s *BlockListControl) GetId() interface{} {
	return s.ID
}

type BlockListById struct {
	dto.ObjectById
}

func (s *BlockListById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *BlockListById) GenerateM() (common.ActiveRecord, error) {
	return &models.BlockList{}, nil
}
