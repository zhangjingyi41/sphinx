package service

import "sphinx/internal/db/mysql"

func CheckClient(client_id, redirect_uri, scope string) (bool, error) {
	// 检查client配置是否存在
	if client_id == "" || redirect_uri == "" {
		return false, nil
	}
	if scope == "" {
		scope = "default" // 默认授权范围
	}
	sql := "SELECT COUNT(*) FROM oauth_clients WHERE client_id = ? AND scope = ?"
	row := mysql.AuthDB.QueryRow(sql, client_id, scope)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}

	return true, nil
}
