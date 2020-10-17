package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/customer"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCustomerRouter)
}

// 需认证的路由代码
func registerCustomerRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &customer.Customer{}
	r := v1.Group("/customer").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetCustomerList)
		r.GET("/:id", api.GetCustomer)
		r.POST("", api.InsertCustomer)
		r.PUT("/:id", api.UpdateCustomer)
		r.DELETE("", api.DeleteCustomer)
	}
}
