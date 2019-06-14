package implementation

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestPasswordManager(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var passwordManager interactors.PasswordManager
	var auth interactors.Authenticator

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Change password process", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			passwordManager = NewPasswordManager(userRepo, sender)
			auth = NewAuthenticator(userRepo, sender)
		})

		g.It("Should return an error if invalid email", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			user.Password = "anotherPassword"
			user.Email = invalidEmail
			err = passwordManager.ChangePassword(user, "Asdqwe123")
			Expect(err.Error()).To(Equal(util.SignupInvalidEmailMessage))
		})

		g.It("Should return an error when new password doesn't match length", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			err = passwordManager.ChangePassword(user, invalidPassword)
			Expect(err.Error()).To(Equal(util.SignupPasswordLengthMessage))
		})

		g.It("Should return an error when old password doesn't match", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			user.Password = "anotherPassword"
			err = passwordManager.ChangePassword(user, "Asdqwe123")
			Expect(err).NotTo(BeNil())
		})

		g.It("Should end successfully when change password happened correctly", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			_ = passwordManager.ChangePassword(user, "Asdqwe123")
			_, err = auth.Login(user.Email, "Asdqwe123")
			Expect(err).To(BeNil())
		})

	})

	g.Describe("Send Reset Password Request process", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			passwordManager = NewPasswordManager(userRepo, sender)
			auth = NewAuthenticator(userRepo, sender)
		})

		g.It("Should return an error when an invalid email is provided", func() {
			_, err := passwordManager.ForgotPassword(invalidEmail)
			Expect(err.Error()).To(Equal(util.SignupInvalidEmailMessage))
		})

		g.It("Should return an error when the email is not present in the storage", func() {
			_, err := passwordManager.ForgotPassword(validEmail)
			Expect(err.Error()).To(Equal("Email does not exist"))
		})

		g.It("Should generate a code and send an email", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			_, err := passwordManager.ForgotPassword(validEmail)
			Expect(err).To(BeNil())
		})
	})

	g.Describe("Reset Password process", func() {
		g.BeforeEach(func() {
			userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
			passwordManager = NewPasswordManager(userRepo, sender)
			auth = NewAuthenticator(userRepo, sender)
		})

		g.It("Should return an error when an invalid email is provided", func() {
			err := passwordManager.ResetPassword(invalidEmail, "0", "")
			Expect(err.Error()).To(Equal(util.SignupInvalidEmailMessage))
		})

		g.It("Should return an error when new password doesn't match length", func() {
			err := passwordManager.ResetPassword(validEmail, "0", invalidPassword)
			Expect(err.Error()).To(Equal(util.SignupPasswordLengthMessage))
		})

		g.It("Should return an error if the provided code is not correct", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			_, _ = passwordManager.ForgotPassword(validEmail)

			err := passwordManager.ResetPassword(validEmail, "0", validPassword)
			Expect(err.Error()).To(Equal(resetPasswordWrongCodeError))
		})

		g.It("Should change the password given the correct code", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			code, _ := passwordManager.ForgotPassword(validEmail)

			_ = passwordManager.ResetPassword(validEmail, code, "secured_password2")
			_, err := auth.Login(user.Email, "secured_password2")
			Expect(err).To(BeNil())
		})
	})

}
