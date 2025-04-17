package main

import (
	"fmt"
	"sphinx/configs"
	"sphinx/internal/db/mysql"
	"sphinx/internal/router"
	"sphinx/logger"
)

func main() {
	if err := configs.InitConfig(); err != nil {
		fmt.Printf("初始化配置失败: %v\n", err)
		return
	}

	// 初始化日志组件
	if err := logger.InitLogger(configs.App.LogConfig); err != nil {
		fmt.Printf("初始化日志组件失败: %v\n", err)
		return
	}

	if err := mysql.InitAuthDB(configs.App.MysqlConfig); err != nil {
		fmt.Printf("初始化数据库失败: %v\n", err)
		return
	}
	defer mysql.AuthDB.Close()

	server := router.StartServer(configs.App.Mode)
	server.Run(fmt.Sprintf(":%d", configs.App.Port))
}
