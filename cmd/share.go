package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func shareCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "share",
		Short: "Create a one time secret to share",
		Long:  `Create a one time secret URL that can be shared with others`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Share command")
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(shareCmd())
}
