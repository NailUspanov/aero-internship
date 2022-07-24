package v1

import (
	"aero-internship/gen/api"
	"context"
)

type HealthCheckHandler struct {
	api.UnimplementedContentCheckServiceServer
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (s *HealthCheckHandler) CheckHealth(ctx context.Context, request *api.EmptyRequest) (*api.HealthResponse, error) {
	return &api.HealthResponse{
		ServiceName:   "Some service",
		ServiceStatus: "200",
	}, nil
}
