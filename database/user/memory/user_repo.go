package memory

import (
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"strconv"
	"time"
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
		temp := u.users[email]
		return temp.Password, nil
	}
	tools.Log().Warn("A user couldn't be found when finding the hashed password of it")
	return "", errors.New("user doesn't exist")
}

func (u *UserRepo) ChangePassword(email, newPassword string) error {
	if temp, ok := u.users[email]; ok {
		temp.Password = newPassword
		u.users[email] = temp
	}

	return nil
}

func (u *UserRepo) UpdateResetToken(email, resetToken string) error {
	temp := u.users[email]
	temp.ResetToken = resetToken
	u.users[temp.Email] = temp
	return nil
}

func (u *UserRepo) Find(email string) (structures.User, error) {
	temp, exist := u.users[email]
	if !exist || temp.DeletedAt != nil {
		return structures.User{}, errors.New("email does not exist")
	}
	return temp, nil
}

// Precondition: user should already exist
func (u *UserRepo) Delete(user structures.User) error {
	if temp, ok := u.users[user.Email]; ok {
		deletedAt := time.Now()
		temp.DeletedAt = &deletedAt
		u.users[user.Email] = temp
	}

	return nil
}

func (u *UserRepo) Update(email string, updateAttrs structures.User) (structures.User, error) {
	_, err := u.Find(email)

	if err != nil {
		return structures.User{}, err
	}

	if email != updateAttrs.Email {
		delete(u.users, email)
		email = updateAttrs.Email
	}

	u.users[email] = updateAttrs

	return updateAttrs, nil
}
