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

func (u *UserRepo) Create(user structures.User) (error, structures.User) {
	u.lastId = u.lastId + 1
	user.ID = strconv.Itoa(u.lastId)
	u.users[user.Email] = user
	return nil, user
}

func (u UserRepo) IsEmailAvailable(email string) (error, bool) {
	_, isUsed := u.users[email]
	return nil, !isUsed
}

func (u UserRepo) VerifyPassword(password string) (error, bool) {
	panic("implement me")
}

func (u UserRepo) ChangePassword(password string) error {
	panic("implement me")
}

func (u UserRepo) SaveResetPasswordToken(token string) error {
	panic("implement me")
}

func (u UserRepo) VerifyResetPasswordToken(token string) (error, bool) {
	panic("implement me")
}
