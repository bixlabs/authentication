package authentication

import (
	"errors"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_change_password"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_login"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_reset_password"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_reset_password_request"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_signup"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authenticatorRESTConfigurator struct {
	handler         interactors.Authenticator
	passwordManager interactors.PasswordManager
}

func NewAuthenticatorRESTConfigurator(handler interactors.Authenticator, pm interactors.PasswordManager, router *gin.Engine) {
	configureAuthRoutes(authenticatorRESTConfigurator{handler, pm}, router)
}

func configureAuthRoutes(restConfig authenticatorRESTConfigurator, r *gin.Engine) *gin.Engine {
	router := r.Group("/user")
	router.POST("/login", restConfig.login)
	router.POST("/signup", restConfig.signup)
	router.PUT("/change-password", restConfig.changePassword)
	router.PUT("/reset-password", restConfig.resetPassword)
	router.PUT("/reset-password-request", restConfig.resetPasswordRequest)
	return r
}

// @Summary Login functionality
// @Description Attempts to authenticate the user with the given credentials.
// @Accept  json
// @Produce  json
// @Param login body rest_login.Request true "Login Request"
// @Success 200 {object} rest_login.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 401 {object} rest.ResponseWrapper
// @Failure 403 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 405 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/login [post]
func (config authenticatorRESTConfigurator) login(c *gin.Context) {
	auth := implementation.NewAuthenticator(in_memory.NewUserRepo(), in_memory.DummySender{})
	user := structures.User{Email: "email@bixlabs.com", Password: "password1"}
	_, _ = auth.Signup(user)
	var request rest_login.Request
	if isInvalidLoginRequest(c, &request) {
		c.JSON(http.StatusBadRequest, rest_login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or password missing")))
	} else {
		c.JSON(loginHandler(request.Email, request.Password, auth))
	}
}

func isInvalidLoginRequest(c *gin.Context, request *rest_login.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func loginHandler(email, password string, handler interactors.Authenticator) (int, rest_login.Response) {
	response, err := handler.Login(email, password)
	if err != nil {
		return handleLoginErrors(err)
	}

	return http.StatusOK, rest_login.NewResponse(http.StatusOK, response)
}

func handleLoginErrors(err error) (int, rest_login.Response) {
	if _, ok := err.(util.InvalidEmailError); ok {
		return http.StatusBadRequest, rest_login.NewErrorResponse(http.StatusBadRequest, err)
	}
	if _, ok := err.(util.PasswordLengthError); ok {
		return http.StatusBadRequest, rest_login.NewErrorResponse(http.StatusBadRequest, err)
	}
	if _, ok := err.(util.WrongCredentialsError); ok {
		return http.StatusUnauthorized, rest_login.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, rest_login.NewErrorResponse(http.StatusBadRequest, err)
}

// @Summary Signup functionality
// @Description Attempts to create a user provided the correct information.
// @Accept  json
// @Produce  json
// @Param signup body rest_signup.Request true "Signup Request"
// @Success 201 {object} rest_signup.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/signup [post]
func (config authenticatorRESTConfigurator) signup(c *gin.Context) {
	//rest.NotImplemented(c)
	c.JSON(http.StatusCreated, rest_signup.Response{})
}

// @Summary Change password functionality
// @Description It changes the password provided the old one and a new password.
// @Accept  json
// @Produce  json
// @Param changePassword body rest_change_password.Request true "Change password Request"
// @Success 200 {object} rest_change_password.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/change-password [put]
func (config authenticatorRESTConfigurator) changePassword(c *gin.Context) {
	//rest.NotImplemented(c)
	c.JSON(http.StatusOK, rest_change_password.Response{})

}

// @Summary Reset password functionality
// @Description It resets your password given the correct code and new password.
// @Accept  json
// @Produce  json
// @Param resetPassword body rest_reset_password.Request true "Reset password Request"
// @Success 200 {object} rest_reset_password.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/reset-password [put]
func (config authenticatorRESTConfigurator) resetPassword(c *gin.Context) {
	//rest.NotImplemented(c)
	c.JSON(http.StatusOK, rest_reset_password.Response{})

}

// @Summary Reset password request functionality
// @Description It enters into the flow of reset password sending an email with instructions
// @Accept  json
// @Produce  json
// @Param resetPassword body rest_reset_password_request.Request true "Reset password Request"
// @Success 202 {object} rest_reset_password_request.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/reset-password-request [put]
func (config authenticatorRESTConfigurator) resetPasswordRequest(c *gin.Context) {
	userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
	auth := implementation.NewAuthenticator(userRepo, sender)
	passwordManager := implementation.NewPasswordManager(userRepo, in_memory.DummySender{})
	user := structures.User{Email: "email@bixlabs.com", Password: "password1"}
	_, _ = auth.Signup(user)
	var request rest_reset_password_request.Request
	if isInvalidResetPasswordRequest(c, &request) {
		c.JSON(http.StatusBadRequest, rest_reset_password_request.NewErrorResponse(http.StatusBadRequest,
			errors.New("email is required")))
	} else {
		c.JSON(resetPasswordRequestHandler(request.Email, passwordManager))
	}
}

func isInvalidResetPasswordRequest(c *gin.Context, request *rest_reset_password_request.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func resetPasswordRequestHandler(email string, handler interactors.PasswordManager) (int, rest_reset_password_request.Response) {
	_, err := handler.SendResetPasswordRequest(email)
	if err != nil {
		if _, ok := err.(util.InvalidEmailError); ok {
			return http.StatusBadRequest, rest_reset_password_request.NewErrorResponse(http.StatusBadRequest, err)
		}
		return http.StatusInternalServerError, rest_reset_password_request.NewErrorResponse(http.StatusInternalServerError, err)
	}

	return http.StatusAccepted, rest_reset_password_request.NewResponse(http.StatusAccepted, &rest_reset_password_request.Result{Success: true})
}
