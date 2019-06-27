package service

import (
	"CommodityService/entity"
	"CommodityService/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 查询列表
func FindAll() []entity.Codetype {
	codetype := make([]entity.Codetype, 0)
	err := util.Engine.Find(&codetype)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(codetype)
	return codetype
}

// 条件查询
func ConditionCodeType(codeType string) (total int64) {
	item := entity.Codetype{}
	var err error
	total, err = util.Engine.Where("code_type=?", codeType).Limit(1).Count(&item)
	if err != nil {
		fmt.Println(err)
	}
	return total
}

// 添加code_type
func AddCodeType(item *entity.Codetype) (err error) {
	_, err = util.Engine.Insert(item)
	return err
}

// 删除 code_type
func DeleteCodeType(codeType string) (status int, info gin.H) {
	total, err := util.Engine.Where("code_type=?", codeType).Count(&entity.Code{})
	if err != nil {
		return util.ErrorSystem(err)
	}
	if total > 0 {
		return util.Error(util.CODE_INVALID, "数据字典下面有数据被引用，不能进行删除", nil)
	}
	_, err = util.Engine.Where("code_type=?", codeType).Delete(&entity.Codetype{})
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(nil)
}
