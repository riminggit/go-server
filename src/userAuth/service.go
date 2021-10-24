package userAuth

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-server-template/pkg/redis"
	"net/http"
)

func UserLoginService(ctx *gin.Context, data UserInfo) interface{} {

	log.Info("info msg")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello 登陆成功1",
	})

	Redis.SetValue("mykey4", "11", 50)
	val := Redis.GetValue("mykey5")
	// log.Info(val, "val")

	// Redis.DelValue("mykey4")

	return val
}
