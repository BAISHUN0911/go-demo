package configx

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
}

type ServerConfig struct {
	Port int
}

var (
	MySQL  MysqlConfig
	Server ServerConfig
)

func InitConfig() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("读取配置失败:", err)
	}

	MySQL = MysqlConfig{
		Host:         viper.GetString("mysql.host"),
		Port:         viper.GetInt("mysql.port"),
		User:         viper.GetString("mysql.user"),
		Password:     viper.GetString("mysql.password"),
		DBName:       viper.GetString("mysql.dbname"),
		MaxOpenConns: viper.GetInt("mysql.max_open_conns"),
		MaxIdleConns: viper.GetInt("mysql.max_idle_conns"),
		MaxLifetime:  viper.GetInt("mysql.max_lifetime"),
	}

	Server = ServerConfig{
		Port: viper.GetInt("server.port"),
	}

	_ = time.Second
}
