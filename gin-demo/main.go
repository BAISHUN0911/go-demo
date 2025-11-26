package main

import (
	"fmt"

	"gin-demo/configx"
	"gin-demo/db"
	"gin-demo/handler"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	configx.InitConfig()

	// 2. 初始化数据库
	db.InitMySQL()

	// 3. 初始化 Gin
	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	// 4. 路由
	r.GET("/login", handler.LoginHandler)

	// 5. 启动
	port := fmt.Sprintf(":%d", configx.Server.Port)
	r.Run(port)
}
