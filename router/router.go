package router

import (
	. "CommodityService/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	cors_config := cors.DefaultConfig()
	cors_config.AllowAllOrigins = true
	cors_config.AddAllowHeaders("Authorization")
	cors_config.AddAllowHeaders("access_token")
	cors_config.AddAllowHeaders("OPTIONS")
	cors_config.AddAllowHeaders("instance_id")
	cors_config.AddAllowHeaders("application_name")
	router.Use(cors.New(cors_config))

	group := router.Group("/kiki/v1")
	codetype := group.Group("/codetype")

	codetype.GET("/list", FindUserList)
	codetype.POST("/insert", InsertCodeType)
	codetype.POST("/del", DeleteCodeType)

	return router
}
