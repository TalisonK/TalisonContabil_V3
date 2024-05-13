package config

import (
	"github.com/spf13/viper"
)

type config struct {
	API     APIConfig
	DBLocal DBConfig
	DBCloud DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

var cfg *config

func init() {
	viper.SetDefault("api.port", "3033")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	databaseLocal := viper.Sub("database.mysql")
	databaseCloud := viper.Sub("database.mongodb")

	cfg.DBLocal = DBConfig{
		Host:     databaseLocal.GetString("host"),
		Port:     databaseLocal.GetString("port"),
		User:     databaseLocal.GetString("username"),
		Pass:     databaseLocal.GetString("password"),
		Database: databaseLocal.GetString("database"),
	}

	cfg.DBCloud = DBConfig{
		Host:     databaseCloud.GetString("host"),
		User:     databaseCloud.GetString("username"),
		Pass:     databaseCloud.GetString("password"),
		Database: databaseCloud.GetString("database"),
	}

	return nil
}

func GetLocalDB() DBConfig {
	return cfg.DBLocal
}

func GetCloudDB() DBConfig {
	return cfg.DBCloud
}

func GetServerPort() string {
	return cfg.API.Port
}
