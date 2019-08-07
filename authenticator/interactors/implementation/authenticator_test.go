package implementation

import (
	databaseUserPackage "github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"os"
	"testing"
	"time"
)

const validEmail = "test@email.com"
const invalidEmail = "invalid_email"
const validPassword = "secured_password"
const invalidPassword = "07chars"

func TestAuthenticator(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	//tools.Log().Level = logrus.FatalLevel
	var auth interactors.Authenticator
	var repo databaseUserPackage.Repository

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Signup process", func() {
		g.BeforeEach(func() {
			repo = memory.NewUserRepo()
			auth = NewAuthenticator(repo, memory.DummySender{}, NewUserManager(repo))
		})

		g.It("Should check for email duplication ", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = auth.Signup(user)
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(util.SignupDuplicateEmailMessage)
		})

		g.It("Should check for invalid email ", func() {
			user := structures.User{Email: invalidEmail, Password: validPassword}
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should have a password of at least 8 characters ", func() {
			user := structures.User{Email: validEmail, Password: invalidPassword}
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(util.SignupPasswordLengthMessage)
		})

		g.It("Should create a user with an ID in it", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			g.Assert(user.ID).Equal("1")
		})

		g.It("Should hash the password of the user", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}

			// TODO: we should use an utility for this in the overall file
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(validPassword))
			g.Assert(err).Equal(nil)
		})

		g.It("Should hash the password and not be able to match when given a wrong one", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("wrong_password"))
			Expect(err).NotTo(BeNil())
		})
	})

	g.Describe("Login process", func() {
		g.BeforeEach(func() {
			repo := memory.NewUserRepo()
			auth = NewAuthenticator(memory.NewUserRepo(), memory.DummySender{}, NewUserManager(repo))
			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should validate the provided email", func() {
			_, err := auth.Login("wrong_email", "")
			Expect(err.Error()).To(Equal(util.SignupInvalidEmailMessage))
		})

		g.It("Should validate password length", func() {
			_, err := auth.Login("test@test.com", "123456")
			Expect(err.Error()).To(Equal("Password should have at least 8 characters"))
		})

		g.It("Should login if the provided credentials are correct", func() {
			response, err := auth.Login(validEmail, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.User.Email).To(Equal(validEmail))
		})

		g.It("Should provide a JWT token after successful login", func() {
			response, err := auth.Login(validEmail, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.Token).ToNot(Equal(""))
		})

		g.It("Should have an issuedAt and a Expiration date correctly set that are one hour apart", func() {
			response, err := auth.Login(validEmail, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.Expiration - response.IssuedAt).To(Equal(int64(3600)))
		})
	})

	g.Describe("Verify JWT", func() {
		g.BeforeEach(func() {
			secret := "test"
			err := os.Setenv("AUTH_SERVER_SECRET", secret)
			Expect(err).To(BeNil())
			repo := memory.NewUserRepo()
			auth = NewAuthenticator(repo, memory.DummySender{}, NewUserManager(repo))
		})

		g.It("Should return an error in case the JWT token is invalid", func() {
			_, err := auth.VerifyJWT("invalid_token")
			_, ok := err.(util.InvalidJWTToken)
			Expect(ok).To(Equal(true))
		})

		g.It("Should return the user claim in case the JWT token is valid", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			response := login.Response{Expiration: time.Now().Add(1000 * time.Second).Unix(),
				IssuedAt: time.Now().Unix(),
				User:     user}
			aToken, err := generateClaims(response).SignedString([]byte("test"))
			Expect(err).To(BeNil())
			time.Sleep(100 * time.Millisecond)
			result, err := auth.VerifyJWT(aToken)
			Expect(err).To(BeNil())
			Expect(result.Email).To(Equal(validEmail))
		})
	})
}
