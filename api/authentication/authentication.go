package authentication

import (
	"errors"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_change_password"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_login"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_reset_password"
	"github.com/bixlabs/authentication/api/authentication/structures/rest_signup"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authenticatorRESTConfigurator struct {
	handler interactors.Authenticator
}

func NewAuthenticatorRESTConfigurator(handler interactors.Authenticator, router *gin.Engine) {
	configureAuthRoutes(authenticatorRESTConfigurator{handler}, router)
}

func configureAuthRoutes(restConfig authenticatorRESTConfigurator, r *gin.Engine) *gin.Engine {
	router := r.Group("/user")
	router.POST("/login", restConfig.login)
	router.POST("/signup", restConfig.signup)
	router.PUT("/change-password", restConfig.changePassword)
	router.PUT("/reset-password", restConfig.resetPassword)
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
	user := structures.User{Email: "email@bixlabs.com", Password: "password1"}
	_, _ = config.handler.Signup(user)
	var request rest_login.Request
	if c.ShouldBindJSON(&request) != nil || request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, rest_login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or password missing")))
	} else {
		c.JSON(loginHandler(request.Email, request.Password, config.handler))
	}
}

func loginHandler(email, password string, handler interactors.Authenticator) (int, rest_login.Response) {
	response, err := handler.Login(email, password)
	if err != nil {
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

	return http.StatusOK, rest_login.NewResponse(http.StatusOK, response)
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
// @Description It resets your password and provide you with a flow to enter a new one.
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
