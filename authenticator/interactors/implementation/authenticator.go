package implementation

import (
	"encoding/json"
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

const signupDuplicateEmailMessage = "Email is already taken"
const emailValidationRegex = "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
const signupInvalidEmailMessage = "Email is not valid"
const signupPasswordLengthMessage = "Password should have at least 8 characters"
const passwordMinLength = 8
const maxNumberOfResetPasswordCodes = 99999
const minNumberOfResetPasswordCodes = 10000

type authenticator struct {
	repository     user.Repository
	sender email.Sender
	ExpirationTime int    `env:"TOKEN_EXPIRATION" envDefault:"3600"`
	Secret         string `env:"AUTH_SERVER_SECRET"`
}

func NewAuthenticator(repository user.Repository, sender email.Sender) interactors.Authenticator {
	auth := &authenticator{repository: repository, sender: sender}
	err := env.Parse(auth)
	if err != nil {
		tools.Log().Panic("Parsing the env variables for the authenticator failed", err)
	}
	return auth
}

func (auth authenticator) Login(email, password string) (*login.Response, error) {
	if err := isValidEmail(email); err != nil {
		return nil, err
	}

	if err := checkPasswordLength(password); err != nil {
		return nil, err
	}

	hashedPassword, err := auth.repository.GetHashedPassword(email)

	if err != nil {
		return nil, wrongCredentialsError()
	}

	if err := verifyPassword(hashedPassword, password); err != nil {
		return nil, err
	}

	return generateJWT(email, auth)
}

func generateJWT(email string, auth authenticator) (*login.Response, error) {
	currentUser, err := auth.repository.Find(email)
	if err != nil {
		return nil, wrongCredentialsError()
	}

	response := &login.Response{User: currentUser, IssuedAt: time.Now().Unix(), Expiration: time.Now().Add(time.Second * time.Duration(auth.ExpirationTime)).Unix()}

	if err := setToken(response, auth.Secret); err != nil {
		return nil, err
	}

	return response, nil
}

func setToken(response *login.Response, secret string) error {
	jsonUser, err := json.Marshal(response.User)
	if err != nil {
		tools.Log().Error("Parsing user json failed", err)
		return err
	}
	tokenString, err := generateClaims(*response, string(jsonUser)).SignedString([]byte(secret))
	if err != nil {
		tools.Log().Error("Generating jwt signed token failed", err)
		return err
	}

	response.Token = tokenString
	return nil
}

func generateClaims(response login.Response, jsonUser string) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":  response.IssuedAt,
		"exp":  response.Expiration,
		"user": jsonUser,
	})
}

func (auth authenticator) Signup(user structures.User) (structures.User, error) {
	if err := auth.hasValidationIssue(user); err != nil {
		return user, err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword

	user, err = auth.repository.Create(user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (auth authenticator) hasValidationIssue(user structures.User) error {
	err := isValidEmail(user.Email)
	if err != nil {
		return err
	}

	if isAvailable, err := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return errors.New(signupDuplicateEmailMessage)
	}

	if err := checkPasswordLength(user.Password); err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) error {
	if isValidEmail, _ := regexp.MatchString(emailValidationRegex, email); !isValidEmail {
		tools.Log().Debug("An invalid email was provided: " + email)
		return errors.New(signupInvalidEmailMessage)
	}
	return nil
}

func checkPasswordLength(password string) error {
	if len(password) < passwordMinLength {
		tools.Log().Debug("A password with incorrect length was provided")
		return errors.New(signupPasswordLengthMessage)
	}
	return nil
}

func hashPassword(password string) (string, error) {
	pass := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		tools.Log().WithField("error", err).Error("Password hash failed")
		return "", err
	}
	return string(hashedPassword), nil
}

func (auth authenticator) ChangePassword(user structures.User, newPassword string) error {
	oldHashedPassword, err := auth.repository.GetHashedPassword(user.Email)
	if err != nil {
		return err
	}

	err = verifyPassword(oldHashedPassword, user.Password)

	if err != nil {
		return err
	}

	if err := checkPasswordLength(user.Password); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(newPassword)

	if err != nil {
		return err
	}

	return auth.repository.ChangePassword(user.Email, hashedPassword)
}

func verifyPassword(hashedPassword, plainPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		tools.Log().Debug("A wrong password was provided")
		return wrongCredentialsError()
	}
	return nil
}

func wrongCredentialsError() error {
	return errors.New("wrong credentials")
}

func (auth authenticator) ResetPassword(email string, code string, newPassword string) error {
	return nil
}

func (auth authenticator) SendResetPasswordRequest(email string) error {
	if err := isValidEmail(email); err != nil {
		return err
	}

	userAccount, err := auth.repository.Find(email)
	if err != nil {
		return err
	}

	code, err := auth.generateCode(userAccount)
	if err != nil {
		return err
	}

	return auth.sender.SendEmailPasswordRequest(userAccount, code)
}

func(auth authenticator) generateCode(user structures.User) (string, error) {
	code := generateRandomNumber()
	resetToken, err := hashPassword(code)
	if err != nil {
		return "", err
	}
	if err := auth.repository.SaveResetToken(user.Email, resetToken); err != nil {
		return "", err
	}

    return code, nil
}

func generateRandomNumber() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(maxNumberOfResetPasswordCodes - minNumberOfResetPasswordCodes) +
		minNumberOfResetPasswordCodes)
}

