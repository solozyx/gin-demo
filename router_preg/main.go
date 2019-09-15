package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()

	// 泛绑定 /user 前缀 ,所有/user的请求都打到该回调函数
	router.GET("/user/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.Run(":8080")
}
