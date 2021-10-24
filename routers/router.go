package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "go-server-template/docs" // 不引入会报错 Fetch error Internal Server Error doc.json
	"go-server-template/middleware/global/log"
	"go-server-template/routers/api"
	// "go-server-template/middleware/global/jwt"
)

// 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	addMiddleware(r)

	api.InitApi(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// 添加全局中间件
func addMiddleware(router *gin.Engine) {
	router.Use(logMiddleware.LogerMiddleware())
	// r.Use(JWTMiddleware.JWT())
}
