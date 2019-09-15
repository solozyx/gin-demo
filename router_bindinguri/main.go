package main

import "github.com/gin-gonic/gin"

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	router := gin.Default()
	// 获取url的 name id 参数
	router.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": c.Param("name"), "uuid": c.Param("id")})

		var person Person
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	router.Run(":8080")
}

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/solo/1007"
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    58  100    58    0     0     58      0  0:00:01 --:--:--  0:00:01 58000
{"name":"solo","uuid":"1007"}{"name":"solo","uuid":"1007"}
*/
