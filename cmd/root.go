package cmd

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	logger  log.Logger

	rootCmd = &cobra.Command{
		Use:   "berichtsheft-cli",
		Short: "A Berichtsheft is a booklet you need to do in Germany due to the regulations of IHK",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	logger = log.NewLogfmtLogger(os.Stdout)
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.berichtsheft/config.yaml)")
	viper.SetDefault("author", "Tom Siewert <tom@siewert.io>")
	viper.SetDefault("license", "apache")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigName(".berichtsheft")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig; err == nil {
		level.Debug(logger).Log("Using config file:", viper.ConfigFileUsed())
	}
}
