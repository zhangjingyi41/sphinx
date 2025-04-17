package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sphinx/internal/db/mysql"
	"sphinx/internal/models/dao"
	"time"
)

// 生成随机字符串
func generateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

// 授权码模式第一步：获取授权码
func GenerateAuthorizationCode(clientID, redirectURI string, userID int64) (string, error) {
	// 验证客户端信息
	var client dao.OAuthClient
	err := mysql.AuthDB.Get(&client, "SELECT * FROM oauth_clients WHERE client_id = ?", clientID)
	if err != nil {
		return "", errors.New("invalid client")
	}

	if client.RedirectURI != redirectURI {
		return "", errors.New("invalid redirect URI")
	}

	// 生成授权码
	code := generateRandomString(32)
	expiresAt := time.Now().Add(10 * time.Minute)

	// 保存授权码
	_, err = mysql.AuthDB.Exec(`
        INSERT INTO oauth_codes (code, client_id, user_id, redirect_uri, expires_at, created_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `, code, clientID, userID, redirectURI, expiresAt, time.Now())

	if err != nil {
		return "", err
	}

	return code, nil
}

// 授权码模式第二步：使用授权码获取访问令牌
func GenerateAccessToken(code, clientID, clientSecret string) (*dao.OAuthToken, error) {
	// 验证授权码
	var oauthCode dao.OAuthCode
	err := mysql.AuthDB.Get(&oauthCode, `
        SELECT * FROM oauth_codes 
        WHERE code = ? AND client_id = ? AND expires_at > ?
    `, code, clientID, time.Now())

	if err != nil {
		return nil, errors.New("invalid authorization code")
	}

	// 验证客户端信息
	var client dao.OAuthClient
	err = mysql.AuthDB.Get(&client, "SELECT * FROM oauth_clients WHERE client_id = ? AND client_secret = ?",
		clientID, clientSecret)
	if err != nil {
		return nil, errors.New("invalid client credentials")
	}

	// 生成访问令牌和刷新令牌
	accessToken := generateRandomString(32)
	refreshToken := generateRandomString(32)
	expiresAt := time.Now().Add(24 * time.Hour)

	// 保存令牌
	token := &dao.OAuthToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ClientID:     clientID,
		UserID:       oauthCode.UserID,
		ExpiresAt:    expiresAt,
		CreatedAt:    time.Now(),
	}

	_, err = mysql.AuthDB.Exec(`
        INSERT INTO oauth_tokens (access_token, refresh_token, client_id, user_id, expires_at, created_at)
        VALUES (?, ?, ?, ?, ?, ?)
    `, token.AccessToken, token.RefreshToken, token.ClientID, token.UserID, token.ExpiresAt, token.CreatedAt)

	if err != nil {
		return nil, err
	}

	// 删除已使用的授权码
	_, err = mysql.AuthDB.Exec("DELETE FROM oauth_codes WHERE code = ?", code)
	if err != nil {
		return nil, err
	}

	return token, nil
}
