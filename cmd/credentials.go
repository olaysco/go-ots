package cmd

import (
	"os"

	"github.com/olaysco/ots/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitCredentials(cfgFile *string) func() {
	return initCredentials
}

func initCredentials() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Use config file from current working directory
		home, err := os.Getwd()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("ots")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	cobra.CheckErr(err)
}

func GetCredentials() *server.Credentials {
	return &server.Credentials{
		Username: viper.GetString("username"),
		Password: viper.GetString("password"),
	}
}
