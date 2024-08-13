package xray

import (
	"encoding/json"
	"fmt"
)

type XRayConfig struct {
	API         *APIConfig     `json:"api,omitempty"`
	Stats       interface{}    `json:"stats,omitempty"`
	Policy      *PolicyConfig  `json:"policy,omitempty"`
	Inbounds    []*Inbound     `json:"inbounds,omitempty"`
	Routing     *RoutingConfig `json:"routing,omitempty"`
	Log         *LogConfig     `json:"log,omitempty"`
	PeerIP      string         `json:"-"`
	SslCertFile string         `json:"-"`
	SslKeyFile  string         `json:"-"`
}

type APIConfig struct {
	Services []string `json:"services"`
	Tag      string   `json:"tag"`
}

type PolicyConfig struct {
	Levels map[string]PolicyLevel `json:"levels"`
	System PolicySystem           `json:"system"`
}

type PolicyLevel struct {
	StatsUserUplink   bool `json:"statsUserUplink"`
	StatsUserDownlink bool `json:"statsUserDownlink"`
}

type PolicySystem struct {
	StatsInboundDownlink  bool `json:"statsInboundDownlink"`
	StatsInboundUplink    bool `json:"statsInboundUplink"`
	StatsOutboundDownlink bool `json:"statsOutboundDownlink"`
	StatsOutboundUplink   bool `json:"statsOutboundUplink"`
}

type Inbound struct {
	Listen         string          `json:"listen"`
	Port           int             `json:"port"`
	Protocol       string          `json:"protocol"`
	Settings       InboundSettings `json:"settings"`
	StreamSettings StreamSettings  `json:"streamSettings"`
	Tag            string          `json:"tag"`
}

type InboundSettings struct {
	Address string `json:"address"`
}

type StreamSettings struct {
	Security    string       `json:"security"`
	TlsSettings *TlsSettings `json:"tlsSettings"`
}

type TlsSettings struct {
	Certificates []TlsCertificate `json:"certificates"`
}

type TlsCertificate struct {
	CertificateFile string `json:"certificateFile"`
	KeyFile         string `json:"keyFile"`
}

type RoutingConfig struct {
	Rules []RoutingRule `json:"rules"`
}

type RoutingRule struct {
	InboundTag  []string `json:"inboundTag"`
	Source      []string `json:"source"`
	OutboundTag string   `json:"outboundTag"`
	Type        string   `json:"type"`
}

type LogConfig struct {
	LogLevel string `json:"logLevel"`
}

func NewConfig(configJson, peerIP string) (*XRayConfig, error) {
	var config XRayConfig
	if err := json.Unmarshal([]byte(configJson), &config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %v", err)
	}

	config.PeerIP = peerIP
	config.applyAPI()
	return &config, nil
}

func (c *XRayConfig) ToJson() string {
	bytes, _ := json.Marshal(c)
	return string(bytes)
}

func (c *XRayConfig) applyAPI() {
	c.API = &APIConfig{
		Services: []string{"HandlerService", "StatsService", "LoggerService"},
		Tag:      "API",
	}
	c.Stats = map[string]interface{}{}
	c.Policy = &PolicyConfig{
		Levels: map[string]PolicyLevel{
			"0": {
				StatsUserUplink:   true,
				StatsUserDownlink: true,
			},
		},
		System: PolicySystem{
			StatsInboundDownlink:  false,
			StatsInboundUplink:    false,
			StatsOutboundDownlink: true,
			StatsOutboundUplink:   true,
		},
	}
	c.Inbounds = []*Inbound{
		{
			Listen:   "0.0.0.0",
			Port:     8080,
			Protocol: "dokodemo-door",
			Settings: InboundSettings{
				Address: "127.0.0.1",
			},
			StreamSettings: StreamSettings{
				Security: "tls",
				TlsSettings: &TlsSettings{
					Certificates: []TlsCertificate{
						{
							CertificateFile: c.SslCertFile,
							KeyFile:         c.SslKeyFile,
						},
					},
				},
			},
			Tag: "API_INBOUND",
		},
	}

	c.Routing = &RoutingConfig{
		Rules: []RoutingRule{
			{
				InboundTag:  []string{"API_INBOUND"},
				Source:      []string{"127.0.0.1", c.PeerIP},
				OutboundTag: "API",
				Type:        "field",
			},
		},
	}
}
