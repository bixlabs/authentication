package user

import "github.com/bixlabs/authentication/authenticator/structures"

type Repository interface {
	Create(user structures.User) (error, structures.User)
	IsEmailAvailable(email string) (error, bool)
	VerifyPassword(password string) (error, bool)
	ChangePassword(password string) error
	SaveResetPasswordToken(token string) error
	VerifyResetPasswordToken(token string) (error, bool)
}

//TODO: This is temporal until we have the real implementation, the linter is complaining about the interface being unused.
type FakeRepo struct {
}

func (FakeRepo) IsEmailAvailable(email string) (error, bool) {
	panic("implement me")
}

func (FakeRepo) Create(user structures.User) (error, structures.User) {
	panic("implement me")
}

func (FakeRepo) VerifyPassword(password string) (error, bool) {
	panic("implement me")
}

func (FakeRepo) ChangePassword(password string) error {
	panic("implement me")
}

func (FakeRepo) SaveResetPasswordToken(token string) error {
	panic("implement me")
}

func (FakeRepo) VerifyResetPasswordToken(token string) (error, bool) {
	panic("implement me")
}
