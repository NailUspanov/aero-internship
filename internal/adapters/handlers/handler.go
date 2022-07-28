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

type FilesHandler interface {
	Create(w http.ResponseWriter, r *http.Request, pathParams map[string]string)
}

type Handler struct {
	api.ContentCheckServiceServer
	Cfg *config.Config
	AuthHandler
	FilesHandler
}

func NewHandler(cfg *config.Config, service *usecase.Service) *Handler {
	return &Handler{
		Cfg:                       cfg,
		ContentCheckServiceServer: v1.NewHealthCheckHandler(*service),
		AuthHandler:               v1.NewAuthHandler(cfg, *service),
		FilesHandler:              v1.NewFilesHandler(*service, cfg),
	}
}
