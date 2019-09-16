package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/test", func(c *gin.Context) {
		// 不带默认值的参数
		firstName := c.Query("first_name")
		// 带默认值的参数
		lastName := c.DefaultQuery("last_name", "default")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})

	router.GET("/user/*action", func(c *gin.Context) {
		firstName := c.DefaultQuery("first_name", "wang")
		lastName := c.DefaultQuery("last_name", "kai")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})

	router.Run(":8080")
}

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/test?first_name=zhao"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    12  100    12    0     0     12      0  0:00:01 --:--:--  0:00:01   193
zhao,default

[C:\~]$ curl -X GET "http://127.0.0.1:8080/user/xxx"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     8  100     8    0     0      8      0  0:00:01 --:--:--  0:00:01  8000
wang,kai

[C:\~]$ curl -X GET "http://127.0.0.1:8080/user/111"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100     8  100     8    0     0      8      0  0:00:01 --:--:--  0:00:01   533
wang,kai
*/
