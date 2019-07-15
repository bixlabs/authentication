package implementation

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"math/rand"
	"strconv"
	"time"
)

type passwordManager struct {
	repository           user.Repository
	sender               email.Sender
	ResetPasswordCodeMax int `env:"AUTH_SERVER_RESET_PASSWORD_MAX" envDefault:"99999"`
	ResetPasswordCodeMin int `env:"AUTH_SERVER_RESET_PASSWORD_MIN" envDefault:"10000"`
}

func NewPasswordManager(repository user.Repository, sender email.Sender) interactors.PasswordManager {
	pm := passwordManager{repository: repository, sender: sender}
	err := env.Parse(&pm)
	if err != nil {
		tools.Log().Panic("Parsing the env variables for the password manager failed", err)
	}
	return pm
}

func (pm passwordManager) ChangePassword(user structures.User, newPassword string) error {
	if err := util.IsValidEmail(user.Email); err != nil {
		return err
	}

	if user.Password == newPassword {
		return util.SamePasswordChangeError{}
	}

	if err := util.CheckPasswordLength(newPassword); err != nil {
		return err
	}

	if err := pm.isPasswordMatch(user); err != nil {
		return err
	}

	return pm.hashAndSavePassword(user.Email, newPassword)
}

func (pm passwordManager) isPasswordMatch(user structures.User) error {
	oldHashedPassword, err := pm.repository.GetHashedPassword(user.Email)
	if err != nil {
		return err
	}

	if err := util.VerifyPassword(oldHashedPassword, user.Password); err != nil {
		return err
	}

	return nil
}

func (pm passwordManager) hashAndSavePassword(email, newPassword string) error {
	hashedPassword, err := util.HashPassword(newPassword)

	if err != nil {
		return err
	}

	if err := pm.repository.ChangePassword(email, hashedPassword); err != nil {
		return err
	}
	return nil
}


func (pm passwordManager) ForgotPassword(email string) (string, error) {
	if err := util.IsValidEmail(email); err != nil {
		return "", err
	}

	userAccount, err := pm.repository.Find(email)
	if err != nil {
		return "", err
	}

	code, err := pm.generateCode(userAccount)
	if err != nil {
		return "", err
	}

	return code, pm.sender.SendEmailPasswordRequest(userAccount, code)
}

func (pm passwordManager) generateCode(user structures.User) (string, error) {
	code := pm.generateRandomNumber()
	resetToken, err := util.HashPassword(code)
	if err != nil {
		return "", err
	}
	if err := pm.repository.UpdateResetToken(user.Email, resetToken); err != nil {
		return "", err
	}

	return code, nil
}

func (pm passwordManager) generateRandomNumber() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(pm.ResetPasswordCodeMax-pm.ResetPasswordCodeMin) +
		pm.ResetPasswordCodeMin)
}

func (pm passwordManager) ResetPassword(email string, code string, newPassword string) error {
	if err := util.IsValidEmail(email); err != nil {
		return err
	}

	if err := util.CheckPasswordLength(newPassword); err != nil {
		return err
	}

	if err := pm.isValidCode(email, code); err != nil {
		return err
	}

	if err := pm.isNewPasswordSameAsOld(email, newPassword); err != nil {
		return err
	}

	if err := pm.hashAndSavePassword(email, newPassword); err != nil {
		return err
	}

	return pm.repository.UpdateResetToken(email, "")
}

func (pm passwordManager) isValidCode(email, code string) error {
	account, err := pm.repository.Find(email)

	if err != nil {
		return err
	}

	if err := util.VerifyPassword(account.ResetToken, code); err != nil {
		return util.InvalidResetPasswordCode{}
	}

	return nil
}

func (pm passwordManager) isNewPasswordSameAsOld(email, newPassword string) error {
	oldHashedPassword, err := pm.repository.GetHashedPassword(email)
	if err != nil {
		return err
	}

	if err := util.VerifyPassword(oldHashedPassword, newPassword); err == nil {
		return util.SamePasswordChangeError{}
	}
	return nil
}
