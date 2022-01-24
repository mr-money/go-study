package main

import (
	"github.com/gin-gonic/gin"
	"go-study/router"
)

//
//  main
//  @Description: 入口
//
func main() {
	// 创建路由引擎
	r := gin.Default()

	router.Router(r)

	// 监听端口，默认在8080
	// Run(":8000")
	_ = r.Run()
}
