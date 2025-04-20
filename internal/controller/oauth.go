package controller

import (
	"net/http"
	"sphinx/internal/models/vo"
	"sphinx/internal/service"

	"github.com/gin-gonic/gin"
)

func OAuthorization(c *gin.Context) {
	// 获取请求参数
	clientID := c.Query("client_id")
	redirectURI := c.Query("redirect_uri")
	responseType := c.Query("response_type")
	scope := c.Query("scope") // 授权范围
	state := c.Query("state") // 第三方应用发起请求时提供的随机字符串，用于防止CSRF攻击
	// 检查response type
	if responseType == "" || (responseType != "code" && responseType != "token") {
		responseType = "code" // 默认授权码模式
	}

	// 检查clientID是否有效
	exist, err := service.CheckClient(clientID, redirectURI, scope)
	if err != nil {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.RequestParamCheckFailed).Finished())
		return
	}
	if !exist {
		c.JSON(http.StatusOK, vo.Builder().Interrupted().Code(vo.InvalidClient).Finished())
		return
	}
	// client_config := &dao.OAuthClient{
	// 	ClientID:    clientID,
	// 	RedirectURI: redirectURI,
	// 	Scope:       scope,
	// }
	// 跳转到授权页面
	// 跳转到vue登录界面
	// c.Redirect(http.StatusFound, "http://localhost:5173?client_id="+clientID+"&redirect_uri="+redirectURI+"&response_type="+responseType+"&scope="+scope+"&state="+state)
	c.JSON(http.StatusOK, vo.Builder().Ok().Data(map[string]string{
		"client_id":     clientID,
		"redirect_uri":  redirectURI,
		"scope":         scope,
		"target_url":    "http://localhost:5173",
		"response_type": responseType,
		"state":         state,
	}).Finished())
}
