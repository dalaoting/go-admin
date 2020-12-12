package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/media"
	"go-admin/common/actions"
	jwt "go-admin/pkg/jwtauth"
	"go-admin/pkg/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerMediaRouter)
}

// 需认证的路由代码
func registerMediaRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &media.Media{}
	r := v1.Group("/media").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.POST("/upload", api.UploadFile)
	}
}
