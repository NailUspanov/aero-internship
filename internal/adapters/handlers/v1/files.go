package v1

import (
	"aero-internship/internal/domain/usecase"
	"aero-internship/pkg/config"
	"net/http"
)

type FilesHandler struct {
	service usecase.Service
	cfg     *config.Config
}

func NewFilesHandler(service usecase.Service, cfg *config.Config) *FilesHandler {
	return &FilesHandler{service: service, cfg: cfg}
}

func (f FilesHandler) Create(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

}
