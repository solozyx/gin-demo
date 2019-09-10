package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// 路由 回调函数
	r.GET("/ping", func(c *gin.Context) {
		// 返回JSON数据
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 默认 listen and serve on 0.0.0.0:8080
}

/*
http://localhost:8080/ping
{"message":"pong"}
*/
