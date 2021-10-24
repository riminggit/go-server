package authMiddleware

import (
	"github.com/gin-gonic/gin"
)

// 用户校验
func UserAuth() gin.HandlerFunc {
	//自定义逻辑

	//返回中间件
	return func(c *gin.Context) {
		//中间件逻辑
	}
}
