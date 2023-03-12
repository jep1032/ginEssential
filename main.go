package main

import (
	"ginessential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = common.InitDb()

	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())
}
