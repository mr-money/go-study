package main

import (
	"go-study/Routes"
)

//
//  main
//  @Description: 入口
//
func main() {
	// 加载路由
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api) //TODO api路由，需要token中间件验证

	//
	//// 监听端口，默认在8080
	// Run(":8000")
	_ = Routes.GinEngine.Run()
}
