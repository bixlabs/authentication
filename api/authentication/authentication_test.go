package authentication

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

const validEmail = "test@email.com"
const invalidEmail = "invalid_email"
const validPassword = "secured_password"
const invalidPassword = "07chars"

func TestRest(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Login rest handler", func() {
		g.BeforeEach(func() {
			auth = implementation.NewAuthenticator(in_memory.NewUserRepo(), in_memory.DummySender{})
		})

		g.It("should return 400 if email is invalid", func() {
			code, _ := loginHandler(invalidEmail, validPassword, auth)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if password is invalid", func() {
			code, _ := loginHandler(validEmail, invalidPassword, auth)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 401 if credentials are wrong", func() {
			code, _ := loginHandler(validEmail, validPassword, auth)
			Expect(code).To(Equal(http.StatusUnauthorized))
		})

		g.It("should return 200 if credentials are correct", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, response := loginHandler(validEmail, validPassword, auth)
			Expect(code).To(Equal(http.StatusOK))
			Expect(response.Result.User.Email).To(Equal(validEmail))
		})
	})
}
