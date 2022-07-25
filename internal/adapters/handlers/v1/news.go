package v1

import (
	"aero-internship/gen/api"
	"aero-internship/internal/domain/usecase"
	"aero-internship/pkg/config"
	"context"
)

type NewsHandler struct {
	usecase.Service
	*config.Config
}

func NewNewsHandler(service usecase.Service, config *config.Config) *NewsHandler {
	return &NewsHandler{Service: service, Config: config}
}

func (n NewsHandler) GetNews(ctx context.Context, params *api.NewsRequestParams) (*api.NewsList, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) GetOne(ctx context.Context, id *api.ObjectId) (*api.NewsObject, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) GetOneBySlug(ctx context.Context, slug *api.ObjectSlug) (*api.NewsObject, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) Create(ctx context.Context, object *api.RequestNewsObject) (*api.BaseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) Update(ctx context.Context, object *api.RequestNewsObject) (*api.BaseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) Delete(ctx context.Context, id *api.ObjectId) (*api.BaseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) GetFileLink(ctx context.Context, id *api.FileId) (*api.FileLink, error) {
	//TODO implement me
	panic("implement me")
}

func (n NewsHandler) mustEmbedUnimplementedNewsServiceServer() {
	//TODO implement me
	panic("implement me")
}
