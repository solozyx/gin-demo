package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	fmt.Println("getting")
}

func posting(c *gin.Context) {
	fmt.Println("posting")
}

func putting(c *gin.Context) {
	fmt.Println("putting")
}

func deleting(c *gin.Context) {
	fmt.Println("deleting")
}

func main() {
	r := gin.Default()

	r.GET("/someGet", getting)

	r.POST("/somePost", posting)

	r.DELETE("/someDelete", deleting)

	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})

	r.PUT("/somePut", putting)

	//router.PATCH("/somePatch", patching)
	//router.HEAD("/someHead", head)
	//router.OPTIONS("/someOptions", options)

	// 匹配任意请求类型 支持
	r.Any("/any", func(c *gin.Context) {
		c.String(http.StatusOK, "any")
	})

	r.Run() // default 8080
	// r.Run(":3000")
}

/*
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) IRoutes {
	group.handle("GET", relativePath, handlers)
	group.handle("POST", relativePath, handlers)
	group.handle("PUT", relativePath, handlers)
	group.handle("PATCH", relativePath, handlers)
	group.handle("HEAD", relativePath, handlers)
	group.handle("OPTIONS", relativePath, handlers)
	group.handle("DELETE", relativePath, handlers)
	group.handle("CONNECT", relativePath, handlers)
	group.handle("TRACE", relativePath, handlers)
	return group.returnObj()
}
*/
