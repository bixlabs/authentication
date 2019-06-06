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
			//repo = in_memory.NewUserRepo()
			auth = NewAuthenticator(in_memory.NewUserRepo())
		})

		g.It("Should check for email duplication ", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, _ = auth.Signup(user)
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupDuplicateEmailMessage)
		})

		g.It("Should check for invalid email ", func() {
			user := structures.User{Email: badEmail, Password: validPassword}
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupInvalidEmailMessage)
		})

		g.It("Should have a password of at least 8 characters ", func() {
			user := structures.User{Email: email, Password: invalidPassword}
			_, err := auth.Signup(user)
			g.Assert(err.Error()).Equal(signupPasswordLengthMessage)
		})

		g.It("Should create a user with an ID in it", func() {
			user := structures.User{Email: email, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			g.Assert(user.ID).Equal("1")
		})

		g.It("Should hash the password of the user", func() {
			user := structures.User{Email: email, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(validPassword))
			g.Assert(err).Equal(nil)
		})

		g.It("Should hash the password and not be able to match when given a wrong one", func() {
			user := structures.User{Email: email, Password: validPassword}
			user, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("wrong_password"))
			Expect(err).NotTo(BeNil())
		})
	})

	g.Describe("Change password process", func() {
		g.BeforeEach(func() {
			auth = NewAuthenticator(in_memory.NewUserRepo())
		})

		g.It("Should return an error when old password doesn't match", func() {
			user := structures.User{Email: email, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
			user.Password = "anotherPassword"
			err = auth.ChangePassword(user, "Asdqwe123")
			Expect(err).NotTo(BeNil())
		})

		//g.It("Should be able to change password in case the provided data is correct", func() {
		//	user := structures.User{Email: email, Password: validPassword}
		//	_, _ = auth.Signup(user)
		//	_ = auth.ChangePassword(user, "12345678")
		//	hashedPassword, _ := repo.GetHashedPassword(user.Email)
		//	err := verifyPassword(hashedPassword, "12345678")
		//	Expect(err).To(BeNil())
		//})
	})

	g.Describe("Login process", func() {
		g.BeforeEach(func() {
			auth = NewAuthenticator(in_memory.NewUserRepo())
			user := structures.User{Email: email, Password: validPassword}
			_, err := auth.Signup(user)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should validate the provided email", func() {
			_, err := auth.Login("wrong_email", "")
			Expect(err.Error()).To(Equal(signupInvalidEmailMessage))
		})

		g.It("Should validate password length", func() {
			_, err := auth.Login("test@test.com", "123456")
			Expect(err.Error()).To(Equal("Password should have at least 8 characters"))
		})

		g.It("Should login if the provided credentials are correct", func() {
			response, err := auth.Login(email, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.User.Email).To(Equal(email))
		})

		g.It("Should provide a JWT token after successful login", func() {
			response, err := auth.Login(email, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.Token).ToNot(Equal(""))
		})

		g.It("Should have an issuedAt and a Expiration date correctly set that are one hour apart", func() {
			response, err := auth.Login(email, validPassword)
			if err != nil {
				panic(err)
			}
			Expect(response.Expiration - response.IssuedAt).To(Equal(int64(3600)))
		})

	})
}
