package memory

import (
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"strconv"
)

type UserRepo struct {
	lastID int
	users  map[string]structures.User
}

// We don't use data mappers here because this implementation is merely for testing purpose.
// and the things we are going to do here are trivial
func NewUserRepo() user.Repository {
	return &UserRepo{0, make(map[string]structures.User)}
}

func (u *UserRepo) Create(user structures.User) (structures.User, error) {
	u.lastID++
	user.ID = strconv.Itoa(u.lastID)
	u.users[user.Email] = user
	return user, nil
}

func (u *UserRepo) IsEmailAvailable(email string) (bool, error) {
	_, isUsed := u.users[email]
	return !isUsed, nil
}

func (u *UserRepo) GetHashedPassword(email string) (string, error) {
	if isAvailable, err := u.IsEmailAvailable(email); err == nil && !isAvailable {
		user := u.users[email]
		return user.Password, nil
	}
	tools.Log().Warn("A user couldn't be found when finding the hashed password of it")
	return "", errors.New("user doesn't exist")
}

func (u *UserRepo) ChangePassword(email, newPassword string) error {
	user := u.users[email]
	user.Password = newPassword
	u.users[email] = user
	return nil
}

func (u *UserRepo) UpdateResetToken(email, resetToken string) error {
	user := u.users[email]
	user.ResetToken = resetToken
	u.users[user.Email] = user
	return nil
}

func (u *UserRepo) Find(email string) (structures.User, error) {
	user, exist := u.users[email]
	if !exist {
		return structures.User{}, errors.New("email does not exist")
	}
	user.Password = ""
	return user, nil
}
