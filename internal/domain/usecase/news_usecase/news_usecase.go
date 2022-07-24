package news_usecase

import (
	"aero-internship/internal/adapters/postgres"
	"aero-internship/pkg/config"
)

type NewsService struct {
	postgres.Repository
	cfg *config.Config
}

func NewNewsService(cfg *config.Config, repository postgres.Repository) *NewsService {
	return &NewsService{Repository: repository, cfg: cfg}
}
