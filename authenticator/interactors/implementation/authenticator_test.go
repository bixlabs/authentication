package implementation

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator
	g.Describe("Signup process", func() {
		g.BeforeEach(func() {
			auth = NewAuthenticator(in_memory.NewUserRepo())
		})

		g.It("Should check for email duplication ", func() {
			user := structures.User{Email: "test@email.com", Password: "secure_password"}
			_, _ = auth.Signup(user)
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupDuplicateEmailMessage)
		})

		g.It("Should check for invalid email ", func() {
			user := structures.User{Email: "invalid_email", Password: "secure_password"}
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupInvalidEmailMessage)
		})

		g.It("Should have a password of at least 8 characters ", func() {
			user := structures.User{Email: "test2@email.com", Password: "07chars"}
			err, _ := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupPasswordLengthMessage)
		})

		g.It("Should create a user with an ID in it", func() {
			user := structures.User{Email: "test2@email.com", Password: "secure_password"}
			_, user = auth.Signup(user)
			g.Assert(user.ID).Equal("1")
		})
	})
}
