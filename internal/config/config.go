package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort        int    `mapstructure:"SERVICE_PORT"`
	XrayApiPort        int    `mapstructure:"XRAY_API_PORT"`
	XrayExecutablePath string `mapstructure:"XRAY_EXECUTABLE_PATH"`
	XrayAssetsPath     string `mapstructure:"XRAY_ASSETS_PATH"`
	SslCertFile        string `mapstructure:"SSL_CERT_FILE"`
	SslKeyFile         string `mapstructure:"SSL_KEY_FILE"`
	SslClientCertFile  string `mapstructure:"SSL_CLIENT_CERT_FILE"`
	Debug              bool   `mapstructure:"DEBUG"`
	ServiceProtocol    string `mapstructure:"SERVICE_PROTOCOL"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
