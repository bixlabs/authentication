package test

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/database/user/sqlite"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

const validEmail = "test@email.com"
const validPassword = "secured_password"

func TestIntegration(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator
	var passwordManager interactors.PasswordManager
	var databaseName = "test.s3db"
	var closeDB func()
	var userRepo user.Repository
	var newPassword = "secure_password2"

	configureDatabaseEnvironments(databaseName)

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Integration tests", func() {
		g.BeforeEach(func() {
			userRepo, closeDB = sqlite.NewSqliteStorage()
			auth = implementation.NewAuthenticator(userRepo, memory.DummySender{})
			passwordManager = implementation.NewPasswordManager(userRepo, memory.DummySender{})
		})

		g.AfterEach(func() {
			closeDB()
			err := os.Remove(databaseName)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should be able to login after signup ", func() {
			aUser := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(aUser)
			Expect(err).To(BeNil())
			response, err := auth.Login(aUser.Email, aUser.Password)
			Expect(err).To(BeNil())
			Expect(response.Token).ToNot(Equal(BeEmpty()))
			Expect(response.User.Email).To(Equal(validEmail))
		})

		g.It("Should be able to change the password", func() {
			aUser := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(aUser)
			Expect(err).To(BeNil())
			err = passwordManager.ChangePassword(aUser, newPassword)
			Expect(err).To(BeNil())
			response, err := auth.Login(aUser.Email, newPassword)
			Expect(err).To(BeNil())
			Expect(response.Token).ToNot(Equal(BeEmpty()))
			Expect(response.User.Email).To(Equal(validEmail))
		})

		g.It("Should be able to forgot password and then reset the password ", func() {
			aUser := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(aUser)
			Expect(err).To(BeNil())
			code, err := passwordManager.ForgotPassword(aUser.Email)
			Expect(err).To(BeNil())
			err = passwordManager.ResetPassword(aUser.Email, code, newPassword)
			Expect(err).To(BeNil())
			response, err := auth.Login(aUser.Email, newPassword)
			Expect(err).To(BeNil())
			Expect(response.Token).ToNot(Equal(BeEmpty()))
			Expect(response.User.Email).To(Equal(validEmail))
		})
	})
}

func configureDatabaseEnvironments(name string) {
	err := os.Setenv("AUTH_SERVER_DATABASE_NAME", name)
	if err != nil {
		panic(err)
	}
	err = os.Setenv("AUTH_SERVER_DATABASE_USER", "admin")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("AUTH_SERVER_DATABASE_PASSWORD", "password")
	if err != nil {
		panic(err)
	}
	err = os.Setenv("AUTH_SERVER_DATABASE_SALT", "salt")
	if err != nil {
		panic(err)
	}
}
