package authentication

import (
	"errors"
	"github.com/bixlabs/authentication/api/authentication/structures/change_password"
	"github.com/bixlabs/authentication/api/authentication/structures/forgot_password"
	"github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/api/authentication/structures/login/mappers"
	"github.com/bixlabs/authentication/api/authentication/structures/reset_password"
	"github.com/bixlabs/authentication/api/authentication/structures/signup"
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
	router.PUT("/reset-password-request", restConfig.forgotPassword)
	return r
}

// @Summary Login functionality
// @Description Attempts to authenticate the user with the given credentials.
// @Accept  json
// @Produce  json
// @Param login body login.Request true "Login Request"
// @Success 200 {object} login.SwaggerResponse
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
	var request login.Request
	if isInvalidLoginRequest(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or password missing")))
	} else {
		c.JSON(loginHandler(request.Email, request.Password, auth))
	}
}

func isInvalidLoginRequest(c *gin.Context, request *login.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func loginHandler(email, password string, handler interactors.Authenticator) (int, login.Response) {
	response, err := handler.Login(email, password)
	if err != nil {
		return handleLoginErrors(err)
	}

	return http.StatusOK, login.NewResponse(http.StatusOK, mappers.LoginResponseToResult(*response))
}

func handleLoginErrors(err error) (int, login.Response) {
	if _, ok := err.(util.InvalidEmailError); ok {
		return http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest, err)
	}
	if _, ok := err.(util.PasswordLengthError); ok {
		return http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest, err)
	}
	if _, ok := err.(util.WrongCredentialsError); ok {
		return http.StatusUnauthorized, login.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, login.NewErrorResponse(http.StatusBadRequest, err)
}

// @Summary Signup functionality
// @Description Attempts to create a user provided the correct information.
// @Accept  json
// @Produce  json
// @Param signup body signup.Request true "Signup Request"
// @Success 201 {object} signup.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/signup [post]
func (config authenticatorRESTConfigurator) signup(c *gin.Context) {
	//rest.NotImplemented(c)
	c.JSON(http.StatusCreated, signup.Response{})
}

// @Summary Change password functionality
// @Description It changes the password provided the old one and a new password.
// @Accept  json
// @Produce  json
// @Param changePassword body change_password.Request true "Change password Request"
// @Success 200 {object} change_password.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/change-password [put]
func (config authenticatorRESTConfigurator) changePassword(c *gin.Context) {
	//rest.NotImplemented(c)
	c.JSON(http.StatusOK, change_password.Response{})

}


// @Summary Reset password functionality
// @Description It resets your password given the correct code and new password.
// @Accept  json
// @Produce  json
// @Param resetPassword body reset_password.Request true "Reset password Request"
// @Success 200 {object} reset_password.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/reset-password [put]
func (config authenticatorRESTConfigurator) resetPassword(c *gin.Context) {
	var request reset_password.Request
	if isInvalidResetPassword(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or code missing")))
	} else {
		config.handleNoContentOrErrorResponse(request, c)
	}
}

func isInvalidResetPassword(c *gin.Context, request *reset_password.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Code == ""
}

func resetPasswordHandler(email string, code string, newPassword string, handler interactors.PasswordManager) (int, reset_password.Response ) {
	err := handler.ResetPassword(email, code, newPassword)

	if err != nil {
		switch err.(type) {
		case util.InvalidEmailError:
			return http.StatusBadRequest, reset_password.Response{}
		case util.PasswordLengthError:
			return http.StatusBadRequest, reset_password.Response{}
		case util.InvalidResetPasswordCode:
			return http.StatusBadRequest, reset_password.Response{}
		default:
			return http.StatusInternalServerError, reset_password.Response{}
		}
	}
	return http.StatusNoContent, reset_password.Response{}
}

func (config authenticatorRESTConfigurator) handleNoContentOrErrorResponse(request reset_password.Request, c *gin.Context) {
	if code, response := resetPasswordHandler(request.Email, request.Code, request.NewPassword, config.passwordManager); code == http.StatusNoContent {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(code, response)
	}
}

// @Summary Forgot password request functionality
// @Description It enters into the flow of reset password sending an email with instructions
// @Accept  json
// @Produce  json
// @Param resetPassword body forgot_password.Request true "Forgot password request"
// @Success 202 {object} forgot_password.SwaggerResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/reset-password-request [put]
func (config authenticatorRESTConfigurator) forgotPassword(c *gin.Context) {
	userRepo, sender := in_memory.NewUserRepo(), in_memory.DummySender{}
	auth := implementation.NewAuthenticator(userRepo, sender)
	passwordManager := implementation.NewPasswordManager(userRepo, in_memory.DummySender{})
	user := structures.User{Email: "email@bixlabs.com", Password: "password1"}
	_, _ = auth.Signup(user)
	var request forgot_password.Request
	if isInvalidForgotPassword(c, &request) {
		c.JSON(http.StatusBadRequest, forgot_password.NewErrorResponse(http.StatusBadRequest,
			errors.New("email is required")))
	} else {
		c.JSON(forgotPasswordHandler(request.Email, passwordManager))
	}
}

func isInvalidForgotPassword(c *gin.Context, request *forgot_password.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func forgotPasswordHandler(email string, handler interactors.PasswordManager) (int, forgot_password.Response) {
	_, err := handler.ForgotPassword(email)
	if err != nil {
		if _, ok := err.(util.InvalidEmailError); ok {
			return http.StatusBadRequest, forgot_password.NewErrorResponse(http.StatusBadRequest, err)
		}
		return http.StatusInternalServerError, forgot_password.NewErrorResponse(http.StatusInternalServerError, err)
	}

	return http.StatusAccepted, forgot_password.NewResponse(http.StatusAccepted, &forgot_password.Result{Success: true})
}
