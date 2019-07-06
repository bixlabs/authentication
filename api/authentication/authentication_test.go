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
	var passwordManager interactors.PasswordManager

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Login rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
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

	g.Describe("Reset password request rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			passwordManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("should return 400 if email is invalid", func() {
			code, _ := forgotPasswordHandler(invalidEmail, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 500 if email doesn't exist", func() {
			code, _ := forgotPasswordHandler(validEmail, passwordManager)
			Expect(code).To(Equal(http.StatusInternalServerError))
		})

		g.It("should return 202 if everything goes well", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := forgotPasswordHandler(validEmail, passwordManager)
			Expect(code).To(Equal(http.StatusAccepted))
		})
	})

	g.Describe("Reset password rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			auth = implementation.NewAuthenticator(userRepo, sender)
			passwordManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("should return 400 if email is invalid", func() {
			code, _ := resetPasswordHandler(invalidEmail, "4000", validPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if password length is not correct", func() {
			code, _ := resetPasswordHandler(validEmail, "4000", invalidPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if reset token is invalid", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := forgotPasswordHandler(validEmail, passwordManager)
			code, _ = resetPasswordHandler(validEmail, "23423423424", validPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 204 if password is changed successfully", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := passwordManager.ForgotPassword(validEmail)
			httpCode, _ := resetPasswordHandler(validEmail, code, validPassword, passwordManager)
			Expect(httpCode).To(Equal(http.StatusNoContent))
		})
	})
}
