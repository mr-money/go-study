package main

import (
	"context"
	"go-study/Routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

//
//  main
//  @Description: 入口
//
func main() {
	// 加载路由
	Routes.Include(
		Routes.Web, //默认web路由
		Routes.Api, //api路由，需要token中间件验证
	)

	//
	//// 监听端口，默认在8080
	// Run(":8000")
	//_ = Routes.GinEngine.Run()

	//优雅关闭
	srv := &http.Server{
		Addr:    ":8080",
		Handler: Routes.GinEngine,
	}

	shutdown(srv)
}

//
// shutdown
// @Description: 优雅关闭服务
// @param srv
//
func shutdown(srv *http.Server) {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}

	log.Println("Server exiting")
}
