package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	db_host     string
	db_port     string
	db_name     string
	db_user     string
	db_password string
	db_ssl      string

	grpc_host string
	grpc_port string

	rest_host string
	rest_port string

	main_dir    string
	jwt_secret  string
	jwt_ttl     string
	refresh_ttl string

	minio_endpoint          string
	minio_access_key_id     string
	minio_secret_access_key string

	migrations_path string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		db_host:                 os.Getenv("DB_HOST"),
		db_port:                 os.Getenv("DB_PORT"),
		db_name:                 os.Getenv("DB_NAME"),
		db_user:                 os.Getenv("DB_USERNAME"),
		db_password:             os.Getenv("DB_PASSWORD"),
		db_ssl:                  os.Getenv("DB_SSLMODE"),
		grpc_host:               os.Getenv("GRPC_HOST"),
		grpc_port:               os.Getenv("GRPC_PORT"),
		rest_host:               os.Getenv("REST_HOST"),
		rest_port:               os.Getenv("REST_PORT"),
		main_dir:                os.Getenv("MAIN_DIR"),
		jwt_secret:              os.Getenv("JWT_SECRET"),
		jwt_ttl:                 os.Getenv("JWT_TTL"),
		refresh_ttl:             os.Getenv("REFRESH_TTL"),
		minio_endpoint:          os.Getenv("MINIO_ENDPOINT"),
		minio_secret_access_key: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		minio_access_key_id:     os.Getenv("MINIO_ACCESS_KEY_ID"),
		migrations_path:         os.Getenv("MIGRATIONS_PATH"),
	}
}

func (cfg *Config) GetMigrationPath() string {
	return cfg.migrations_path
}

func (cfg *Config) GetRefreshTTL() string {
	return cfg.refresh_ttl
}

func (cfg *Config) GetJWTSecret() string {
	return cfg.jwt_secret
}

func (cfg *Config) GetJWTttl() string {
	return cfg.jwt_ttl
}

func (cfg *Config) GetMinioEndpoint() string {
	return cfg.minio_endpoint
}

func (cfg *Config) GetMinioSecretAccessKey() string {
	return cfg.minio_secret_access_key
}

func (cfg *Config) GetMinioAccessKeyId() string {
	return cfg.minio_access_key_id
}

func (cfg *Config) GetDBHost() string {
	return cfg.db_host
}

func (cfg *Config) GetDBPort() string {
	return cfg.db_port
}

func (cfg *Config) GetDBName() string {
	return cfg.db_name
}

func (cfg *Config) GetDBUsername() string {
	return cfg.db_user
}

func (cfg *Config) GetDBPassword() string {
	return cfg.db_password
}

func (cfg *Config) GetDBSSLmode() string {
	return cfg.db_ssl
}
func (cfg *Config) GetGRPCHost() string {
	return cfg.grpc_host
}

func (cfg *Config) GetGRPCPort() string {
	return cfg.grpc_port
}

func (cfg *Config) GetRESTHost() string {
	return cfg.rest_host
}

func (cfg *Config) GetRESTPort() string {
	return cfg.rest_port
}

func (cfg *Config) GetMainDir() string {
	return cfg.main_dir
}
