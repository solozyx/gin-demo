package main

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"gopkg.in/go-playground/validator.v8"
)

type Booking struct {
	// 自定义验证标签 bookabledate
	CheckIn time.Time `form:"check_in"  binding:"required,bookabledate"    time_format:"2006-01-02"`
	// gtfield=CheckIn 登出字段比登入字段 时间大 要写结构体字段CheckIn而不是tag名
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn" time_format:"2006-01-02"`
}

// 自定义验证标签 bookabledate 验证器
func bookableDate(
	v *validator.Validate,
	topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value, // 该参数获取请求对应参数
	fieldType reflect.Type,
	fieldKind reflect.Kind,
	param string) bool {
	// 类型断言
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		// 预约时间 > 今天
		if today.Unix() < date.Unix() {
			return true
		}
	}
	return false
}

func main() {
	r := gin.Default()
	// 自定义验证规则
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}
	r.GET("/bookable", getBookable)
	r.Run(":8080")
}

func getBookable(c *gin.Context) {
	var b Booking
	// if err := c.ShouldBindWith(&b, binding.Query); err == nil {
	if err := c.ShouldBind(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Booking dates success :)", "booking": b})
}

/*
[C:\~]$ curl -X GET "http://127.0.0.1:8080/bookable?check_in=2019-09-01&check_out=2019-09-30"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   104  100   104    0     0    104      0  0:00:01 --:--:--  0:00:01  3250
{"error":"Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"}

[C:\~]$ curl -X GET "http://127.0.0.1:8080/bookable?check_in=2019-09-25&check_out=2019-09-30"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   127  100   127    0     0    127      0  0:00:01 --:--:--  0:00:01  2702
{"booking":{"CheckIn":"2019-09-25T00:00:00+08:00","CheckOut":"2019-09-30T00:00:00+08:00"},"message":"Booking dates success :)"}

[C:\~]$ curl -X GET "http://127.0.0.1:8080/bookable?check_in=2019-09-25&check_out=2019-09-22"
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100   101  100   101    0     0    101      0  0:00:01 --:--:--  0:00:01  6733
{"error":"Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"}

*/
