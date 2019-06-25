package router

import (
	. "CommodityService/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/codetype/v1/list", FindUserList)
	return router
}
