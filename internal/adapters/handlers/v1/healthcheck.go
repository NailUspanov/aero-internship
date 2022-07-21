package v1

import (
	"aero-internship/gen/api"
	"context"
)

type GRPCServer struct {
	api.UnimplementedContentCheckServiceServer
}

func (s *GRPCServer) CheckHealth(ctx context.Context, request *api.EmptyRequest) (*api.HealthResponse, error) {
	return &api.HealthResponse{
		ServiceName:   "Some service",
		ServiceStatus: "200",
	}, nil
}
