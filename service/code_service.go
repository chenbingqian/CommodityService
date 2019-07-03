package service

import (
	"CommodityService/entity"
	"CommodityService/util"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

// select all code list
func FindCodeAll() []entity.Code {
	code := make([]entity.Code, 0)
	err := util.Engine.OrderBy("code_type,value").Find(&code)
	if err != nil {
		fmt.Println(err)
	}
	return code
}

// select code by code_type
func SelectCodeByType(codeType string) (status int, info gin.H) {

	code := make([]entity.Code, 0)
	err := util.Engine.Where("code_type = ?", codeType).Find(&code)
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(code)
}

// insert code
func InsertCode(item *entity.Code) (status int, info gin.H) {
	total, err := util.Engine.Where("code_type = ? ", item.Codetype).And(" (name =? or value = ?)", item.Name, item.Value).Count(&entity.Code{})
	if total != 0 {
		err = errors.New("{'result_code':20000,'msg':'数据不合法'}")
		return util.Error(util.CODE_INVALID, util.CODE_INVALID_MSG, nil)
	}
	_, err = util.Engine.Insert(item)
	if err != nil {
		return util.ErrorSystem(err)
	}
	return util.Success(nil)
}

// del code
func DelCode(item *entity.Code) (err error) {
	fmt.Println(item)
	_, err = util.Engine.Where("code_type = ?", item.Codetype).And("value=?", item.Value).Delete(&entity.Code{})
	return err
}
