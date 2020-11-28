package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/blocklist"
	"go-admin/common/actions"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerBlockListRouter)
}

// 需认证的路由代码
func registerBlockListRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &blocklist.BlockList{}
	r := v1.Group("/blocklist").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetBlockListList)
		r.GET("/:id", api.GetBlockList)
		r.POST("", api.InsertBlockList)
		r.PUT("/:id", api.UpdateBlockList)
		r.DELETE("", api.DeleteBlockList)
	}
}
