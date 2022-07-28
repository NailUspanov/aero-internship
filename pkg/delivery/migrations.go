package delivery

import (
	"aero-internship/pkg/config"
	"github.com/golang-migrate/migrate/v4"
	pg_migrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

func MigrateUp(db *sqlx.DB, cfg *config.Config) error {

	driver, err := pg_migrate.WithInstance(db.DB, &pg_migrate.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(cfg.GetMigrationPath(), cfg.GetDBName(), driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}
