package authentication

import (
	"github.com/bixlabs/authentication/api/authentication/structures/changepass"
	"github.com/bixlabs/authentication/api/authentication/structures/signup"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/memory"
	email2 "github.com/bixlabs/authentication/email"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"testing"
)

const validEmail = "test@email.com"
const invalidEmail = "invalid_email"
const validPassword = "secured_password"
const invalidPassword = "07chars"

func TestAuthenticatorRest(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator
	var passwordManager interactors.PasswordManager
	var sender email.Sender

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Login rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), email2.NewDummySender()
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
			userRepo, sender := memory.NewUserRepo(), email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
			passwordManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("should return 400 if email is invalid", func() {
			code, _ := startForgotPasswordHandler(invalidEmail, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 500 if email doesn't exist", func() {
			code, _ := startForgotPasswordHandler(validEmail, passwordManager)
			Expect(code).To(Equal(http.StatusInternalServerError))
		})

		g.It("should return 202 if everything goes well", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := startForgotPasswordHandler(validEmail, passwordManager)
			Expect(code).To(Equal(http.StatusAccepted))
		})
	})

	g.Describe("Change Password process", func() {
		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
			passwordManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("Should return 400 if email is not valid", func() {
			request := changepass.Request{Email: invalidEmail}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("Should return 400 if password length is less than 8", func() {
			request := changepass.Request{Email: validEmail, NewPassword: invalidPassword}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("Should return 500 if we can't get the hashed password from db", func() {
			request := changepass.Request{Email: validEmail, NewPassword: validPassword}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusInternalServerError))
		})

		g.It("Should return 400 if user password is not valid", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			request := changepass.Request{Email: user.Email, NewPassword: validPassword}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusUnauthorized))
		})

		g.It("should return 400 if newPassword is the same as the actual one", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			request := changepass.Request{Email: user.Email, OldPassword: validPassword, NewPassword: validPassword}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("Should return 200 if user provides the correct information", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			request := changepass.Request{Email: user.Email, OldPassword: validPassword, NewPassword: "12345678"}
			code, _ := changePasswordHandler(request, passwordManager)
			Expect(code).To(Equal(http.StatusOK))
		})
	})

	g.Describe("Sign up rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender := memory.NewUserRepo(), email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
		})

		g.It("should return 400 if email is invalid", func() {
			request := signup.Request{Email: invalidEmail, Password: validPassword}
			code, _ := signupHandler(request, auth)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if password is invalid", func() {
			request := signup.Request{Email: validEmail, Password: invalidPassword}
			code, _ := signupHandler(request, auth)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if email is duplicated", func() {
			request := signup.Request{Email: validEmail, Password: validPassword}
			_, _ = signupHandler(request, auth)
			code, _ := signupHandler(request, auth)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 201 if user is created successfully", func() {
			request := signup.Request{Email: validEmail, Password: validPassword}
			code, _ := signupHandler(request, auth)
			Expect(code).To(Equal(http.StatusCreated))
		})
	})

	g.Describe("Reset password rest handler", func() {
		g.BeforeEach(func() {
			userRepo := memory.NewUserRepo()
			sender = email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
			passwordManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("should return 400 if email is invalid", func() {
			code, _ := finishResetPasswordHandler(invalidEmail, "4000", validPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if password length is not correct", func() {
			code, _ := finishResetPasswordHandler(validEmail, "4000", invalidPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if reset token is invalid", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			_, _ = startForgotPasswordHandler(validEmail, passwordManager)
			code, _ := finishResetPasswordHandler(validEmail, "23423423424", validPassword, passwordManager)
			Expect(code).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 400 if newPassword is the same as the actual one", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := passwordManager.StartResetPassword(validEmail)
			httpCode, _ := finishResetPasswordHandler(validEmail, code, validPassword, passwordManager)
			Expect(httpCode).To(Equal(http.StatusBadRequest))
		})

		g.It("should return 204 if password is changed successfully", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := passwordManager.StartResetPassword(validEmail)
			httpCode, _ := finishResetPasswordHandler(validEmail, code, "secured_password2", passwordManager)
			Expect(httpCode).To(Equal(http.StatusNoContent))
		})

	})

	g.Describe("Verify JWT rest handler", func() {
		g.BeforeEach(func() {
			secret := "test"
			err := os.Setenv("AUTH_SERVER_SECRET", secret)
			Expect(err).To(BeNil())
			userRepo, sender := memory.NewUserRepo(), email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
		})

		g.It("should return 401 if the token is invalid", func() {
			code, _ := verifyJWTHandler("invalid_token", auth)
			Expect(code).To(Equal(http.StatusUnauthorized))
		})

		g.It("should return 200 if the token is valid", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			Expect(err).To(BeNil())
			token, err := auth.Login(validEmail, validPassword)
			Expect(err).To(BeNil())
			code, result := verifyJWTHandler(token.Token, auth)
			Expect(code).To(Equal(http.StatusOK))
			Expect(result.Result.User.Email).To(Equal(validEmail))
		})
	})
}
