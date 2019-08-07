package cmd

import (
	"github.com/bixlabs/authentication/admincli/authentication/structures/createuser"
	"github.com/bixlabs/authentication/admincli/authentication/structures/mappers"
	"github.com/spf13/cobra"
)

var CreateAttrs createuser.Command

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:     "create-user <email>",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		user, err := rootCmd.Authenticator.Create(mappers.CreateUserCommandToUser(CreateAttrs))

		if err != nil {
			return err
		}

		cmd.Printf("%+v\n", user)
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
