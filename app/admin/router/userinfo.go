package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/userinfo"
	"go-admin/common/actions"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerUserInfoRouter)
}

// 需认证的路由代码
func registerUserInfoRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &userinfo.UserInfo{}
	r := v1.Group("/userinfo").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetUserInfoList)
		r.GET("/:id", api.GetUserInfo)
		r.POST("", api.InsertUserInfo)
		r.PUT("/:id", api.UpdateUserInfo)
		r.DELETE("", api.DeleteUserInfo)
	}
}
