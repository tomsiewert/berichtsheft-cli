package database

import (
	"log"
	"os"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/tomsdevsn/berichtsheft-cli/internal/tools"
)

func InitDatabase(path *string, force bool) bool {
	if _, err := os.Stat(tools.StringPointer(path)); err == nil {
		if force {
			if !forceConfirmation() {
				os.Exit(0)
			}
		} else {
			log.Fatal("A file already exists at this path.")
		}
	}
	if err := createDatabaseFile(path); err != nil {
		log.Fatal(err)
	}
	if err := runDatabaseMigration(path); err != nil {
		log.Fatal(err)
	}
	return false
}

func createDatabaseFile(path *string) error {
	file, err := os.Create(tools.StringPointer(path))
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	file.Close()
	log.Println("Database file created")
	return nil
}

func runDatabaseMigration(path *string) error {
	m, err := migrate.New(
		"github://tomsdevsn/berichtsheft-cli/db/migrations",
		"sqlite3://"+tools.StringPointer(path)+"?query")

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	return err
}

func forceConfirmation() bool {
	confirmation := tools.AskForConfirmation("A file already exists at this path and force parameter is set. Do you want to continue?", false)
	return confirmation
}
