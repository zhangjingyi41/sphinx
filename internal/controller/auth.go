package controller

import (
	"fmt"
	"sphinx/internal/models/vo"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	a := 0
	b := 1 / a
	fmt.Println(b)
	c.JSON(200, vo.Builder().Ok().Msg("login success").Finished())
}

func Logout(c *gin.Context) {
	c.JSON(200, vo.Builder().Ok().Msg("logout success").Finished())
}
func Register(c *gin.Context) {
	c.JSON(200, vo.Builder().Ok().Msg("registration success").Finished())
}
