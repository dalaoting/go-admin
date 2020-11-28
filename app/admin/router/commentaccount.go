package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/commentaccount"
	"go-admin/common/actions"
	"go-admin/pkg/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCommentAccountRouter)
}

// 需认证的路由代码
func registerCommentAccountRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := &commentaccount.CommentAccount{}
	r := v1.Group("/commentaccount").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole(), actions.PermissionAction())
	{
		r.GET("", api.GetCommentAccountList)
		r.GET("/:id", api.GetCommentAccount)
		r.POST("", api.InsertCommentAccount)
		r.PUT("/:id", api.UpdateCommentAccount)
		r.DELETE("", api.DeleteCommentAccount)
	}
}
