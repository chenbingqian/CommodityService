package test
import (
	
	"github.com/gin-gonic/gin"
	"net/http"
	
   )
   
   func main1() {
	

	router := gin.Default()
   
	router.GET("/", func(c *gin.Context) {
	 c.String(http.StatusOK, "It works")
	})
   
	router.Run(":8000")
   }