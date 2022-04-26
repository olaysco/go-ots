package cmd

import (
	"fmt"

	"github.com/olaysco/ots/server"
	"github.com/olaysco/ots/utils"
	"github.com/spf13/cobra"
)

const flagSecretKey = "secret_key"

func retrieveCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "retrieve",
		Short: "Get/Retrieve a secret",
		Long:  `Get/Retrieve a previously created secret.`,
		Run: func(cmd *cobra.Command, args []string) {
			retrieveRequest := server.RetrieveRequest{
				Credentials: *GetCredentials(),
			}

			retrieveRequest.Secret, _ = cmd.Flags().GetString(flagSecretKey)
			retrieveRequest.Passphrase, _ = cmd.Flags().GetString(flagPassphrase)

			response, err := server.Retrieve(retrieveRequest)
			utils.HandleError(err)

			fmt.Println("Secret retrieved successfully.")
			fmt.Println("______________________________________________________________________________________")
			fmt.Printf("%-15s : %s \n", "Secret Key", response.Secret)
			fmt.Printf("%-15s : %s \n", "Shared secret", response.Value)
			fmt.Println("______________________________________________________________________________________")
		},
	}

	cmd.Flags().String(flagSecretKey, "", "Secret key")
	cmd.Flags().String(flagPassphrase, "", "Key that the recipient must know to view the secret")

	cmd.MarkFlagRequired(flagSecretKey)

	return cmd
}

func init() {
	RootCmd.AddCommand(retrieveCmd())
}
