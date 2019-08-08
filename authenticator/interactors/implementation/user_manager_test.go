package implementation

import (
	databaseUserPackage "github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/tools"
	"github.com/franela/goblin"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestUserManager(t *testing.T) {
	g := goblin.Goblin(t)
	tools.InitializeLogger()
	// This line prevents the logs to appear in the tests.
	//tools.Log().Level = logrus.FatalLevel
	var um interactors.UserManager
	var userRepo databaseUserPackage.Repository

	//special hook for gomega
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Create User process", func() {
		g.BeforeEach(func() {
			userRepo = memory.NewUserRepo()
			um = NewUserManager(userRepo)
		})

		g.It("Should return an error in case the email is invalid", func() {
			user := structures.User{Email: invalidEmail, Password: validPassword}
			_, err := um.Create(user)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should have a password of at least 8 characters", func() {
			user := structures.User{Email: validEmail, Password: invalidPassword}
			_, err := um.Create(user)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupPasswordLengthMessage)
		})

		g.It("Should create a user with an ID", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := um.Create(user)

			Expect(err).To(BeNil())
			g.Assert(user.ID).Equal("1")
		})

		g.It("Should create a random password if not provided", func() {
			user := structures.User{Email: validEmail}
			user, err := um.Create(user)

			Expect(err).To(BeNil())

			savedUser, _ := userRepo.Find(user.Email)

			err = bcrypt.CompareHashAndPassword([]byte(savedUser.Password), []byte(user.GeneratedPassword))
			g.Assert(err).Equal(nil)
		})

		g.It("Should hash the password of the user", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := um.Create(user)

			Expect(err).To(BeNil())

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(validPassword))
			g.Assert(err).Equal(nil)
		})

		g.It("Should hash the password and not be able to match when given a wrong one", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			user, err := um.Create(user)

			Expect(err).To(BeNil())

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(invalidPassword))
			Expect(err).NotTo(BeNil())
		})

		g.It("Should check for email duplication", func() {
			user := structures.User{Email: validEmail, Password: validPassword}
			_, _ = um.Create(user)
			_, err := um.Create(user)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupDuplicateEmailMessage)
		})
	})

	g.Describe("Delete User process", func() {
		g.BeforeEach(func() {
			um = NewUserManager(memory.NewUserRepo())

			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := um.Create(user)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error in case the email is invalid", func() {
			err := um.Delete(invalidEmail)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should return an error in case the user does not exist", func() {
			err := um.Delete("nonexistingemail@example.com")

			Expect(err).NotTo(BeNil())
			g.Assert(err.Error()).Equal(util.UserNotFoundMessage)
		})

		g.It("Should delete a valid user", func() {
			_ = um.Delete(validEmail)
			err := um.Delete(validEmail)
			Expect(err).NotTo(BeNil())
		})

		g.It("Should return an error in case the user was already deleted", func() {
			_ = um.Delete(validEmail)
			err := um.Delete(validEmail)

			Expect(err).NotTo(BeNil())
			g.Assert(err.Error()).Equal(util.UserNotFoundMessage)
		})
	})

	g.Describe("Find User process", func() {
		g.Before(func() {
			um = NewUserManager(memory.NewUserRepo())

			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := um.Create(user)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error in case the email is invalid", func() {
			_, err := um.Find(invalidEmail)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should return an error in case the user does not exist", func() {
			_, err := um.Find("nonexistingemail@example.com")

			Expect(err).NotTo(BeNil())
			g.Assert(err.Error()).Equal(util.UserNotFoundMessage)
		})

		g.It("Should find a valid user", func() {
			user, err := um.Find(validEmail)

			Expect(err).To(BeNil())
			g.Assert(user.Email).Equal(validEmail)
		})
	})

	g.Describe("Update User process", func() {
		g.Before(func() {
			um = NewUserManager(memory.NewUserRepo())

			user := structures.User{Email: validEmail, Password: validPassword}
			_, err := um.Create(user)
			if err != nil {
				panic(err)
			}
		})

		g.It("Should return an error in case the email is invalid", func() {
			updateAttrs := structures.UpdateUser{GivenName: "GivenName"}
			_, err := um.Update(invalidEmail, updateAttrs)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should return an error in case the user does not exist", func() {
			updateAttrs := structures.UpdateUser{}
			_, err := um.Update("nonexistingemail@example.com", updateAttrs)

			Expect(err).NotTo(BeNil())
			g.Assert(err.Error()).Equal(util.UserNotFoundMessage)
		})

		g.It("Should return an error in case the update email is invalid", func() {
			updateAttrs := structures.UpdateUser{Email: invalidEmail, Password: validPassword}
			_, err := um.Update(validEmail, updateAttrs)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupInvalidEmailMessage)
		})

		g.It("Should return an error in case the update password does not have at least 8 characters", func() {
			updateAttrs := structures.UpdateUser{Email: validEmail, Password: invalidPassword}
			_, err := um.Update(validEmail, updateAttrs)

			Expect(err).NotTo(BeNil())

			// TODO: rename this error property
			g.Assert(err.Error()).Equal(util.SignupPasswordLengthMessage)
		})

		g.It("Should Update a valid user", func() {
			updateAttrs := structures.UpdateUser{
				Email:            "newemail@example.com",
				Password:         "newPassWord",
				GivenName:        "newGivenName",
				SecondName:       "newSecondName",
				FamilyName:       "newFamilyName",
				SecondFamilyName: "newSecondFamilyName",
			}
			updatedUser, err := um.Update(validEmail, updateAttrs)

			Expect(err).To(BeNil())

			Expect(updatedUser.ID).NotTo(BeNil())
			g.Assert(updatedUser.Email).Equal(updateAttrs.Email)
			g.Assert(updatedUser.GivenName).Equal(updateAttrs.GivenName)
			g.Assert(updatedUser.SecondName).Equal(updateAttrs.SecondName)
			g.Assert(updatedUser.FamilyName).Equal(updateAttrs.FamilyName)
			g.Assert(updatedUser.SecondFamilyName).Equal(updateAttrs.SecondFamilyName)

			err = bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(updateAttrs.Password))
			g.Assert(err).Equal(nil)
		})
	})
}
