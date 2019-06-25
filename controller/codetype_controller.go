package controller

import (
	"fmt"
	"net/http"
	"product-kiki-master/service"

	"github.com/gin-gonic/gin"
)

func FindUserList(in *gin.Context) {
	result1 := service.FindAll()

	fmt.Println(result1)
	in.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"data": result1,
	})
}
