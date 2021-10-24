package userRouter

import (
	"github.com/gin-gonic/gin"
	// "go-server-template/middleware/global/jwt"
)

// 初始化模块路由
func UserInitRouter(r *gin.RouterGroup) {
	userAuth := r.Group("/user-auth")
	// userAuth.Use(JWTMiddleware.JWT())
	UserAuthRouter(userAuth)
}
