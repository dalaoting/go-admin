package commentdemandassign

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	common "go-admin/common/models"
	"net/http"

	"go-admin/pkg/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/tools"
)

type CommentDemandAssign struct {
	apis.Api
}

func (e *CommentDemandAssign) GetCommentDemandAssignList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.CommentDemandAssignSearch)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := d.Generate()

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	list := make([]models.CommentDemandAssign, 0)
	var count int64
	serviceStudent := service.CommentDemandAssign{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetCommentDemandAssignPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e *CommentDemandAssign) GetCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object models.CommentDemandAssign

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.MsgID = msgID
	serviceCommentDemandAssign.Orm = db
	err = serviceCommentDemandAssign.GetCommentDemandAssign(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

func (e *CommentDemandAssign) InsertCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(tools.GetUserIdUint(c))

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.InsertCommentDemandAssign(object)
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

func (e *CommentDemandAssign) UpdateCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	req := control.Generate()
	//更新操作
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	object.SetUpdateBy(tools.GetUserIdUint(c))

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.UpdateCommentDemandAssign(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

func (e *CommentDemandAssign) DeleteCommentDemandAssign(c *gin.Context) {
	control := new(dto.CommentDemandAssignById)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//删除操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	// 设置编辑人
	object.SetUpdateBy(tools.GetUserIdUint(c))

	// 数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCommentDemandAssign := service.CommentDemandAssign{}
	serviceCommentDemandAssign.Orm = db
	serviceCommentDemandAssign.MsgID = msgID
	err = serviceCommentDemandAssign.RemoveCommentDemandAssign(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}