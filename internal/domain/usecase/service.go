package usecase

import (
	"aero-internship/internal/adapters/postgres"
	"aero-internship/internal/domain/entity/tokens"
	"aero-internship/internal/domain/entity/users"
	"aero-internship/internal/domain/usecase/auth_usecase"
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

type NewsService interface {
}

type Service struct {
	AuthService
	NewsService
}

func NewService(cfg *config.Config, repository postgres.Repository) *Service {
	return &Service{
		AuthService: auth_usecase.NewAuthService(cfg, repository),
		NewsService: news_usecase.NewNewsService(cfg, repository)}
}
