package router

import (
	"fmt"
	"net/http"
	"sphinx/logger"

	"github.com/gin-gonic/gin"
)

func StartServer(mode string) *gin.Engine {
	fmt.Println("初始化web服务器")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	// 中间件，先不管
	server.Use(
		logger.GinLogger(),       // 接收gin框架默认的日志
		logger.GinRecovery(true), // recover掉项目可能出现的panic，并使用zap记录相关日志
	)
	// 静态文件

	server.StaticFile("/favicon.ico", "static/favicon.ico") // favicon图标
	server.Static("/assets", "static/assets")               // 静态文件目录
	server.LoadHTMLFiles("static/index.html")               // 加载html文件
	// server.LoadHTMLGlob("./static/*") // 加载html文件夹下的所有html文件
	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Sphinx",
		})
	})

	// swagger路由先不管

	// 加载其他模块的路由
	AuthRouter(server, "/api")

	fmt.Println("web服务器初始化完成")
	return server
}
