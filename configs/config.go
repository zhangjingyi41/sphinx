package configs

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Mode        string `mapstructure:"mode"`
	Port        int    `mapstructure:"port"`
	Version     string `mapstructure:"version"`
	StartTime   string `mapstructure:"start_time"`
	MachineCode string `mapstructure:"machine_code"`

	*AuthConfig  `mapstructure:"auth"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type AuthConfig struct {
	JwtExpireTime int64 `mapstructure:"jwt_expire"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`    // MB
	MaxAge     int    `mapstructure:"max_age"`     // 最长保存时间，单位：天
	MaxBackups int    `mapstructure:"max_backups"` // 备份数量限制
}

type MysqlConfig struct {
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DbName             string `mapstructure:"dbname"`
	MaxOpenConnections int    `mapstructure:"max_open_conns"`
	MaxIdleConnections int    `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

var App = new(AppConfig)

func InitConfig() (err error) {
	fmt.Println("正在初始化应用配置")
	var configFilePath string
	if len(os.Args) < 2 {
		configFilePath = "configs/dev.yaml"
	} else {
		configFilePath = os.Args[1]
	}
	viper.SetConfigFile(configFilePath)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(App); err != nil {
		return err
	}

	// 监听配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变化")
		if err := viper.Unmarshal(App); err != nil {
			fmt.Printf("配置文件重新加载失败: %v\n", err)
		}
	})
	fmt.Println("应用配置初始化完成")
	return nil
}
