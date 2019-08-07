package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"

	"github.com/spf13/cobra"
)

var resetPassword string

// resetPasswordCmd represents the reset-password command
var resetPasswordCmd = &cobra.Command{
	Use:   "reset-password <email>",
	Short: "Reset a user password",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]

		// TODO: we need a mapper here
		user, err := rootCmd.Authenticator.Update(email, structures.UpdateUser{Password: resetPassword})
		if err != nil {
			return err
		}

		cmd.Printf("user with email %s was reset", user.Email)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
	resetPasswordCmd.Flags().StringVar(&resetPassword, "new-password", "", "The new password that will be reset")

	err := resetPasswordCmd.MarkFlagRequired("new-password")
	if err != nil {
		panic(err)
	}
}
