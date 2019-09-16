package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Any("/user/*action", func(c *gin.Context) {
		firstName := c.DefaultPostForm("first_name", "wang")
		lastName := c.PostForm("last_name")
		c.String(http.StatusOK, "firstName = %s, lastName = %s", firstName, lastName)
	})

	router.Run(":8080")
}

/*

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx" -d "first_name=solo&last_name=zyx"
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    61  100    32  100    29     32     29  0:00:01 --:--:--  0:00:01 61000
firstName = solo, lastName = zyx

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx" -d '{"first_name":"solo,"last_name":"zyx"}'
% Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
Dload  Upload   Total   Spent    Left  Speed
100    62  100    29  100    33     29     33  0:00:01 --:--:--  0:00:01 62000
firstName = wang, lastName =

[C:\~]$ curl -X POST "http://127.0.0.1:8080/user/xxx?first_name=f&last_name=l"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100    29  100    29    0     0     29      0  0:00:01 --:--:--  0:00:01 29000
firstName = wang, lastName =

*/
