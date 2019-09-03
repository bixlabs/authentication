package implementation

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

type passwordManager struct {
	repository           user.Repository
	emailSender          email.Sender
	ResetPasswordCodeMax int `env:"AUTH_SERVER_RESET_PASSWORD_MAX" envDefault:"99999"`
	ResetPasswordCodeMin int `env:"AUTH_SERVER_RESET_PASSWORD_MIN" envDefault:"10000"`
}

// NewPasswordManager returns a new instance of the passwordManager
func NewPasswordManager(repository user.Repository, sender email.Sender) interactors.PasswordManager {
	pm := passwordManager{repository: repository, emailSender: sender}
	err := env.Parse(&pm)
	if err != nil {
		tools.Log().Panic("Parsing the env variables for the password manager failed", err)
	}
	return pm
}

func (pm passwordManager) ChangePassword(user structures.User, newPassword string) error {
	loggerFields := logrus.Fields{"email": user.Email, "meth": "passwordManager:ChangePassword"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(user.Email); err != nil {
		contextLogger.WithError(err).Debug("invalid email provided")

		return err
	}

	if user.Password == newPassword {
		contextLogger.Debug("same password provided")

		return util.SamePasswordChangeError{}
	}

	if err := util.CheckPasswordLength(newPassword); err != nil {
		contextLogger.Debug("invalid password length")

		return err
	}

	if err := pm.isPasswordMatch(user); err != nil {
		contextLogger.Debug("password did not match")

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
	loggerFields := logrus.Fields{"email": email, "meth": "passwordManager:hashAndSavePassword"}
	contextLogger := tools.Log().WithFields(loggerFields)

	hashedPassword, err := util.HashPassword(newPassword)

	if err != nil {
		contextLogger.Debug("an error happened when hashing the new password")

		return err
	}

	if err := pm.repository.ChangePassword(email, hashedPassword); err != nil {
		return err
	}
	return nil
}

func (pm passwordManager) ForgotPassword(email string) (string, error) {
	loggerFields := logrus.Fields{"email": email, "meth": "passwordManager:ForgotPassword"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(email); err != nil {
		contextLogger.WithError(err).Debug("invalid email provided")

		return "", err
	}

	userAccount, err := pm.repository.Find(email)
	if err != nil {
		contextLogger.WithError(err).Debug("wrong email provided")

		return "", err
	}

	code, err := pm.generateCode(userAccount)
	if err != nil {
		contextLogger.WithError(err).Debug("an error happened generating the forgot code")

		return "", err
	}

	return code, pm.emailSender.ForgotPasswordRequest(userAccount, code)
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
	loggerFields := logrus.Fields{"email": email, "meth": "passwordManager:ResetPassword"}
	contextLogger := tools.Log().WithFields(loggerFields)

	if err := util.IsValidEmail(email); err != nil {
		contextLogger.WithError(err).Debug("invalid email provided")

		return err
	}

	if err := util.CheckPasswordLength(newPassword); err != nil {
		contextLogger.Debug("invalid password length")

		return err
	}

	if err := pm.isValidCode(email, code); err != nil {
		contextLogger.WithError(err).Debug("invalid code provided")

		return err
	}

	if err := pm.isNewPasswordSameAsOld(email, newPassword); err != nil {
		contextLogger.Debug("same password provided")

		return err
	}

	if err := pm.hashAndSavePassword(email, newPassword); err != nil {
		contextLogger.WithError(err).Debug("an error happened hashing and saving the password")

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
