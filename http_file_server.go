package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 定义ID与PW
	authoried := r.Group("/", gin.BasicAuth(gin.Accounts{
		"kevin_coding": "5Xq0mQsdrifMMoajeD5sQSo0qkA=",
	}))

	// 定义访问 url 与文件路径
	authoried.StaticFS("/file/list", http.Dir("./files"))

	// 启动服务
	r.Run(":8080")
}
