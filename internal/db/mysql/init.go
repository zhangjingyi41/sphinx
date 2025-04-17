package mysql

import (
	"fmt"
	"sphinx/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var AuthDB *sqlx.DB

func InitAuthDB(config *configs.MysqlConfig) (err error) {
	fmt.Println("初始化AuthDB")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	if AuthDB, err = sqlx.Connect("mysql", dsn); err != nil {
		return err
	}
	AuthDB.SetMaxOpenConns(config.MaxOpenConnections)
	AuthDB.SetMaxIdleConns(config.MaxIdleConnections)
	fmt.Println("AuthDB初始化完成")
	return nil
}
