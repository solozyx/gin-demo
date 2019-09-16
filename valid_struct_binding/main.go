package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	// 验证规则 使用 binding 标签
	// 需要同时满足n个条件使用 , 分割
	// 满足n个的1个即可 使用 | 分割
	// 规则 age name address 必传, age > 10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Address  string    `form:"address" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/test", startPage)
	r.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	// 先把请求参数绑定到struct,再基于struct做验证,在struct增加验证规则
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		c.Abort()
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}
