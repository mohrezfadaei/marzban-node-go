package services

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/mohrezfadaei/marzban-node-go/internal/config"
	"github.com/mohrezfadaei/marzban-node-go/internal/logger"
	"github.com/mohrezfadaei/marzban-node-go/internal/xray"
	"github.com/mohrezfadaei/marzban-node-go/proto/xrayservice"
)

type XrayServiceServer struct {
	xrayservice.UnimplementedXrayServiceServer
	core *xray.XRayCore
}

func (s *XrayServiceServer) Start(ctx context.Context, req *xrayservice.StartRequest) (*xrayservice.StartResponse, error) {
	config, err := xray.NewConfig(req.Config, "127.0.0.1")
	if err != nil {
		return nil, err
	}
	err = s.core.Start(config)
	if err != nil {
		return nil, err
	}
	return &xrayservice.StartResponse{Message: "Xray started successfully"}, nil
}

func (s *XrayServiceServer) Stop(ctx context.Context, req *xrayservice.StopRequest) (*xrayservice.StopResponse, error) {
	err := s.core.Stop()
	if err != nil {
		return nil, err
	}
	return &xrayservice.StopResponse{Message: "Xray stopped successfully"}, nil
}

func (s *XrayServiceServer) Restart(ctx context.Context, req *xrayservice.RestartRequest) (*xrayservice.RestartResponse, error) {
	config, err := xray.NewConfig(req.Config, "127.0.0.1")
	if err != nil {
		return nil, err
	}
	err = s.core.Restart(config)
	if err != nil {
		return nil, err
	}
	return &xrayservice.RestartResponse{Message: "Xray restarted successfully"}, nil
}

func (s *XrayServiceServer) FetchXrayVersion(ctx context.Context, req *xrayservice.FetchXrayVersionRequest) (*xrayservice.FetchXrayVersionResponse, error) {
	version := s.core.GetVersion()
	return &xrayservice.FetchXrayVersionResponse{Version: version}, nil
}

func (s *XrayServiceServer) FetchLogs(req *xrayservice.FetchLogsRequest, stream xrayservice.XrayService_FetchLogsServer) error {
	logs := s.core.GetLogs()
	for log := range logs {
		if err := stream.Send(&xrayservice.LogMessage{Log: log}); err != nil {
			return err
		}
	}
	return nil
}

func NewXrayServiceServer(core *xray.XRayCore) *XrayServiceServer {
	return &XrayServiceServer{core: core}
}

func RunGRPCServer(conf *config.Config, core *xray.XRayCore) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.ServicePort))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	xrayservice.RegisterXrayServiceServer(s, NewXrayServiceServer(core))
	logger.Info.Printf("gRPC server listening on port %d", conf.ServicePort)
	return s.Serve(lis)
}
