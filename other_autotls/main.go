package main

import (
	"log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// 不使用 r.Run
	// 实现自动化证书配置,流程 只需要调用对应的自动下载证书的包 即可
	// autotls.Run(r, "www.itpp.tk")
	// 1.生成本地密钥
	// 2.把密钥发给证书颁发机构,获取私钥
	// 3.拿到私钥后,本地进行私钥验证
	// 4.验证成功,把对应的私钥信息保存,下次请求使用该私钥进行加密,实现自动下载证书功能
	// TODO : NOTICE 实际的外网ip才行
	log.Fatal(autotls.Run(r, "www.itpp.tk"))
}
