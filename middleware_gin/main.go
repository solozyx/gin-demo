package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request
		// 要么是下一重中间件
		// 要么是实际的action handler
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	// Default默认使用 engine.Use(Logger(), Recovery()) 中间件
	// r := gin.Default()
	r := gin.New()

	// 使用中间件
	// 默认把日志打印到 控制台 这里打印到日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		// example := c.MustGet("example").(string)
		name := c.DefaultQuery("name", "default_name")
		panic("test gin.Recovery middleware")
		c.String(http.StatusOK, "%s", name)
	})

	r.Run(":8080")
}
