package cmd

import (
	"github.com/bixlabs/authentication/admincli/usermanager/structures/mappers"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/resetpassword"

	"github.com/spf13/cobra"
)

var resetPassword resetpassword.Command

// resetPasswordCmd represents the reset-password command
var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password <email>",
	Short: "Reset a user password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		user, err := rootCmd.userManager.Update(email, mappers.ResetUserCommandToUpdateUser(resetPassword))
		if err != nil {
			return err
		}

		cmd.Printf("user with email %s was reset", user.Email)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
	resetPasswordCmd.Flags().StringVar(&resetPassword.Password, "new-password", "", "The new password that will be reset")

	err := resetPasswordCmd.MarkFlagRequired("new-password")
	if err != nil {
		panic(err)
	}
}
