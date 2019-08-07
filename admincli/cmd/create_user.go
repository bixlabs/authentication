package cmd

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/spf13/cobra"
)

type CreateUser struct {
	email            string
	password         string
	givenName        string
	secondName       string
	familyName       string
	secondFamilyName string
}

var CreateAttrs CreateUser

// createUserCmd represents the create-user command
var createUserCmd = &cobra.Command{
	Use:     "create-user <email>",
	Aliases: []string{"create"},
	Short:   "Create a user",
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: add a mapper
		user, err := rootCmd.Authenticator.Create(structures.User{
			Email:            CreateAttrs.email,
			Password:         CreateAttrs.password,
			GivenName:        CreateAttrs.givenName,
			SecondName:       CreateAttrs.secondName,
			FamilyName:       CreateAttrs.familyName,
			SecondFamilyName: CreateAttrs.secondFamilyName,
		})

		if err != nil {
			return err
		}

		cmd.Printf("%+v\n", user)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
	createUserCmd.Flags().StringVar(&CreateAttrs.email, "email", "", "email")
	createUserCmd.Flags().StringVar(&CreateAttrs.password, "password", "", "password")
	createUserCmd.Flags().StringVar(&CreateAttrs.givenName, "given-name", "", "given name")
	createUserCmd.Flags().StringVar(&CreateAttrs.secondName, "second-name", "", "second name")
	createUserCmd.Flags().StringVar(&CreateAttrs.familyName, "family-name", "", "family name")
	createUserCmd.Flags().StringVar(&CreateAttrs.secondFamilyName, "second-family-name", "", "second family name")

	err := createUserCmd.MarkFlagRequired("email")
	if err != nil {
		panic(err)
	}
}
