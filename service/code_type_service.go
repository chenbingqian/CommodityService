package service

import (
	"fmt"
	"product-kiki-master/entity"
	"product-kiki-master/util"
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
