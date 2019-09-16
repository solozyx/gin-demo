package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	// form 标签把json格式转换为form表单格式
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
	route := gin.Default()
	route.GET("/test", startPage)
	route.POST("/test", startPage)
	route.Run(":8080")
}

func startPage(c *gin.Context) {
	var person Person
	// 根据前端请求 GET/POST 类型 Content-Type 动态binding解析请求数据
	// GET 把参数识别为 name=x?address=y&birthday=z
	// POST 把参数识别为表单或json
	if err := c.ShouldBind(&person); err != nil {
		c.String(http.StatusOK, fmt.Sprintf("%v", err))
	}
	c.String(http.StatusOK, fmt.Sprintf("%v", person))
}

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/test?name=solo&address=china&birthday=2000-01-01
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    42  100    42    0     0     42      0  0:00:01 --:--:--  0:00:01 42000
{solo china 2000-01-01 00:00:00 +0000 UTC}

[C:\~]$ curl -X POST "http://127.0.0.1:8080/test?name=solo&address=china&birthday=2000-01-01
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    42  100    42    0     0     42      0  0:00:01 --:--:--  0:00:01 42000
{solo china 2000-01-01 00:00:00 +0000 UTC}

[C:\~]$ curl -X POST "http://127.0.0.1:8080/test" -d "name=solo&address=china&birthday=2000-01-01"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    85  100    42  100    43     42     43  0:00:01 --:--:--  0:00:01  5312
{solo china 2000-01-01 00:00:00 +0000 UTC}

[C:\~]$ curl -X GET "http://127.0.0.1:8080/test" -d "name=solo&address=china&birthday=2000-01-01"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    76  100    33  100    43     33     43  0:00:01 --:--:--  0:00:01 76000
{  0001-01-01 00:00:00 +0000 UTC}

*/
