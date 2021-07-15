package config

import "github.com/spf13/viper"

type Config struct {
	HTTPServer HTTPServerConfig
	Postgres   PostgresConfig
	EthConfig  EthConfig
	Fiat       FiatConfig
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

type HTTPServerConfig struct {
	Host string
	Port int
}

type EthConfig struct {
	EthNetwork  string
	HezRollup   string
	UsdtAddress string
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
	}
}
