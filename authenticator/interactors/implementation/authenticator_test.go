package implementation

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const email = "test@email.com"
const badEmail = "invalid_email"
const validPassword = "secured_password"
const invalidPassword = "07chars"

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Signup process", func() {
		g.BeforeEach(func() {
			auth = NewAuthenticator(in_memory.NewUserRepo())
		})

		g.It("Should check for email duplication ", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, _ = auth.Signup(user)
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupDuplicateEmailMessage)
		})

		g.It("Should check for invalid email ", func() {
			user := structures.User{Email: badEmail, Password: validPassword}
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupInvalidEmailMessage)
		})

		g.It("Should have a password of at least 8 characters ", func() {
			user := structures.User{Email: email, Password: invalidPassword}
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupPasswordLengthMessage)
		})

		g.It("Should create a user with an ID in it", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, user = auth.Signup(user)
			g.Assert(user.ID).Equal("1")
		})

		g.It("Should hash the password of the user", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, user = auth.Signup(user)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(validPassword))
			g.Assert(err).Equal(nil)
		})

		g.It("Should hash the password and not be able to match when given a wrong one", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, user = auth.Signup(user)
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("wrong_password"))
			Expect(err).NotTo(BeNil())
		})
	})
}
