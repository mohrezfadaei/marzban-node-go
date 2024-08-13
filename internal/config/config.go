package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServicePort        int    `envconfig:"SERVICE_PORT" default:"62050"`
	XrayApiPort        int    `envconfig:"XRAY_API_PORT" default:"62051"`
	XrayExecutablePath string `envconfig:"XRAY_EXECUTABLE_PATH" default:"/usr/local/bin/xray"`
	XrayAssetsPath     string `envconfig:"XRAY_ASSETS_PATH" default:"/usr/local/share/xray"`
	SslCertFile        string `envconfig:"SSL_CERT_FILE" default:"/var/lib/marzban-node/ssl_cert.pem"`
	SslKeyFile         string `envconfig:"SSL_KEY_FILE" default:"/var/lib/marzban-node/ssl_key.pem"`
	SslClientCertFile  string `envconfig:"SSL_CLIENT_CERT_FILE" default:""`
	Debug              bool   `envconfig:"DEBUG" default:"false"`
	ServiceProtocol    string `envconfig:"SERVICE_PROTOCOL" default:"rpyc"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
