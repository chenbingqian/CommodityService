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
	var item entity.Attribute
	err := in.Bind(&item)
	fmt.Println(&item.AttributeTypeId)
	if err != nil {
		in.JSON(util.ErrorSystem(err))
		return
	}
	in.JSON(service.ConditionAttributeList(&item))
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
