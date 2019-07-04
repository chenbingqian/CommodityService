package service

import (
	"CommodityService/entity"
	"CommodityService/util"
	"strings"

	"github.com/gin-gonic/gin"
)

// condition select attribute
func ConditionAttributeList(item *entity.Attribute, page *util.PageInfo) (status int, data gin.H) {
	attribute := make([]entity.Attribute, 0)
	sessionItem := util.Engine.NewSession()
	//根据属性类型查询
	if item.AttributeTypeId != -1 {
		sessionItem.Where("attribute_type_id =? ", item.AttributeTypeId)
	}
	// 根据属性ID查询
	if item.AttributeId != -1 {
		sessionItem.Where("attribute_id = ?", item.AttributeId)
	}
	// 根据属性名称模糊查询
	if !strings.EqualFold(item.AttributeName, "") {
		sessionItem.Where("attribute_name like '%" + item.AttributeName + "%'")
	}
	// 总条数
	totalNumber, err := sessionItem.Count(&entity.Attribute{})
	if err != nil {
		return util.ErrorSystem(err)
	}
	page.TotalRowNumber = totalNumber

	if totalNumber > 0 {
		// 分页
		if page.RowStartNumber != -1 && page.RowCount != -1 {
			sessionItem.Limit(page.RowCount, page.RowStartNumber)
		}
		err = sessionItem.OrderBy("attribute_id").Find(&attribute)
		if err != nil {
			return util.ErrorSystem(err)
		}
	}
	return util.SuccessPage(attribute, page)
}

// insert attribute
func InsertAttribute(item *entity.Attribute) (status int, info gin.H) {

	// 使用cols用于设置默认值
	_, err := util.Engine.Cols("`attribute_id`,`attribute_name`,`attribute_type_id`,`attribute_type_name`,`attribute_value`,`attribute_value_style`,`add_time`,`update_time`").Insert(item)
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

// update
func UpdateAttribute(item *entity.Attribute) (status int, info gin.H) {
	_, err := util.Engine.Id(item.AttributeId).Update(item)
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(nil)
}
