package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-server-template/src/userAuth"
)

func UserAuthRouter(g *gin.RouterGroup) {
	g.GET("/login", userAuth.UserLoginController)
}
