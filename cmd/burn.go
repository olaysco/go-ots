package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func burnCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "burn",
		Short: "burn a secret",
		Long: `burn a previously created secret
				Once a secret is burnt it can no logner be retrieved.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("burn command")
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(burnCmd())
}
