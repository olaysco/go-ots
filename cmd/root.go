package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ots",
	Short: "A simple command to generate one-time secret.",
	Long: `Generate a onetime secret for securely sharing your password/secrets
		    using the one-time secret service API https://onetimesecret.com/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var cfgFile string

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitCredentials(&cfgFile))
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./ots.yaml)")
}
