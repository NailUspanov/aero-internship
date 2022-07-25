package handlers

import (
	"aero-internship/gen/api"
	v1 "aero-internship/internal/adapters/handlers/v1"
	"aero-internship/internal/domain/usecase"
	"aero-internship/pkg/config"
	"net/http"
)

type AuthHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request, pathParams map[string]string)
	SignIn(w http.ResponseWriter, r *http.Request, pathParams map[string]string)
}

type Handler struct {
	api.ContentCheckServiceServer
	Cfg *config.Config
	AuthHandler
}

func NewHandler(cfg *config.Config, service *usecase.Service) *Handler {
	return &Handler{
		Cfg:                       cfg,
		ContentCheckServiceServer: v1.NewHealthCheckHandler(),
		AuthHandler:               v1.NewAuthHandler(cfg, *service),
	}
}
