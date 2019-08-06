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

func TestAdminCli(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })
	var auth interactors.Authenticator

	g.Describe("Find user command", func() {
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
			_, err := executeCommand(&rootCmd.Command, "find-user")
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("accepts 1 arg(s), received 0"))
		})

		g.It("Should return an error when the email is invalid", func() {
			_, err := executeCommand(&rootCmd.Command, "find-user", invalidEmail)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("Email is not valid"))
		})

		g.It("Should return an error when the email does not exist", func() {
			_, err := executeCommand(&rootCmd.Command, "find-user", emailWithoutUser)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(Equal("User provided was not found"))
		})

		g.It("Should return a valid user", func() {
			output, err := executeCommand(&rootCmd.Command, "find-user", validEmail)
			Expect(err).To(BeNil())
			Expect(output).Should(ContainSubstring(fmt.Sprintf("Email:%s", validEmail)))
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
