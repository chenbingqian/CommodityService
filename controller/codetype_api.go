package controller

import (
	"CommodityService/entity"
	"CommodityService/service"
	"fmt"

	"CommodityService/util"

	"github.com/gin-gonic/gin"
)

// 查询所有codetype数据
func FindCodeTypeList(in *gin.Context) {
	result := service.FindAll()
	in.JSON(util.Success(result))
}

// 添加code_type
func InsertCodeType(in *gin.Context) {
	// 完成数据映射
	var codeType entity.Codetype
	if err := in.Bind(&codeType); err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}

	total := service.ConditionCodeType(codeType.Codetype)
	fmt.Println(total)
	if total != 0 {
		fmt.Println("log: 已经添加过")
		in.JSON(util.Error(util.CODE_INVALID, util.CODE_INVALID_MSG, nil))
		return
	}

	service.AddCodeType(&codeType)

	in.JSON(util.Success(nil))
}

// delete code_type method
func DeleteCodeType(in *gin.Context) {
	var item struct {
		CodeType string `json:"code_type"`
	}
	if err := in.Bind(&item); err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.DeleteCodeType(item.CodeType))

}
