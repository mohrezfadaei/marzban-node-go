package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/mohrezfadaei/marzban-node-go/internal/certificate"
	"github.com/mohrezfadaei/marzban-node-go/internal/config"
	"github.com/mohrezfadaei/marzban-node-go/internal/logger"
	"github.com/mohrezfadaei/marzban-node-go/internal/services"
	"github.com/mohrezfadaei/marzban-node-go/internal/xray"
)

func generateSSLFiles(certFile, keyFile string) {
	certPem, keyPem, err := certificate.GenerateCertificate()
	if err != nil {
		log.Fatalf("Failed to generate certificate: %v", err)
	}

	// Ensure the directory exists
	certDir := filepath.Dir(certFile)
	if err := os.MkdirAll(certDir, 0755); err != nil {
		log.Fatalf("Failed to create certificate directory: %v", err)
	}

	if err := os.WriteFile(certFile, []byte(certPem), 0644); err != nil {
		log.Fatalf("Failed to write certificate file: %v", err)
	}

	if err := os.WriteFile(keyFile, []byte(keyPem), 0644); err != nil {
		log.Fatalf("Failed to write key file: %v", err)
	}
}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if conf.SslCertFile == "" || conf.SslKeyFile == "" {
		log.Fatalf("SSL_CERT_FILE or SSL_KEY_FILE not specified in configuration")
	}

	// Check if SSL certificate and key files exist
	if _, err := os.Stat(conf.SslCertFile); os.IsNotExist(err) {
		generateSSLFiles(conf.SslCertFile, conf.SslKeyFile)
	} else if _, err := os.Stat(conf.SslKeyFile); os.IsNotExist(err) {
		generateSSLFiles(conf.SslCertFile, conf.SslKeyFile)
	}

	if conf.SslClientCertFile == "" {
		logger.Warn.Println("You are running node without SSL_CLIENT_CERT_FILE, be aware that everyone can connect to this node and this isn't secure!")
	}

	if conf.SslClientCertFile != "" && !fileExists(conf.SslClientCertFile) {
		logger.Error.Println("Client's certificate file specified on SSL_CLIENT_CERT_FILE is missing")
		os.Exit(1)
	}

	core := xray.NewXRayCore(conf.XrayExecutablePath, conf.XrayAssetsPath)

	switch conf.ServiceProtocol {
	case "rpyc":
		if err := services.RunGRPCServer(conf, core); err != nil {
			logger.Error.Printf("Failed to run gRPC server: %v", err)
		}
	case "rest":
		if err := services.RunRestServer(conf, core); err != nil {
			logger.Error.Printf("Failed to run REST server: %v", err)
		}
	default:
		logger.Error.Println("SERVICE_PROTOCOL is not any of (rpyc, rest).")
		os.Exit(1)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}
