package config

import "github.com/spf13/viper"

type Config struct {
	Mongo      MongoConfig
	HTTPServer HTTPServerConfig
}

type MongoConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type HTTPServerConfig struct {
	Host string
	Port int
}

func readDotEnvWithViper() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(".env file not found")
	}
}

func Load() Config {
	readDotEnvWithViper()

	return Config{
		Mongo: MongoConfig{
			User:     viper.GetString("MONGO_USER"),
			Password: viper.GetString("MONGO_PASSWORD"),
			Host:     viper.GetString("MONGO_HOST"),
			Port:     viper.GetInt("MONGO_PORT"),
			Database: viper.GetString("MONGO_DATABASE"),
		},
		HTTPServer: HTTPServerConfig{
			Host: viper.GetString("HTTP_HOST"),
			Port: viper.GetInt("HTTP_PORT"),
		},
	}
}
