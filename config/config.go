package config

import (
	"github.com/spf13/viper"
	"time"
)

// Config stores all configuration of the application
// The values are read by viper from a config file or environment variables
type Config struct {
	DBSource             string        `mapstructure:"DB_SOURCE"`
	SymmetricKey         string        `mapstructure:"SYMMETRIC_KEY"`
	AsaasUrl             string        `mapstructure:"ASAAS_URL"`
	AsaasApiKey          string        `mapstructure:"ASAAS_API_KEY"`
	FilePath             string        `mapstructure:"FILE_PATH"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
