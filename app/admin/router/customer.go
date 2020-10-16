package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/pkg/middleware"

	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
	"go-admin/pkg/models"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCustomerRouter)
}

// 需认证的路由代码
func registerCustomerRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r := v1.Group("/customer").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		model := &models.Customer{}
		r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.CustomerSearch), func() interface{} {
			list := make([]models.Customer, 0)
			return &list
		}))
		r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.CustomerById), nil))
		r.POST("", actions.CreateAction(new(dto.CustomerControl)))
		r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.CustomerControl)))
		r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.CustomerById)))
	}
}
