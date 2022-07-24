package refresh_sessions

import (
	"aero-internship/pkg/config"
	"github.com/jmoiron/sqlx"
	"time"
)

type RefreshSessionsRepository struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewRefreshSessionsRepository(cfg *config.Config, db *sqlx.DB) *RefreshSessionsRepository {
	return &RefreshSessionsRepository{cfg: cfg, db: db}
}

func (r *RefreshSessionsRepository) MakeNewSession(cfg *config.Config, userId int, refreshToken string) error {

	//начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	q := `
		insert into RefreshSessions (userId,refreshToken,expiresIn,createdAt) 
		values ($1,$2,$3,$4)
	`

	createdAt := time.Now()
	refresh_ttl, _ := time.ParseDuration(cfg.GetRefreshTTL())

	_, err = tx.Exec(
		q,
		userId,
		refreshToken,
		createdAt.Add(refresh_ttl).Unix(),
		createdAt.Unix(),
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
