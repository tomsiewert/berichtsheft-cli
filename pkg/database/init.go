package database

import (
	"log"
	"os"
	"strconv"

	"github.com/tomsdevsn/berichtsheft-cli/pkg/tools"
)

func InitDatabase(path string, force bool) {
	log.Println("Ensure if database already exists")
	log.Println(path)

	if _, err := os.Stat(path); err == nil {
		if !force {
			confirmation := tools.AskForConfirmation("The database path already exists and force parameter is not set. Do you want to continue?", false)
			log.Println(strconv.FormatBool(confirmation))
		}
	}
}
