package implementation

import (
	"encoding/json"
	"errors"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/joho/godotenv/autoload"
	"time"
)

const signupDuplicateEmailMessage = "Email is already taken"

type authenticator struct {
	repository     user.Repository
	sender         email.Sender
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
	if err := util.IsValidEmail(email); err != nil {
		return nil, err
	}

	if err := util.CheckPasswordLength(password); err != nil {
		return nil, err
	}

	hashedPassword, err := auth.repository.GetHashedPassword(email)

	if err != nil {
		return nil, util.WrongCredentialsError{}
	}

	if err := util.VerifyPassword(hashedPassword, password); err != nil {
		return nil, err
	}

	return generateJWT(email, auth)
}

func generateJWT(email string, auth authenticator) (*login.Response, error) {
	currentUser, err := auth.repository.Find(email)
	if err != nil {
		return nil, util.WrongCredentialsError{}
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

	hashedPassword, err := util.HashPassword(user.Password)
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
	if err := util.IsValidEmail(user.Email); err != nil {
		return err
	}

	if isAvailable, err := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return errors.New(signupDuplicateEmailMessage)
	}

	if err := util.CheckPasswordLength(user.Password); err != nil {
		return err
	}

	return nil
}
