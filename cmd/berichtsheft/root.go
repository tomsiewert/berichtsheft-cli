package cmd

import (
	"os"

	"github.com/go-kit/kit/log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	logger       log.Logger
	BuildDate    string
	BuildVersion string

	RootCmd = &cobra.Command{
		Use:   "berichtsheft-cli",
		Short: "A Berichtsheft is a booklet you need to do in Germany due to the regulations of IHK",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initConfig(cmd)
		},
	}
)

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	logger = log.NewLogfmtLogger(os.Stdout)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.berichtsheft/config.yaml)")
}

func initConfig(cmd *cobra.Command) error {
	v := viper.New()

	if cfgFile != "" {
		v.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		v.AddConfigPath(home + "/.berichtsheft")
		v.SetConfigName("config")
		v.SetConfigType("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.AutomaticEnv()

	return nil
}
