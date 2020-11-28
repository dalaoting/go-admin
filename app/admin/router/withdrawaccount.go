package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/withdrawaccount"
	"go-admin/common/actions"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerWithdrawAccountRouter)
}

// 需认证的路由代码
func registerWithdrawAccountRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &withdrawaccount.WithdrawAccount{}
	r := v1.Group("/withdrawaccount").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetWithdrawAccountList)
		r.GET("/:id", api.GetWithdrawAccount)
		r.POST("", api.InsertWithdrawAccount)
		r.PUT("/:id", api.UpdateWithdrawAccount)
		r.DELETE("", api.DeleteWithdrawAccount)
	}
}
