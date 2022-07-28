package usecase

import (
	"aero-internship/gen/api"
	"aero-internship/internal/adapters"
	"aero-internship/internal/domain/entity/files"
	"aero-internship/internal/domain/entity/tokens"
	"aero-internship/internal/domain/entity/users"
	"aero-internship/internal/domain/usecase/auth_usecase"
	"aero-internship/internal/domain/usecase/files_usecase"
	"aero-internship/internal/domain/usecase/news_usecase"
	"aero-internship/pkg/config"
	"context"
	"google.golang.org/grpc"
)

type AuthService interface {
	GenerateTokens(tknDTO *tokens.TokenDTO) (*auth_usecase.Tokens, error)
	GenerateTokensFromUserDTO(userDTO *users.UserDTO) (*auth_usecase.Tokens, error)
	ParseToken(tkn string) (auth_usecase.TokenDTO, error)
	RegisterUser(userDTO *users.UserDTO) (*auth_usecase.Tokens, error)
	SignIn(signInDTO *auth_usecase.SignInDTO) (*auth_usecase.Tokens, error)
	UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

type FileService interface {
	GetFile(ctx context.Context, bucketName, fileName string) (*api.File, error)
	GetFilesByNewsID(ctx context.Context, newsID string) ([]*api.File, error)
	CreateFile(ctx context.Context, newsID string, file files.FileDTO) error
	DeleteFile(ctx context.Context, newsID, fileId string) error
}

type NewsService interface {
}

type Service struct {
	AuthService
	NewsService
	FileService
}

func NewService(cfg *config.Config, dataTransfer adapters.DataTransfer) *Service {
	return &Service{
		AuthService: auth_usecase.NewAuthService(cfg, dataTransfer),
		NewsService: news_usecase.NewNewsService(cfg, dataTransfer),
		FileService: files_usecase.NewFileService(cfg, dataTransfer)}
}
