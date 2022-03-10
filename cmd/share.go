package cmd

import (
	"fmt"

	"github.com/olaysco/ots/server"
	"github.com/olaysco/ots/utils"
	"github.com/spf13/cobra"
)

const flagTTL = "ttl"
const flagSecret = "secret"
const flagRecipient = "recipient"
const flagPassphrase = "passphrase"

func shareCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "share",
		Short: "Create a one time secret to share",
		Long:  `Create a one time secret URL that can be shared with others`,
		Run: func(cmd *cobra.Command, args []string) {
			shareRequest := &server.ShareRequest{
				Credentials: *GetCredentials(),
			}

			shareRequest.TTL, _ = cmd.Flags().GetUint(flagTTL)
			shareRequest.Secret, _ = cmd.Flags().GetString(flagSecret)
			shareRequest.Recipient, _ = cmd.Flags().GetString(flagRecipient)
			shareRequest.Passphrase, _ = cmd.Flags().GetString(flagPassphrase)

			response, err := server.Share(*shareRequest)
			utils.HandleError(err)

			fmt.Println("Secret generated successfully.")
			fmt.Println("______________________________________________________________________________________")
			fmt.Printf("%-15s : %s \n", "Secret key", response.SecretKey)
			fmt.Printf("%-15s : %s \n", "Created", utils.UnixToDate(int64(response.Created)))
			fmt.Printf("%-15s : %s \n", "Secret link", response.Link)
			fmt.Printf("%-15s : %d seconds \n", "Secret TTL", response.TTL)
			fmt.Println("______________________________________________________________________________________")
		},
	}

	cmd.Flags().String(flagSecret, "", "Secret to share")
	cmd.Flags().String(flagRecipient, "", "Email address to send the secret link to")
	cmd.Flags().String(flagPassphrase, "", "Key that the recipient must know to view the secret")
	cmd.Flags().Uint(flagTTL, 60, "The maximum amount of time, in seconds, that the secret should survive (default 60)")

	cmd.MarkFlagRequired(flagSecret)

	return cmd
}

func init() {
	RootCmd.AddCommand(shareCmd())
}
