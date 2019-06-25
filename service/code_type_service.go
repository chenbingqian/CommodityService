package service

import (
	"CommodityService/entity"
	"CommodityService/util"
	"fmt"
)

func FindAll() []entity.Codetype {
	codetype := make([]entity.Codetype, 0)
	err := util.Engine.Find(&codetype)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(codetype)
	return codetype
}
