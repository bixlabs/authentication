package in_memory

import (
	"github.com/bixlabs/authentication/authenticator/structures"
	"strconv"
)

type UserRepo struct {
	lastId int
	users  map[string]structures.User
}

// We don't use data mappers here because this implementation is merely for testing purpose.
// and the things we are going to do here are trivial
func NewUserRepo() *UserRepo {
	return &UserRepo{0, make(map[string]structures.User)}
}

func (u *UserRepo) Create(user structures.User) (structures.User, error) {
	u.lastId = u.lastId + 1
	user.ID = strconv.Itoa(u.lastId)
	u.users[user.Email] = user
	return user, nil
}

func (u *UserRepo) IsEmailAvailable(email string) (bool, error) {
	_, isUsed := u.users[email]
	return !isUsed, nil
}

func (u *UserRepo) VerifyPassword(password string) (bool, error) {
	panic("implement me")
}

func (u *UserRepo) VerifyResetPasswordToken(token string) (bool, error) {
	panic("implement me")
}

func (u *UserRepo) ChangePassword(password string) error {
	panic("implement me")
}

func (u *UserRepo) SaveResetPasswordToken(token string) error {
	panic("implement me")
}
