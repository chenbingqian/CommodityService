package controller

import (
	"CommodityService/entity"
	"CommodityService/service"
	"CommodityService/util"

	"github.com/gin-gonic/gin"
)

// find all code list
func FindCodeList(in *gin.Context) {
	result := service.FindCodeAll()
	in.JSON(util.Success(result))
}

// find code by code_type
func SelectCodeByType(in *gin.Context) {
	var info struct {
		CodeType string `json:"code_type"`
	}
	err := in.Bind(&info)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.SelectCodeByType(info.CodeType))
}

// insert code
func InsertCode(in *gin.Context) {
	var item entity.Code
	err := in.Bind(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.InsertCode(&item))
}

// delete code
func DelCode(in *gin.Context) {
	var item entity.Code
	err := in.Bind(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	err = service.DelCode(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(util.Success(nil))
}
