package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/spf13/cobra"
)

var email, password, givenName, secondName, familyName, secondFamilyName string

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:     "create-user <email>",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: add a mapper
		user, err := rootCmd.Authenticator.Create(structures.User{
			Email:            email,
			Password:         password,
			GivenName:        givenName,
			SecondName:       secondName,
			FamilyName:       familyName,
			SecondFamilyName: secondFamilyName,
		})

		if err != nil {
			return err
		}

		cmd.Printf("%+v\n", user)
		return nil
	},
}

func init() {
	rootCmd.Command.AddCommand(createUserCmd)
	createUserCmd.Flags().StringVar(&email, "email", "", "email")
	createUserCmd.Flags().StringVar(&password, "password", "", "password")
	createUserCmd.Flags().StringVar(&givenName, "given-name", "", "given name")
	createUserCmd.Flags().StringVar(&secondName, "second-name", "", "second name")
	createUserCmd.Flags().StringVar(&familyName, "family-name", "", "family name")
	createUserCmd.Flags().StringVar(&secondFamilyName, "second-family-name", "", "second family name")
}
