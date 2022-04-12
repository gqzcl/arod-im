package service

import (
	"context"

	pb "arod-im/api/connector/v1"
)

type ConnectorService struct {
	pb.UnimplementedConnectorServer
}

func NewConnectorService() *ConnectorService {
	return &ConnectorService{}
}

func (s *ConnectorService) CreateConnector(ctx context.Context, req *pb.CreateConnectorRequest) (*pb.CreateConnectorReply, error) {
	return &pb.CreateConnectorReply{}, nil
}
func (s *ConnectorService) UpdateConnector(ctx context.Context, req *pb.UpdateConnectorRequest) (*pb.UpdateConnectorReply, error) {
	return &pb.UpdateConnectorReply{}, nil
}
func (s *ConnectorService) DeleteConnector(ctx context.Context, req *pb.DeleteConnectorRequest) (*pb.DeleteConnectorReply, error) {
	return &pb.DeleteConnectorReply{}, nil
}
func (s *ConnectorService) GetConnector(ctx context.Context, req *pb.GetConnectorRequest) (*pb.GetConnectorReply, error) {
	return &pb.GetConnectorReply{}, nil
}
func (s *ConnectorService) ListConnector(ctx context.Context, req *pb.ListConnectorRequest) (*pb.ListConnectorReply, error) {
	return &pb.ListConnectorReply{}, nil
}
