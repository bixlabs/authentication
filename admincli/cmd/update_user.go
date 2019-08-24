package cmd

import (
	"encoding/json"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/mappers"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/updateuser"
	"github.com/spf13/cobra"
)

// UpdateAttrs represents the params values received by this command
var UpdateAttrs updateuser.Command

// updateUserCmd represents the update-user command
var updateUserCmd = &cobra.Command{
	Use:     "update-user <email>",
	Aliases: []string{"update"},
	Short:   "Update a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentEmail := args[0]
		user, err := rootCmd.userManager.Update(currentEmail, mappers.UpdateUserCommandToUpdateUser(UpdateAttrs))

		if err != nil {
			return err
		}

		jsonUser, err := json.Marshal(mappers.UserToUpdateUserResult(user))
		if err != nil {
			return err
		}

		cmd.Print(string(jsonUser))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateUserCmd)

	updateUserCmd.Flags().StringVar(&UpdateAttrs.Email, "new-email", "", "new email")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.Password, "password", "", "password")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.GivenName, "given-name", "", "given name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.SecondName, "second-name", "", "second name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.FamilyName, "family-name", "", "family name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.SecondFamilyName, "second-family-name", "", "second family name")
}
