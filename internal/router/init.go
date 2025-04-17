package router

import (
	"fmt"
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
	// 静态文件，先不管

	// swagger路由先不管

	// 加载其他模块的路由
	AuthRouter(server)

	fmt.Println("web服务器初始化完成")
	return server
}
