package controller

import (
	"net/http"
	"sphinx/internal/service"

	"github.com/gin-gonic/gin"
)

// 授权端点
func Authorize(ctx *gin.Context) {
	clientID := ctx.Query("client_id")
	redirectURI := ctx.Query("redirect_uri")
	responseType := ctx.Query("response_type")

	// 验证必要参数
	if clientID == "" || redirectURI == "" || responseType != "code" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	// 这里假设用户已经登录，从会话中获取用户ID
	userID := ctx.GetInt64("user_id")
	if userID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// 生成授权码
	code, err := service.GenerateAuthorizationCode(clientID, redirectURI, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 重定向回客户端
	ctx.Redirect(http.StatusFound, redirectURI+"?code="+code)
}

// 令牌端点
func Token(ctx *gin.Context) {
	grantType := ctx.PostForm("grant_type")
	code := ctx.PostForm("code")
	clientID := ctx.PostForm("client_id")
	clientSecret := ctx.PostForm("client_secret")

	// 验证必要参数
	if grantType != "authorization_code" || code == "" || clientID == "" || clientSecret == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid_request"})
		return
	}

	// 生成访问令牌
	token, err := service.GenerateAccessToken(code, clientID, clientSecret)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  token.AccessToken,
		"token_type":    "Bearer",
		"expires_in":    86400, // 24小时
		"refresh_token": token.RefreshToken,
	})
}
