package adapters

import (
	miniostorage "aero-internship/internal/adapters/minio"
	"aero-internship/internal/adapters/postgres"
	rediscache "aero-internship/internal/adapters/redis"
	"aero-internship/pkg/config"
	"github.com/go-redis/redis/v7"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

type DataTransfer struct {
	Repo    postgres.Repository
	Storage miniostorage.Storage
	Cache   rediscache.CacheRedis
	cfg     *config.Config
}

func NewDataTransfer(db *sqlx.DB, cfg *config.Config, minio *minio.Client, redis *redis.Client) *DataTransfer {
	return &DataTransfer{
		Repo:    *postgres.NewRepository(db, cfg),
		Storage: *miniostorage.NewStorage(minio),
		Cache:   *rediscache.NewCacheRedis(redis),
	}
}
