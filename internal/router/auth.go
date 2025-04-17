package router

import (
	"sphinx/internal/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(server *gin.Engine) {
	auth := server.Group("/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.POST("/logout", controller.Logout)
	}

	// oauth := server.Group("/oauth")
	// {
	// 	oauth.GET("/authorize", controller.Authorize) // 授权端点
	// 	oauth.POST("/token", controller.Token)        // 令牌端点
	// }

}
