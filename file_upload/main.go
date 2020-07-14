package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/uploadfiles", func(c *gin.Context) {
		// 获取表单数据 参数为name值
		f, err := c.FormFile("f1")
		// 错误处理
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		} else {
			//将文件保存至本项目根目录中
			c.SaveUploadedFile(f, f.Filename)
			//保存成功返回正确的Json数据
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		}
	})

	// 运行 默认为80端口
	r.Run()
}
