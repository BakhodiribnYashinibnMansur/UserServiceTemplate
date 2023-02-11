package repository

import (
	"database/sql"
	"fmt"
	"users_service/config"
	"users_service/util/logrus_log"

	migrate "github.com/rubenv/sql-migrate"
)

func psqlInitURL(cfg *config.Config) string {

	// URL for Migration
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	)
	return url
}

func MigratePsql(cfg *config.Config, logrus *logrus_log.Logger, up bool) error {

	// OR: Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: "./schema",
	}

	db, err := sql.Open("postgres", psqlInitURL(cfg))
	if err != nil {
		logrus.Fatal("error in creating migrations: ", err.Error())
	}
	defer db.Close()
	migrateState := migrate.Up
	if !up {
		migrateState = migrate.Down
	}
	n, err := migrate.Exec(db, "postgres", migrations, migrateState)
	if err != nil {
		logrus.Fatal("error in creating migrations: ", err.Error())
	}

	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
