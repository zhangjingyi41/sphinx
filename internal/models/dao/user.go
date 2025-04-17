package dao

import (
	"time"
)

// User 用户表
type User struct {
	ID        int64     `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"-"` // 密码不返回给前端
	Phone     string    `db:"phone" json:"phone"`
	Status    int       `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Externalldentity struct {
	UserID         int64     `db:"user_id" json:"user_id"`
	Provider       string    `db:"provider" json:"provider"`                 // 第三方平台
	ExternalUserID string    `db:"external_user_id" json:"external_user_id"` // 第三方用户ID
	AccessToken    string    `db:"access_token" json:"access_token"`
	RefreshToken   string    `db:"refresh_token" json:"refresh_token"`
	ExpiredAt      time.Time `db:"expired_at" json:"expired_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

type OAuthClient struct {
	ClientID     string    `db:"client_id" json:"client_id"`
	ClientSecret string    `db:"client_secret" json:"client_secret"`
	RedirectURI  string    `db:"redirect_uri" json:"redirect_uri"` // 回调地址
	Scope        string    `db:"scope" json:"scope"`               // 权限范围
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type OAuthorizationCode struct {
	Code        string    `db:"code" json:"code"`
	ClientID    string    `db:"client_id" json:"client_id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	RedirectURI string    `db:"redirect_uri" json:"redirect_uri"` // 回调地址
	Scope       string    `db:"scope" json:"scope"`               // 权限范围
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	ExpiresAt   time.Time `db:"expires_at" json:"expires_at"`
}

type OAuthToken struct {
	AccessToken string    `db:"access_token" json:"access_token"`
	ClientID    string    `db:"client_id" json:"client_id"`
	UserID      int64     `db:"user_id" json:"user_id"`
	ExpiresAt   time.Time `db:"expires_at" json:"expires_at"`
	Scope       string    `db:"scope" json:"scope"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}
