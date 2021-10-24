package api

import (
	"github.com/gin-gonic/gin"
	"go-server-template/routers/api/userRouter"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/api")
	userRouter.UserInitRouter(api)
}
