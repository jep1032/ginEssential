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
	fmt.Println("master 在网页修改 修改")
	fmt.Println("master 在goland 修改")
	r = CollectRoute(r)
	panic(r.Run())
}
