package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yottachain/YTSGX/controller"
)

//InitRouter 初始化路由
func InitRouter() (router *gin.Engine) {
	router = gin.Default()
	gin.SetMode(gin.DebugMode)
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	v1 := router.Group("/api/v1")
	{
		v1.GET("/addUser", controller.AddUser)
		v1.GET("/getObject", controller.DownloadFileForSGX)
		v1.GET("/getRandPublicKey", controller.GetPubKey)
		v1.GET("/getUserInfo", controller.GetUserInfo)
		v1.GET("/updateUserInfo", controller.UpdateUserInfo)
		v1.GET("/getExcelUsers", controller.GetExcelUsers)
		v1.GET("/getAllExcelUsers", controller.GetAllAuthExcelUsers)
	}

	return
}
