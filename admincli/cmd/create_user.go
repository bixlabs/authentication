package cmd

import (
	"encoding/json"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/createuser"
	"github.com/bixlabs/authentication/admincli/usermanager/structures/mappers"
	"github.com/spf13/cobra"
)

// CreateAttrs represents the params values received by this command
var CreateAttrs createuser.Command

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:     "create-user <email>",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := rootCmd.userManager.Create(mappers.CreateUserCommandToUser(CreateAttrs))

		if err != nil {
			return err
		}

		jsonUser, err := json.Marshal(mappers.UserToCreateUserResult(user))
		if err != nil {
			return err
		}

		cmd.Print(string(jsonUser))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
	createUserCmd.Flags().StringVar(&CreateAttrs.Email, "email", "", "email")
	createUserCmd.Flags().StringVar(&CreateAttrs.Password, "password", "", "password")
	createUserCmd.Flags().StringVar(&CreateAttrs.GivenName, "given-name", "", "given name")
	createUserCmd.Flags().StringVar(&CreateAttrs.SecondName, "second-name", "", "second name")
	createUserCmd.Flags().StringVar(&CreateAttrs.FamilyName, "family-name", "", "family name")
	createUserCmd.Flags().StringVar(&CreateAttrs.SecondFamilyName, "second-family-name", "", "second family name")

	err := createUserCmd.MarkFlagRequired("email")
	if err != nil {
		panic(err)
	}
}
