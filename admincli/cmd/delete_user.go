package cmd

import (
	"github.com/spf13/cobra"
)

var deleteUserCmd = &cobra.Command{
	Use:     "delete-user <email>",
	Aliases: []string{"delete"},
	Short:   "Delete a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		err := rootCmd.userManager.Delete(email)

		if err != nil {
			return err
		}

		cmd.Printf("user with email %s was deleted", email)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
