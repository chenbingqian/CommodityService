package main

import (
	"CommodityService/router"
	_ "CommodityService/util"
)

func main() {

	router := router.InitRouter()
	router.Run(":8000")
}
