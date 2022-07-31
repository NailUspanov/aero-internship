package v1

import (
	"aero-internship/gen/api"
	"aero-internship/internal/domain/usecase"
	"context"
)

type RequestNewsObject struct {
	api.UnimplementedNewsServiceServer
	service usecase.NewsService
}

func (r *RequestNewsObject) Create(ctx context.Context, request *api.RequestNewsObject) (*api.BaseResponse, error) {
	return &api.BaseResponse{
		Success: true,
		Message: "200",
	}, nil
}
