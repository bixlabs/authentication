package cmd

import (
	"bytes"
	"fmt"
	"github.com/bixlabs/authentication/admincli/authentication/structures/createuser"
	"github.com/bixlabs/authentication/admincli/authentication/structures/updateuser"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"testing"
)

const emailWithoutUser = "nouser@email.com"
const validEmail = "test@email.com"
const invalidEmail = "invalid_email"
const validPassword = "very_strong_and_secure_password"
const invalidPassword = "weak"

const oneArgumentErrorMessage = "accepts 1 arg(s), received 0"
const emailRequiredArgumentErrorMessage = "required flag(s) \"email\" not set"
const invalidEmailErrorMessage = "Email is not valid"
const notFoundEmailErrorMessage = "User provided was not found"
const invalidPasswordErrorMessage = "Password should have at least 8 characters"

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
			_, err := executeCommand(findUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(findUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(findUserCommandUse, emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should return a valid user", func() {
			output, err := executeCommand(findUserCommandUse, validEmail)
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
			_, err := executeCommand(deleteUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(deleteUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(deleteUserCommandUse, emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should delete an existing user", func() {
			output, err := executeCommand(deleteUserCommandUse, validEmail)
			Expect(err).To(BeNil())
			Expect(output).To(Equal(fmt.Sprintf("user with email %s was deleted", validEmail)))
		})

		g.It("Should return an error when the user was already deleted", func() {
			_, _ = executeCommand(deleteUserCommandUse, validEmail)
			_, err := executeCommand(deleteUserCommandUse, validEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})
	})

	g.Describe("Create user command", func() {
		const createUserCommandUse = "create-user"

		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			rootCmd.setAuth(auth)

			// reset the create attributes, otherwise the flags are kept before each test
			CreateAttrs = createuser.Command{}
		})

		g.It("Should return an error when email argument is not provided", func() {
			_, err := executeCommand(createUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(emailRequiredArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(createUserCommandUse, "--email", invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the password is invalid", func() {
			_, err := executeCommand(createUserCommandUse, "--email", validEmail, "--password", invalidPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidPasswordErrorMessage))
		})

		g.It("Should create a user", func() {
			output, err := executeCommand(createUserCommandUse, "--email", validEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", validEmail)))
			Expect(output).Should(ContainSubstring("Password"))
		})

		g.It("Should create a user with password", func() {
			output, err := executeCommand(createUserCommandUse, "--email", validEmail, "--password", validPassword)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring("Password:"))
		})

		g.It("Should create a user with given name", func() {
			givenName := "Isabella"
			output, err := executeCommand(createUserCommandUse, "--email", validEmail, "--given-name", givenName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("GivenName:%s", givenName)))
		})

		g.It("Should create a user with second name", func() {
			secondName := "Rose"
			output, err := executeCommand(createUserCommandUse, "--email", validEmail, "--second-name", secondName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondName:%s", secondName)))
		})

		g.It("Should create a user with family name", func() {
			familyName := "Foreman"
			output, err := executeCommand(createUserCommandUse, "--email", validEmail, "--family-name", familyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("FamilyName:%s", familyName)))
		})

		g.It("Should create a user with second family name", func() {
			secondFamilyName := "Barclay"
			output, err := executeCommand(createUserCommandUse, "--email", validEmail, "--second-family-name", secondFamilyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondFamilyName:%s", secondFamilyName)))
		})
	})

	g.Describe("Update user command", func() {
		const updateUserCommandUse = "update-user"

		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			rootCmd.setAuth(auth)

			// reset the update attributes, otherwise the flags are kept before each test
			UpdateAttrs = updateuser.Command{}

			_, err := auth.Create(structures.User{Email: validEmail})

			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error when email argument is not provided", func() {
			_, err := executeCommand(updateUserCommandUse)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(updateUserCommandUse, invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(updateUserCommandUse, emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should return an error when the update email is invalid", func() {
			_, err := executeCommand(updateUserCommandUse, validEmail, "--new-email", invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the update password is invalid", func() {
			_, err := executeCommand(updateUserCommandUse, validEmail, "--password", invalidPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidPasswordErrorMessage))
		})

		g.It("Should update a user", func() {
			output, err := executeCommand(updateUserCommandUse, validEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", validEmail)))
		})

		g.It("Should update a user with email", func() {
			newEmail := "update_email@gmail.com"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--new-email", newEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", newEmail)))
		})

		g.It("Should update a user with password", func() {
			password := "very_strong_and_secure_password"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--password", password)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring("Password:"))
		})

		g.It("Should update a user with given name", func() {
			givenName := "Isabella"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--given-name", givenName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("GivenName:%s", givenName)))
		})

		g.It("Should update a user with second name", func() {
			secondName := "Rose"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--second-name", secondName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondName:%s", secondName)))
		})

		g.It("Should update a user with family name", func() {
			familyName := "Foreman"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--family-name", familyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("FamilyName:%s", familyName)))
		})

		g.It("Should update a user with second family name", func() {
			secondFamilyName := "Barclay"
			output, err := executeCommand(updateUserCommandUse, validEmail, "--second-family-name", secondFamilyName)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("SecondFamilyName:%s", secondFamilyName)))
		})
	})

	g.Describe("Reset password command", func() {
		const resetPasswordCommandUse = "reset-password"

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
			_, err := executeCommand(resetPasswordCommandUse, "--new-password", validPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(oneArgumentErrorMessage))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(resetPasswordCommandUse, invalidEmail, "--new-password", validPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidEmailErrorMessage))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(resetPasswordCommandUse, emailWithoutUser, "--new-password", validPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(notFoundEmailErrorMessage))
		})

		g.It("Should return an error when the reset password is invalid", func() {
			_, err := executeCommand(resetPasswordCommandUse, emailWithoutUser, "--new-password", invalidPassword)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal(invalidPasswordErrorMessage))
		})

		g.It("Should reset a valid user", func() {
			output, err := executeCommand(resetPasswordCommandUse, validEmail, "--new-password", validPassword)
			Expect(err).To(BeNil())
			Expect(output).Should(Equal(fmt.Sprintf("user with email %s was reset", validEmail)))
		})
	})
}

func executeCommand(args ...string) (output string, err error) {
	root := rootCmd

	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetArgs(args)

	_, err = root.ExecuteC()

	return buf.String(), err
}
