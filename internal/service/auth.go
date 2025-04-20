package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sphinx/internal/db/mysql"
	"sphinx/internal/models/dao"

	"github.com/google/uuid"
)

// CheckPhoneExists 检查手机号是否存在
func CheckPhoneExists(phone string) (bool, error) {
	sql := "SELECT COUNT(*) FROM users WHERE phone = ?"
	row := mysql.AuthDB.QueryRow(sql, phone)
	var count int
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// SaveAccount 保存账号密码
func SaveAccount(phone, password string) error {
	// 随机用户名
	uuid := uuid.New().String()
	username := "USER_" + uuid
	sql := "INSERT INTO users (phone, password, username) VALUES (?, ?, ?)"
	_, err := mysql.AuthDB.Exec(sql, phone, password, username)
	return err
}

// CheckAccountPassword 检查账号密码是否正确
func CheckAccountPassword(phone string, password string) (bool, error) {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	sql := "SELECT COUNT(*) FROM users WHERE phone = ? AND password = ?"

	fmt.Println(sql, phone, string(hashedPassword))
	row := mysql.AuthDB.QueryRow(sql, phone, hashedPassword)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// GenerateToken 生成token和refreshToken
// 这里可以使用JWT或其他方式生成token和refreshToken
func GenerateToken(phone string) (token string, refreshToken string, err error) {
	return "", "", nil
}

// GeneratrAuthorizationCode 生成授权码
func GeneratrAuthorizationCode(account string, client_id string, redirect_uri string, scope string, state string) (string, error) {
	// 生成随机的授权码
	// 这里可以使用UUID或其他方式生成授权码
	code := uuid.New().String()

	sql := "SELECT * FROM oauth_clients WHERE client_id = ?"
	var oclient dao.OAuthClient
	if err := mysql.AuthDB.Get(&oclient, sql, client_id); err != nil {
		return "", err
	}

	sql = "SELECT * FROM users WHERE phone = ?"
	var user dao.User
	if err := mysql.AuthDB.Get(&user, sql, account); err != nil {
		return "", err
	}

	sql = "INSERT INTO authorization_codes (code, client_id, user_id, redirect_uri, scope) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := mysql.AuthDB.Exec(sql, code, client_id, user.ID, redirect_uri, scope, state)
	if err != nil {
		return "", err
	}

	return code, nil
}
