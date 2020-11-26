package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/customerOrder"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
	"go-admin/pkg/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCustomerOrderRouter)
}

// 需认证的路由代码
func registerCustomerOrderRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &customerOrder.CustomerOrder{}
	r := v1.Group("/customerOrder").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetCustomerOrderList)
		r.GET("/:id", api.GetCustomerOrder)
		r.POST("", api.InsertCustomerOrder)
		r.PUT("/:id", api.UpdateCustomerOrder)
		r.DELETE("", api.DeleteCustomerOrder)
	}
}
