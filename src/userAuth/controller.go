package userAuth

import (
	"github.com/gin-gonic/gin"
)

// @Summary 登陆方法
// @Produce  json
// @Param account query string true "account"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/user-auth/login [post]
func UserLoginController(ctx *gin.Context) {
	myData := UserInfo{ID: 2}
	UserLoginService(ctx, myData)
}
