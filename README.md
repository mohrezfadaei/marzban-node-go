# Marzban Node Go

Marzban Node Go is a Go-based application that provides node services using gRPC and REST protocols. This project includes a multi-platform Docker setup and a `Makefile` to simplify the build and deployment process.

## Prerequisites

- [Docker](https://www.docker.com/get-started)

## Getting Started

### Docker

#### Building Docker Image

To build the Docker image for the local architecture:

```sh
docker build -t marzban-node-go . 
```

### Environment Variables

The application configuration is loaded from environment variables. Below is an example `.env` file:

```env
SERVICE_PORT=62050
XRAY_API_PORT=62051
XRAY_EXECUTABLE_PATH=/usr/local/bin/xray
XRAY_ASSETS_PATH=/usr/local/share/xray
SSL_CERT_FILE=/var/lib/marzban-node/ssl_cert.pem
SSL_KEY_FILE=/var/lib/marzban-node/ssl_key.pem
SSL_CLIENT_CERT_FILE=/path/to/ssl_client_cert.pem
SERVICE_PROTOCOL=rpyc
DEBUG=false
```

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.
