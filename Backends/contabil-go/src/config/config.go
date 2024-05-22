package config

import (
	"os"
	"path/filepath"

	"github.com/TalisonK/TalisonContabil/src/util"
	"github.com/spf13/viper"
)

type config struct {
	API     APIConfig
	Auth    AuthConfig
	DBLocal DBConfig
	DBCloud DBConfig
}

type AuthConfig struct {
	Key                  string
	Google_Client_id     string
	Google_Client_Secret string
}

type APIConfig struct {
	Port   string
	IsProd bool
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

// Load reads the configuration from the file config.toml
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	configPath, err := findConfigFile()

	if err != nil {
		util.LogHandler("Fail to find the config file", err, "config.Load")
		return err
	}

	viper.AddConfigPath(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port:   viper.GetString("api.port"),
		IsProd: viper.GetBool("api.is_prod"),
	}

	databaseLocal := viper.Sub("database.mysql")
	databaseCloud := viper.Sub("database.mongodb")
	auth := viper.Sub("auth")

	cfg.DBLocal = DBConfig{
		Host:     databaseLocal.GetString("host"),
		Port:     databaseLocal.GetString("port"),
		User:     databaseLocal.GetString("user"),
		Pass:     databaseLocal.GetString("pass"),
		Database: databaseLocal.GetString("database"),
	}

	cfg.DBCloud = DBConfig{
		Host:     databaseCloud.GetString("host"),
		User:     databaseCloud.GetString("user"),
		Pass:     databaseCloud.GetString("pass"),
		Database: databaseCloud.GetString("database"),
	}

	cfg.Auth = AuthConfig{
		Key:                  auth.GetString("key"),
		Google_Client_id:     auth.GetString("google_client_id"),
		Google_Client_Secret: auth.GetString("google_client_secret"),
	}

	return nil
}

func GetLocalDB() DBConfig {
	return cfg.DBLocal
}

func GetCloudDB() DBConfig {
	return cfg.DBCloud
}

func GetAuthConfig() AuthConfig {
	return cfg.Auth
}

func IsProd() bool {
	return cfg.API.IsProd
}

func GetServerPort() string {
	return cfg.API.Port
}

// FindConfigFile searches for the "config.toml" file in the parent directories
// until it finds the file or reaches the root directory.
func findConfigFile() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		configPath := filepath.Join(dir, "config.toml")
		_, err := os.Stat(configPath)
		if err == nil {
			return dir, nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}

	return "", os.ErrNotExist
}
