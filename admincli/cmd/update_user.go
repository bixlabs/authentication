package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"

	"github.com/spf13/cobra"
)

// updateUserCmd represents the update-user command
var updateUserCmd = &cobra.Command{
	Use:     "update-user <email>",
	Aliases: []string{"update"},
	Short:   "Update a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentEmail := args[0]
		// TODO: add a mapper
		user, err := rootCmd.Authenticator.Update(currentEmail, structures.UserUpdate{
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
	rootCmd.Command.AddCommand(updateUserCmd)
	updateUserCmd.Flags().StringVar(&email, "new-email", "", "new email")
	updateUserCmd.Flags().StringVar(&password, "password", "", "password")
	updateUserCmd.Flags().StringVar(&givenName, "given-name", "", "given name")
	updateUserCmd.Flags().StringVar(&secondName, "second-name", "", "second name")
	updateUserCmd.Flags().StringVar(&familyName, "family-name", "", "family name")
	updateUserCmd.Flags().StringVar(&secondFamilyName, "second-family-name", "", "second family name")
}
