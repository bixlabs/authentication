package cmd

import (
	"bytes"
	"fmt"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"testing"
)

const emailWithoutUser = "nouser@email.com"
const validEmail = "test@email.com"
const invalidEmail = "invalid_email"

const oneArgumentErrorMessage = "accepts 1 arg(s), received 0"
const invalidEmailErrorMessage = "Email is not valid"
const notFoundEmailErrorMessage = "User provided was not found"

func TestAdminCli(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	var auth interactors.Authenticator

	g.Describe("Find user command", func() {
		const findUserCommandUse = "find-user"

		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			rootCmd.setAuth(auth)
			_, err := auth.Create(structures.User{Email: validEmail})

			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error when email argument is not provided", func() {
			_, err := executeCommand(&rootCmd.Command, findUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(&rootCmd.Command, findUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(&rootCmd.Command, findUserCommandUse, emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should return a valid user", func() {
			output, err := executeCommand(&rootCmd.Command, findUserCommandUse, validEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", validEmail)))
		})
	})

	g.Describe("Delete user command", func() {
		const deleteUserCommandUse = "delete-user"

		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			rootCmd.setAuth(auth)
			_, err := auth.Create(structures.User{Email: validEmail})

			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error when email argument is not provided", func() {
			_, err := executeCommand(&rootCmd.Command, deleteUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(&rootCmd.Command, deleteUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(&rootCmd.Command, deleteUserCommandUse, emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should delete an existing user", func() {
			output, err := executeCommand(&rootCmd.Command, deleteUserCommandUse, validEmail)
			Expect(err).To(BeNil())
			Expect(output).To(Equal(fmt.Sprintf("user with email %s was deleted", validEmail)))
		})
	})

	g.Describe("Create user command", func() {
		const createUserCommandUse = "create-user"

		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			rootCmd.setAuth(auth)
		})

		g.It("Should return an error when email argument is not provided", func() {
			_, err := executeCommand(&rootCmd.Command, createUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(&rootCmd.Command, createUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should create a user", func() {
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", validEmail)))
			Expect(output).Should(ContainSubstring("Password"))
		})

		g.It("Should create a user with password", func() {
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail, "--second-name", secondName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring("Password:"))
		})

		g.It("Should create a user with given name", func() {
			givenName := "Isabella"
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail, "--given-name", givenName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("GivenName:%s", givenName)))
		})

		g.It("Should create a user with second name", func() {
			secondName := "Rose"
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail, "--second-name", secondName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondName:%s", secondName)))
		})

		g.It("Should create a user with family name", func() {
			familyName := "Foreman"
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail, "--family-name", familyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("FamilyName:%s", familyName)))
		})

		g.It("Should create a user with second family name", func() {
			secondFamilyName := "Barclay"
			output, err := executeCommand(&rootCmd.Command, createUserCommandUse, validEmail, "--second-family-name", secondFamilyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondFamilyName:%s", secondFamilyName)))
		})
	})
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}
