package controller

import (
	"CommodityService/entity"
	"CommodityService/service"
	"CommodityService/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// condition select attribute
func ConditionAttributeList(in *gin.Context) {
	var info struct {
		entity.Attribute
		util.PageInfo
	}
	err := in.ShouldBind(&info)
	fmt.Println(info.Attribute.AttributeId)
	fmt.Println(info.PageInfo.RowCount)
	fmt.Println(info.PageInfo.RowStartNumber)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.ConditionAttributeList(&info.Attribute, &info.PageInfo))
}

// insert attribute
func InsertAttribute(in *gin.Context) {
	var item entity.Attribute
	err := in.Bind(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.InsertAttribute(&item))
}

// delete attribute
func DelAttribute(in *gin.Context) {
	var item entity.Attribute
	err := in.Bind(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.DelAttribute(&item))
}

func UpdateAttribute(in *gin.Context) {
	var item entity.Attribute
	err := in.ShouldBind(&item)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.UpdateAttribute(&item))
}
