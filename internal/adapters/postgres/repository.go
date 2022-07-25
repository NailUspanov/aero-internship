package postgres

import (
	"aero-internship/internal/adapters/postgres/refresh_sessions"
	users2 "aero-internship/internal/adapters/postgres/users"
	"aero-internship/internal/domain/entity/tokens"
	"aero-internship/internal/domain/entity/users"
	"aero-internship/pkg/config"
	"github.com/jmoiron/sqlx"
)

type RefreshSessions interface {
	MakeNewSession(cfg *config.Config, userId int, refreshToken string) error
}

type Users interface {
	GetUserDTObyEmail(cfg *config.Config, email string) (*users.User, error)
	GetTokenDTOFromUserDTO(cfg *config.Config, userDTO *users.UserDTO) (*tokens.TokenDTO, error)
	MakeRegistrationTxn(cfg *config.Config, userDTO users.UserDTO) error
}

type Repository struct {
	db  *sqlx.DB
	cfg *config.Config
	RefreshSessions
	Users
}

func NewRepository(db *sqlx.DB, cfg *config.Config) *Repository {
	return &Repository{
		RefreshSessions: refresh_sessions.NewRefreshSessionsRepository(cfg, db),
		Users:           users2.NewUsers(cfg, db),
	}
}
