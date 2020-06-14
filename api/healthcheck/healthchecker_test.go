package healthcheck

import (
	"net/http"
	"testing"

	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/database/user/memory"
	email2 "github.com/bixlabs/authentication/email"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
)

func TestHealthCheckRest(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	tools.Log().Level = logrus.FatalLevel

	var (
		auth        interactors.Authenticator
		passManager interactors.PasswordManager
		sender      email.Sender
		userRepo    user.Repository
	)

	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Healthcheck rest handler", func() {
		g.BeforeEach(func() {
			userRepo, sender = memory.NewUserRepo(), email2.NewDummySender()
			auth = implementation.NewAuthenticator(userRepo, sender)
			passManager = implementation.NewPasswordManager(userRepo, sender)
		})

		g.It("should return 500 if any service is down", func() {
			auth = nil
			code, _ := healthCheckHandler(userRepo, sender, auth, passManager)
			Expect(code).To(Equal(http.StatusInternalServerError))
		})

		g.It("should return 200 if all services are running", func() {
			code, _ := healthCheckHandler(userRepo, sender, auth, passManager)
			Expect(code).To(Equal(http.StatusOK))
		})
	})
}