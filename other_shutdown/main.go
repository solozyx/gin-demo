package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Router")
	})

	// 不能使用 gin.Run 而是用 net/http 构建server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 把构建的server放到1个goroutine执行
	go func() {
		// 注意 ListenAndServe 是非阻塞的 要配合 os.Signal
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// os.Signal 信号
	quit := make(chan os.Signal)
	// 捕获退出信号 SIGINT [Ctrl + C] 和 SIGTERM [kill -15] 这2类信号可以捕获
	// kill -9 不能捕获到
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// 超时上下文
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	// 真正关闭server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown err = ", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
