package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"

	"github.com/spf13/cobra"
)

type UpdateUser struct {
	email            string
	password         string
	givenName        string
	secondName       string
	familyName       string
	secondFamilyName string
}

var UpdateAttrs UpdateUser

// updateUserCmd represents the update-user command
var updateUserCmd = &cobra.Command{
	Use:     "update-user <email>",
	Aliases: []string{"update"},
	Short:   "Update a user",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		currentEmail := args[0]

		// TODO: add a mapper
		userUpdate := structures.UserUpdate{
			Email:            UpdateAttrs.email,
			Password:         UpdateAttrs.password,
			GivenName:        UpdateAttrs.givenName,
			SecondName:       UpdateAttrs.secondName,
			FamilyName:       UpdateAttrs.familyName,
			SecondFamilyName: UpdateAttrs.secondFamilyName,
		}

		user, err := rootCmd.Authenticator.Update(currentEmail, userUpdate)

		if err != nil {
			return err
		}

		cmd.Printf("%+v\n", user)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateUserCmd)

	updateUserCmd.Flags().StringVar(&UpdateAttrs.email, "new-email", "", "new email")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.password, "password", "", "password")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.givenName, "given-name", "", "given name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.secondName, "second-name", "", "second name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.familyName, "family-name", "", "family name")
	updateUserCmd.Flags().StringVar(&UpdateAttrs.secondFamilyName, "second-family-name", "", "second family name")
}
