package config

import (
	"time"

	"github.com/spf13/cobra"
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

func setDefaultValues() {
	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 3000)
	// Goerli values
	viper.SetDefault("ETH_HEZ_ROLLUP", "0xf08a226B67a8A9f99cCfCF51c50867bc18a54F53")
	viper.SetDefault("ETH_USDT_ADDRESS", "0x509ee0d083ddf8ac028f2a56731412edd63223b9")
	viper.SetDefault("BBOLT_LOCATION", "/tmp/price_update.db")
	viper.SetDefault("MAIN_TIME_TO_UPDATE_PRICES", "30s")
}

func readDotEnvWithViper() error {
	setDefaultValues()

	viper.SetConfigType("env")
	viper.SetConfigFile("./.env")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func LoadByFlags(flags *cobra.Command) (Config, error) {
	setDefaultValues()

	viper.BindPFlag("POSTGRES_USER", flags.Flags().Lookup("pg-user"))
	viper.BindPFlag("POSTGRES_PASSWORD", flags.Flags().Lookup("pg-pass"))
	viper.BindPFlag("POSTGRES_HOST", flags.Flags().Lookup("pg-host"))
	viper.BindPFlag("POSTGRES_PORT", flags.Flags().Lookup("pg-port"))
	viper.BindPFlag("POSTGRES_DATABASE", flags.Flags().Lookup("pg-dbname"))
	viper.BindPFlag("ETH_NETWORK", flags.Flags().Lookup("eth-network"))
	viper.BindPFlag("FIAT_API_KEY", flags.Flags().Lookup("fiat-key"))

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
	}, nil
}

func LoadByEnv() (Config, error) {
	err := readDotEnvWithViper()
	if err != nil {
		return Config{}, err
	}

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
	}, nil
}
