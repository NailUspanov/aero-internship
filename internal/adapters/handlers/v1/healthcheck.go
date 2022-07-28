package v1

import (
	"aero-internship/gen/api"
	"aero-internship/internal/domain/usecase"
	"context"
)

type HealthCheckHandler struct {
	api.UnimplementedContentCheckServiceServer
	service usecase.Service
}

func NewHealthCheckHandler(service usecase.Service) *HealthCheckHandler {
	return &HealthCheckHandler{service: service}
}

func (s *HealthCheckHandler) CheckHealth(ctx context.Context, request *api.EmptyRequest) (*api.HealthResponse, error) {
	return &api.HealthResponse{
		ServiceName:   "Some service",
		ServiceStatus: "200",
	}, nil
}
