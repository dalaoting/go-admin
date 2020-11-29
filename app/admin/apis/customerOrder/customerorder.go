package customerOrder

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/service"
	common "go-admin/common/models"
	"go-admin/pkg/constant"
	"go-admin/pkg/service/customerService"
	"net/http"

	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	"go-admin/common/apis"
	"go-admin/common/log"
	"go-admin/pkg/models"
	"go-admin/tools"
)

type CustomerOrder struct {
	apis.Api
}

func (e *CustomerOrder) GetCustomerOrderList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	d := new(dto.CustomerOrderSearch)
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

	list := make([]models.CustomerOrder, 0)
	var count int64
	serviceStudent := service.CustomerOrder{}
	serviceStudent.MsgID = msgID
	serviceStudent.Orm = db
	err = serviceStudent.GetCustomerOrderPage(req, p, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	var (
		createUserId = make([]int, 0)
	)
	for i := range list {
		createUserId = append(createUserId, int(list[i].CreateBy), int(list[i].UpdateBy))
		list[i].CreateTime = list[i].CreatedAt.Format(constant.DefaultTimeFormat)
	}
	user := &models.SysUser{}
	userMap, _ := user.BatchGet(createUserId)
	for i := range list {
		if v, ok := userMap[int(list[i].CreateBy)]; ok {
			list[i].CreateUser = v.NickName
		} else {
			list[i].CreateUser = "用户"
		}

		if v, ok := userMap[int(list[i].UpdateBy)]; ok {
			list[i].UpdateUser = v.NickName
		} else {
			list[i].UpdateUser = ""
		}
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e *CustomerOrder) GetCustomerOrder(c *gin.Context) {
	control := new(dto.CustomerOrderById)
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
	var object models.CustomerOrder

	//数据权限检查
	p := actions.GetPermissionFromContext(c)

	serviceCustomerOrder := service.CustomerOrder{}
	serviceCustomerOrder.MsgID = msgID
	serviceCustomerOrder.Orm = db
	err = serviceCustomerOrder.GetCustomerOrder(req, p, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

func (e *CustomerOrder) InsertCustomerOrder(c *gin.Context) {
	control := new(dto.CustomerOrderControl)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}
	// 添加客户充值记录时，需要给客户添加额度
	tx := db.Set("gorm:query_option", "FOR UPDATE")

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := control.Generate()
	err = req.Bind(c)
	if err != nil {
		tx.Callback()
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		tx.Callback()
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}
	// 设置创建人
	object.SetCreateBy(tools.GetUserIdUint(c))

	serviceCustomerOrder := service.CustomerOrder{}
	serviceCustomerOrder.Orm = tx
	serviceCustomerOrder.MsgID = msgID
	err = serviceCustomerOrder.InsertCustomerOrder(object)
	if err != nil {
		tx.Callback()
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}
	// 是否需要增加客户余额
	if control.IsAddBalance {
		if err = customerService.OperateCustomerWithTx(tx, &models.CustomerOperation{
			CustomerId: int(control.CustomerId),
			Amount:     control.Amount,
			OpType:     models.OpTypeAdd,
			BsType:     models.BsTypeRecharge,
			Detail:     "客户充值金额",
			Ext:        fmt.Sprintf("%v", object.GetId()),
			CreateBy:  tools.GetUserIdUint(c),
		}); err != nil {
			tx.Callback()
			log.Error(err)
			e.Error(c, http.StatusInternalServerError, err, "创建失败")
			return
		}
	}

	tx.Commit()
	e.OK(c, object.GetId(), "创建成功")
}

func (e *CustomerOrder) UpdateCustomerOrder(c *gin.Context) {
	control := new(dto.CustomerOrderControl)
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

	serviceCustomerOrder := service.CustomerOrder{}
	serviceCustomerOrder.Orm = db
	serviceCustomerOrder.MsgID = msgID
	err = serviceCustomerOrder.UpdateCustomerOrder(object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

func (e *CustomerOrder) DeleteCustomerOrder(c *gin.Context) {
	control := new(dto.CustomerOrderById)
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

	serviceCustomerOrder := service.CustomerOrder{}
	serviceCustomerOrder.Orm = db
	serviceCustomerOrder.MsgID = msgID
	err = serviceCustomerOrder.RemoveCustomerOrder(req, object, p)
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}
