package service

import (
	"CommodityService/entity"
	"CommodityService/util"
	"fmt"
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

// 添加code_type
func AddCodeType(item *entity.Codetype) {
	util.Engine.Insert(item)
}

// 删除 code_type
func DeleteCodeType(codeType string) {

	util.Engine.Where("code_type=?", codeType).Delete(&entity.Codetype{})
}
