package commands

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"main/internal/configuration"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Checks the user out of the room based on the configurations",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := args[0]

		switch action {
		case "up":
			migrateDatabase()
		case "rollback":
			rollbackMigrations()
		default:
			log.Info("action is not supported")
		}

	}}

func migrateDatabase() {
	cfg := configuration.GetConfig()
	db, err := sql.Open("postgres", cfg.DB.ConnectionString())
	if err != nil {
		log.Error(err)
		return
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations/",
		"postgres", driver)
	if err != nil {
		log.Error(err)
		return
	}
	err = m.Up()

	if err != nil {
		log.Error(err)
		return
	}
}

func rollbackMigrations() {
	cfg := configuration.GetConfig()
	db, err := sql.Open("postgres", cfg.DB.ConnectionString())
	if err != nil {
		log.Error(err)
		return
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Error(err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations/",
		"postgres", driver)
	if err != nil {
		log.Error(err)
		return
	}
	err = m.Down()

	if err != nil {
		log.Error(err)
		return
	}
}
