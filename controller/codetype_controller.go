package controller

import (
	"CommodityService/entity"
	"CommodityService/service"
	"fmt"

	"CommodityService/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询所有code数据
func FindUserList(in *gin.Context) {
	result := service.FindAll()

	in.JSON(http.StatusOK, gin.H{
		"result_code": util.CODE_SUCCES,
		"data":        result,
	})
}

// 添加code_type
func InsertCodeType(in *gin.Context) {
	var codeType entity.Codetype
	err := in.Bind(&codeType)
	if err != nil {
		fmt.Println(err)
	}

	service.AddCodeType(&codeType)

	in.JSON(http.StatusOK, gin.H{
		"result_code": util.CODE_SUCCES,
		"msg":         util.CODE_SUCCES_MSG,
	})
}

// delete code_type method
func DeleteCodeType(in *gin.Context) {
	//var codeType = in.Request.FormValue("code_type")
	//codeType := in.PostForm("code_type")
	//fmt.Println(codeType)
	//if codeType == "" {
	// 错误
	//}
	var d struct {
		CodeType string `json:"code_type"`
	}

	if err := in.Bind(&d); err != nil {
		fmt.Println(err)
	}

	//fmt.Println(d)
	//fmt.Println("111111111111111111")

	service.DeleteCodeType(d.CodeType)
	in.JSON(http.StatusOK, gin.H{
		"result_code": util.CODE_SUCCES,
		"msg":         util.CODE_SUCCES_MSG,
	})

}
