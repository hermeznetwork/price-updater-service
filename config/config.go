package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServer    HTTPServerConfig
	Postgres      PostgresConfig
	EthConfig     EthConfig
	Fiat          FiatConfig
	Bbolt         BboltConfig
	ProjectConfig Main
}

type FiatConfig struct {
	APIKey string
}

type PostgresConfig struct {
	User           string
	Password       string
	Host           string
	Port           int
	Database       string
	SslModeEnabled bool
	MaxIdleConns   int
	MaxOpenConns   int
}

type BboltConfig struct {
	Location string
}

type HTTPServerConfig struct {
	Host string
	Port int
}

type EthConfig struct {
	EthNetwork  string
	HezRollup   string
	UsdtAddress string
}

type Main struct {
	TimeToUpdate time.Duration
}

func readDotEnvWithViper() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.SetDefault("MAIN_TIME_TO_UPDATE_PRICES", "30s")
	err := viper.ReadInConfig()
	if err != nil {
		panic(".env file not found")
	}
}

func Load() Config {
	readDotEnvWithViper()

	return Config{
		HTTPServer: HTTPServerConfig{
			Host: viper.GetString("HTTP_HOST"),
			Port: viper.GetInt("HTTP_PORT"),
		},
		Postgres: PostgresConfig{
			User:           viper.GetString("POSTGRES_USER"),
			Password:       viper.GetString("POSTGRES_PASSWORD"),
			Host:           viper.GetString("POSTGRES_HOST"),
			Port:           viper.GetInt("POSTGRES_PORT"),
			Database:       viper.GetString("POSTGRES_DATABASE"),
			SslModeEnabled: viper.GetBool("POSTGRES_SSL_ENABLED"),
			MaxIdleConns:   viper.GetInt("POSTGRES_MAX_ID_CONNS"),
			MaxOpenConns:   viper.GetInt("POSTGRES_MAX_OPEN_CONNS"),
		},
		EthConfig: EthConfig{
			EthNetwork:  viper.GetString("ETH_NETWORK"),
			HezRollup:   viper.GetString("ETH_HEZ_ROLLUP"),
			UsdtAddress: viper.GetString("ETH_USDT_ADDRESS"),
		},
		Fiat: FiatConfig{
			APIKey: viper.GetString("FIAT_API_KEY"),
		},
		Bbolt: BboltConfig{
			Location: viper.GetString("BBOLT_LOCATION"),
		},
		ProjectConfig: Main{
			TimeToUpdate: viper.GetDuration("MAIN_TIME_TO_UPDATE_PRICES"),
		},
	}
}
