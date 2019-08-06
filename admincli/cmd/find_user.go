package cmd

import (
	"github.com/spf13/cobra"
)

// findUserCmd represents the find-user command
var findUserCmd = &cobra.Command{
	Use:     "find-user <email>",
	Aliases: []string{"find"},
	Short:   "Find a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		user, err := rootCmd.Authenticator.Find(email)
		if err != nil {
			return err
		}

		cmd.Printf("%+v\n", user)
		return nil
	},
}

func init() {
	rootCmd.Command.AddCommand(findUserCmd)
}
