package service

import (
	"CommodityService/entity"
	"CommodityService/util"

	"github.com/gin-gonic/gin"
)

// condition select attribute
func ConditionAttributeList(item *entity.Attribute) (status int, data gin.H) {
	attribute := make([]entity.Attribute, 0)
	sessionItem := util.Engine.NewSession()
	if item.AttributeTypeId != -1 {
		sessionItem = sessionItem.Where("attribute_type_id =? ", item.AttributeTypeId)
	}
	err := sessionItem.OrderBy("attribute_id").Find(&attribute)
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(attribute)
}

// insert attribute
func InsertAttribute(item *entity.Attribute) (status int, info gin.H) {
	_, err := util.Engine.Insert(item)
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(nil)
}

// delete attribute
func DelAttribute(item *entity.Attribute) (status int, info gin.H) {
	_, err := util.Engine.Where("attribute_id = ?", item.AttributeId).Delete(&entity.Attribute{})
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(nil)
}
