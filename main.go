package main

import (
	"ginessential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = common.InitDb()

	r := gin.Default()
	fmt.Pringln("master test")
	r = CollectRoute(r)
	panic(r.Run())
}
