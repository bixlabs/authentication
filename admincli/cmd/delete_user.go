package cmd

import (
	"github.com/spf13/cobra"
)

// deleteUserCmd represents the delete-user command
var deleteUserCmd = &cobra.Command{
	Use:     "delete-user",
	Aliases: []string{"delete"},
	Short:   "Delete a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		err := rootCmd.Authenticator.Delete(email)

		if err != nil {
			return err
		}

		cmd.Printf("user with email %s was deleted", email)
		return nil
	},
}

func init() {
	rootCmd.Command.AddCommand(deleteUserCmd)
}
