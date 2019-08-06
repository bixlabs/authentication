package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/spf13/cobra"
)

var password, givenName, secondName, familyName, secondFamilyName string

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:     "create-user",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		email := args[0]
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
	createUserCmd.Flags().StringVar(&password, "password", "", "Password")
	createUserCmd.Flags().StringVar(&givenName, "given-name", "", "Given Name")
	createUserCmd.Flags().StringVar(&secondName, "second-name", "", "Second Name")
	createUserCmd.Flags().StringVar(&familyName, "family-name", "", "Family Name")
	createUserCmd.Flags().StringVar(&secondFamilyName, "second-family-name", "", "Second Family Name")
}
