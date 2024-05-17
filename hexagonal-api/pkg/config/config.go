package config

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API ApiConfig
	DB  Database
}

type Database struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

type ApiConfig struct {
	Port string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
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

	cfg.API = ApiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = Database{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDb() Database {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
