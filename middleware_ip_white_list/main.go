package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("example", "example_123456")

		ipList := []string{
			// db -> cache -> ipList
			"127.0.0.2",
		}

		isMatched := false
		for _, ip := range ipList {
			if c.ClientIP() == ip {
				isMatched = true
				break
			}
		}

		if !isMatched {
			c.String(401, fmt.Sprintf("%v, not in ip_white_list", c.ClientIP()))
			c.Abort()
			return
		}
	}
}

func main() {
	r := gin.Default()
	r.Use(IPAuthMiddleware())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)
		c.String(http.StatusOK, "hello world")
	})

	r.Run(":8080")
}

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/test"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    31  100    31    0     0     31      0  0:00:01 --:--:--  0:00:01   500
127.0.0.1, not in ip_white_list

*/
