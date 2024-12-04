package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	DB     DBConfig     `mapstructure:"db"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DBConfig struct {
	Driver   string `mapstructure:"driver"`
	URL      string `mapstructure:"url"`
	MaxConns int    `mapstructure:"max_conns"`
}

func Load() (*Config, error) {
	// load .env file
	if err := gotenv.Load(); err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("db.driver", "sqlite3")
	viper.SetDefault("db.max_conns", 10)

	// bind environment variables
	viper.BindEnv("db.url", "DATABASE_URL")
	viper.BindEnv("server.port", "PORT")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil

}
