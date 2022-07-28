package news_usecase

import (
	"aero-internship/internal/adapters"
	"aero-internship/pkg/config"
)

type NewsService struct {
	adapters.DataTransfer
	cfg *config.Config
}

func NewNewsService(cfg *config.Config, dataTransfer adapters.DataTransfer) *NewsService {
	return &NewsService{DataTransfer: dataTransfer, cfg: cfg}
}
