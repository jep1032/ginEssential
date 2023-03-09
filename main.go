package main

import (
	"fmt"
	"ginessential/common"
	"github.com/gin-gonic/gin"
)

func main() {
	_ = common.InitDb()

	r := gin.Default()
	fmt.Println("master 修改")
	r = CollectRoute(r)
	panic(r.Run())
}
