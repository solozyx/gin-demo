package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.POST("/user/*action", func(c *gin.Context) {
	// router.GET("/user/*action", func(c *gin.Context) {
	router.Any("/user/*action", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}

		// TODO:NOTICE 执行完 ioutil.ReadAll() Form表单的数据就拿不到了
		//  全部的数据已经读到 bodyBytes了 把数据回存到 c.Request.Body
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "kai")
		c.String(http.StatusOK, "firstName = %s, lastName = %s, body = %s", firstName, lastName, string(bodyBytes))
	})

	router.Run(":8080")
}

/*
[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx" -d "first_name=solo&last_name=zyx"
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    95  100    66  100    29     66     29  0:00:01 --:--:--  0:00:01  1507
firstName = , lastName = kai, body = first_name=solo&last_name=zyx

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx" -d '{"first_name":"solo","last_name":"zyx"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100   103  100    70  100    33     70     33  0:00:01 --:--:--  0:00:01  100k
firstName = , lastName = kai, body = '{first_name:solo,last_name:zyx}'

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx?first_name=f&last_name=l" -d '{"first_name":"solo,"last_name":"zyx"}'
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100   103  100    70  100    33     70     33  0:00:01 --:--:--  0:00:01  100k
firstName = , lastName = kai, body = '{first_name:solo,last_name:zyx}'

*/

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/user/xxx" -d "first_name=solo&last_name=zyx"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    95  100    66  100    29     66     29  0:00:01 --:--:--  0:00:01  1532
firstName = , lastName = kai, body = first_name=solo&last_name=zyx

[C:\~]$ curl -X GET "http://127.0.0.1:8080/user/xxx?first_name=f&last_name=l" -d '{"first_name":"solo,"last_name":"zyx"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   103  100    70  100    33     70     33  0:00:01 --:--:--  0:00:01  100k
firstName = , lastName = kai, body = '{first_name:solo,last_name:zyx}'

*/

/*
// TODO:NOTICE 执行完 ioutil.ReadAll() Form表单的数据就拿不到了
//  全部的数据已经读到 bodyBytes了 把数据回存到 c.Request.Body
c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx" -d "first_name=solo&last_name=zyx"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    99  100    70  100    29     70     29  0:00:01 --:--:--  0:00:01 99000
firstName = solo, lastName = zyx, body = first_name=solo&last_name=zyx

[C:\~]$ curl -X GET "http://127.0.0.1:8080/user/xxx?first_name=f&last_name=l" -d '{"first_name":"solo,"last_name":"zyx"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   103  100    70  100    33     70     33  0:00:01 --:--:--  0:00:01  100k
firstName = , lastName = kai, body = '{first_name:solo,last_name:zyx}'

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx?first_name=f&last_name=l" -d '{"first_name":"solo,"last_name":"zyx"}'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   103  100    70  100    33     70     33  0:00:01 --:--:--  0:00:01  100k
firstName = , lastName = kai, body = '{first_name:solo,last_name:zyx}'

*/
