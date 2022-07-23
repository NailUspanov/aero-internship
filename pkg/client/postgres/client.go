package postgres

import (
	"aero-internship/pkg/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	NewsTable = "news"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.GetDBHost(),
		cfg.GetDBPort(),
		cfg.GetDBUsername(),
		cfg.GetDBName(),
		cfg.GetDBPassword(),
		cfg.GetDBSSLmode(),
	)
	db, err := sqlx.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
