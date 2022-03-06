package cmd

import (
	"fmt"

	"github.com/olaysco/ots/server"
	"github.com/spf13/cobra"
)

func statusCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "status",
		Short: "Current system status",
		Long: `Get the current status of the one-time service
					The current system status. One of: nominal, offline.`,
		Run: func(cmd *cobra.Command, args []string) {
			statusRequest := server.StatusRequest{
				Credentials: *GetCredentials(),
			}
			if response, err := server.Status(statusRequest); err != nil {
				fmt.Printf("an error occured %s", err.Error())
			} else {
				fmt.Printf("status: %s", response.Status)
			}
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(statusCmd())
}
