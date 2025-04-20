package router

import (
	"sphinx/internal/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(server *gin.Engine, baseRoute string) {
	auth := server.Group(baseRoute)
	{
		auth.GET("/login", controller.Login)
		auth.POST("/register", controller.Register)
		auth.GET("/oauthorization", controller.OAuthorization)
	}

	// oauth := server.Group("/oauth")
	// {
	// 	oauth.GET("/authorize", controller.Authorize) // 授权端点
	// 	oauth.POST("/token", controller.Token)        // 令牌端点
	// }

}
