package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}
	Clickhouse struct {
		Host     string
		Port     string
		Database string
		User     string
		Password string
	}
	Websocket struct {
		ReadBufferSize  int
		WriteBufferSize int
	}
	Data struct {
		DefaultThrottleMs int
	}
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	
	config.Server.Port = viper.GetString("SERVER_PORT")
	config.Server.Host = viper.GetString("SERVER_HOST")
	
	config.Clickhouse.Host = viper.GetString("CLICKHOUSE_HOST")
	config.Clickhouse.Port = viper.GetString("CLICKHOUSE_PORT")
	config.Clickhouse.Database = viper.GetString("CLICKHOUSE_DATABASE")
	config.Clickhouse.User = viper.GetString("CLICKHOUSE_USER")
	config.Clickhouse.Password = viper.GetString("CLICKHOUSE_PASSWORD")
 
	config.Websocket.ReadBufferSize = viper.GetInt("WS_READ_BUFFER_SIZE")
	config.Websocket.WriteBufferSize = viper.GetInt("WS_WRITE_BUFFER_SIZE")
	
	config.Data.DefaultThrottleMs = viper.GetInt("DEFAULT_THROTTLE_MS")

	return &config, nil
} 