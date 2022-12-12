package repository

import (
	"fmt"
	"strings"
	"users_service/util/logrus_log"

	"users_service/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"          //database is needed for migration
	_ "github.com/golang-migrate/migrate/v4/database/postgres" //postgres is used for database
	_ "github.com/golang-migrate/migrate/v4/source/file"       //file is needed for migration url
)

func migrateInit() string {
	var conf = config.Get()

	// URL for Migration
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.PostgresUser,
		conf.PostgresPassword,
		conf.PostgresHost,
		conf.PostgresPort,
		conf.PostgresDatabase,
	)
	return url
}

// MigrateUP ...
func MigrateUP(logrus *logrus_log.Logger) {
	url := migrateInit()
	m, err := migrate.New("file://schema", url)
	if err != nil {
		logrus.Fatal("error in creating migrations: ", err.Error())
	}
	if err := m.Up(); err != nil {
		if !strings.Contains(err.Error(), "no change") {
			logrus.Error("Error in migrating ", err.Error())
		}
	}
}
