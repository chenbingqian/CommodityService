package router

import (
	api "CommodityService/controller"

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

	codetype.GET("/list", api.FindCodeTypeList)
	codetype.POST("/insert", api.InsertCodeType)
	codetype.POST("/del", api.DeleteCodeType)

	code := group.Group("/code")
	code.GET("/list", api.FindCodeList)
	code.POST("/insert", api.InsertCode)
	code.POST("/del", api.DelCode)
	code.POST("/list/by/codetype", api.SelectCodeByType)

	attribute := group.Group("/attribute")
	attribute.POST("/condition/list", api.ConditionAttributeList)
	attribute.POST("/insert", api.InsertAttribute)
	attribute.POST("/del", api.DelAttribute)
	attribute.POST("/update", api.UpdateAttribute)

	return router
}
