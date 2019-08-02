package implementation

import (
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/authenticator/structures/login"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/dgrijalva/jwt-go"
	"time"
)

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

	response := &login.Response{User: currentUser, IssuedAt: time.Now().Unix(),
		Expiration: time.Now().Add(time.Second * time.Duration(auth.ExpirationTime)).Unix()}

	if err := setToken(response, auth.Secret); err != nil {
		return nil, err
	}

	return response, nil
}

func setToken(response *login.Response, secret string) error {
	tokenString, err := generateClaims(*response).SignedString([]byte(secret))
	if err != nil {
		tools.Log().Error("Generating jwt signed token failed", err)
		return err
	}

	response.Token = tokenString
	return nil
}

func generateClaims(response login.Response) *jwt.Token {
	c := userClaims{
		User: response.User,
		StandardClaims: &jwt.StandardClaims{
			IssuedAt:  response.IssuedAt,
			ExpiresAt: response.Expiration,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &c)
}

type userClaims struct {
	User structures.User `json:"user,omitempty"`
	*jwt.StandardClaims
}

// TODO: This is a workaround because jwt-go is validating iat when it shouldn't (jwt specification doesn't say so)
// let's remove this later when jwt-go removes the iat validation in v4.
func (c *userClaims) Valid() error {
	c.StandardClaims.IssuedAt /= 10
	valid := c.StandardClaims.Valid()
	c.StandardClaims.IssuedAt *= 10
	return valid
}

func (auth authenticator) Signup(user structures.User) (structures.User, error) {
	return auth.createUser(user)
}

func (auth authenticator) hasValidationIssue(user structures.User) error {
	if err := util.IsValidEmail(user.Email); err != nil {
		return err
	}

	if isAvailable, err := auth.repository.IsEmailAvailable(user.Email); err != nil || !isAvailable {
		tools.Log().WithField("error", err).Debug("A duplicated email was provided")
		return util.DuplicatedEmailError{}
	}

	if err := util.CheckPasswordLength(user.Password); err != nil {
		return err
	}

	return nil
}

func (auth authenticator) VerifyJWT(token string) (structures.User, error) {
	jwtToken, err := auth.parseJWTToken(token)
	if err != nil {
		return structures.User{}, err
	}
	return auth.validateAndObtainClaims(*jwtToken)
}

func (auth authenticator) parseJWTToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &userClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(auth.Secret), nil
	})
	if err != nil {
		tools.Log().WithField("error", err).Info("An error happened while validating the JWT token")
		return jwtToken, util.InvalidJWTToken{}
	}

	return jwtToken, nil
}

func (auth authenticator) validateAndObtainClaims(token jwt.Token) (structures.User, error) {
	claims, ok := token.Claims.(*userClaims)
	if !ok {
		tools.Log().Info("Claims object is not of the correct type")
		return structures.User{}, util.InvalidJWTToken{}
	}

	if err := claims.Valid(); err != nil {
		tools.Log().WithField("error", err).Info("An error happened while validating the JWT token")
		return structures.User{}, util.InvalidJWTToken{}
	}
	return claims.User, nil
}

func (auth authenticator) Create(user structures.User) (structures.User, error) {
	if user, err := auth.generatePasswordIfEmpty(&user); err != nil {
		return user, err
	}

	return auth.createUser(user)
}

func (auth authenticator) createUser(user structures.User) (structures.User, error) {
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

func (auth authenticator) generatePasswordIfEmpty(user *structures.User) (structures.User, error) {
	if user.Password != "" {
		return *user, nil
	}

	var err error
	user.GeneratedPassword, err = util.GenerateRandomPassword()
	if err != nil {
		return *user, err
	}

	user.Password = user.GeneratedPassword

	return *user, nil
}

func (auth authenticator) Delete(email string) error {
	if err := util.IsValidEmail(email); err != nil {
		return err
	}

	userToRemove, err := auth.repository.Find(email)
	if err != nil {
		return util.UserNotFoundError{}
	}

	return auth.repository.Delete(userToRemove)
}

func (auth authenticator) Find(email string) (structures.User, error) {
	if err := util.IsValidEmail(email); err != nil {
		return structures.User{}, err
	}

	user, err := auth.repository.Find(email)
	if err != nil {
		return structures.User{}, util.UserNotFoundError{}
	}

	return user, nil
}
