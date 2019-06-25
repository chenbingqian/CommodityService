package main

import (
	"product-kiki-master/router"
	_ "product-kiki-master/util"
)

func main() {

	router := router.InitRouter()
	router.Run(":8000")
}
