package adapters

import (
	minio_storage "aero-internship/internal/adapters/minio"
	"aero-internship/internal/adapters/postgres"
	"aero-internship/pkg/config"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type DataTransfer struct {
	Repo    postgres.Repository
	Storage minio_storage.Storage
	cfg     *config.Config
}

func NewDataTransfer(db *sqlx.DB, cfg *config.Config, minio *minio.Client) *DataTransfer {
	return &DataTransfer{Repo: *postgres.NewRepository(db, cfg), Storage: *minio_storage.NewStorage(minio)}
}
