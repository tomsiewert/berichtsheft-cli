/*
Copyright © 2021 Tom Siewert

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/tomsdevsn/berichtsheft-cli/internal/database"
)

var (
	force        bool
	databasePath string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the database for the Berichtsheft",
	Run: func(cmd *cobra.Command, args []string) {
		database.InitDatabase(&databasePath, force)
	},
}

func init() {
	defaultDatabasePath, _ := homedir.Expand("~/.berichtsheft/berichtsheft.db")
	initCmd.Flags().StringVarP(&databasePath, "database.path", "", defaultDatabasePath, "Path to the database file")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force the initialization of the database (will overwrite if it already exists)")
	RootCmd.AddCommand(initCmd)
}
