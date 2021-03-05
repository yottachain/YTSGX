package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yottachain/YTSGX/controller"
)

//InitRouter 初始化路由
func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/getInfo", controller.GetInfo)

	}

	return
}
