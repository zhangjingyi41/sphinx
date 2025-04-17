package service

import (
	"sphinx/internal/db/mysql"

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
