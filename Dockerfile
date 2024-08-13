FROM golang:1.22-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=$TARGETOS
ENV GOARCH=$TARGETARCH

ENV SERVICE_PORT=62050
ENV XRAY_API_PORT=62051
ENV XRAY_EXECUTABLE_PATH=/usr/local/bin/xray
ENV XRAY_ASSETS_PATH=/usr/local/share/xray
ENV SSL_CERT_FILE=/var/lib/marzban-node/ssl_cert.pem
ENV SSL_KEY_FILE=/var/lib/marzban-node/ssl_key.pem
ENV SSL_CLIENT_CERT_FILE=/var/lib/marzban-node/ssl_client_cert.pem
ENV SERVICE_PROTOCOL=rpyc
ENV DEBUG=false


RUN apk update && apk add --no-cache bash curl unzip

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o marzban-node-go cmd/main.go

FROM alpine:latest AS xray-installer

RUN apk update && apk add --no-cache curl unzip

ARG TARGETARCH

ENV ARCH=$TARGETARCH

RUN bash -c "$(curl -L https://github.com/Gozargah/Marzban-scripts/raw/master/install_latest_xray.sh)"

FROM alpine:latest

ENV PYTHONUNBUFFERED=1

WORKDIR /app

RUN apk update && apk add --no-cache ca-certificates

COPY --from=builder /app/marzban-node-go /usr/local/bin/marzban-node-go
COPY --from=xray-installer /usr/local/bin/xray /usr/local/bin/xray
COPY --from=xray-installer /usr/local/share/xray /usr/local/share/xray

COPY . .

EXPOSE 62050

CMD ["/usr/local/bin/marzban-node-go"]
