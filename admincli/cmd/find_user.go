package cmd

import (
	"encoding/json"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/mappers"
	"github.com/spf13/cobra"
)

var findUserCmd = &cobra.Command{
	Use:     "find-user <email>",
	Aliases: []string{"find"},
	Short:   "Find a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
		user, err := rootCmd.userManager.Find(email)
		if err != nil {
			return err
		}

		jsonUser, err := json.Marshal(mappers.UserToFindUserResult(user))
		if err != nil {
			return err
		}

		cmd.Print(string(jsonUser))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(findUserCmd)
}
