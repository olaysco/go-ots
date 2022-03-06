package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func retrieveCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "",
		Short: "Get/Retrieve a secret",
		Long:  `Get/Retrieve a previously created secret.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Retrieve command")
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(retrieveCmd())
}
