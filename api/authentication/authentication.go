package authentication

import (
	"errors"
	"github.com/bixlabs/authentication/api/authentication/structures/change_password"
	"github.com/bixlabs/authentication/api/authentication/structures/forgot_password"
	"github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/api/authentication/structures/mappers"
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
	authenticator   interactors.Authenticator
	passwordManager interactors.PasswordManager
}

func NewAuthenticatorRESTConfigurator(auth interactors.Authenticator, pm interactors.PasswordManager, router *gin.Engine) {
	configureAuthRoutes(authenticatorRESTConfigurator{auth, pm}, router)
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
	var request login.Request
	if isInvalidLoginRequest(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or password missing")))
	} else {
		c.JSON(loginHandler(request.Email, request.Password, config.authenticator))
	}
}

func isInvalidLoginRequest(c *gin.Context, request *login.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func loginHandler(email, password string, handler interactors.Authenticator) (int, login.Response) {
	response, err := handler.Login(email, password)
	if err != nil {
		code, err := handleBasicErrors(err)
		return code, login.NewErrorResponse(code, err)
	}

	return http.StatusOK, login.NewResponse(http.StatusOK, mappers.LoginResponseToResult(*response))
}

func handleBasicErrors(err error) (int, error) {
	if _, ok := err.(util.InvalidEmailError); ok {
		return http.StatusBadRequest, err
	}
	if _, ok := err.(util.PasswordLengthError); ok {
		return http.StatusBadRequest, err
	}
	if _, ok := err.(util.WrongCredentialsError); ok {
		return http.StatusUnauthorized, err
	}
	return http.StatusInternalServerError, err
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
	var request change_password.Request
	if isInvalidChangePasswordRequest(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email, old password or new password missing")))
	} else {
		c.JSON(changePasswordHandler(request, config.passwordManager))
	}
}

func isInvalidChangePasswordRequest(c *gin.Context, request *change_password.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.OldPassword == "" || request.NewPassword == ""
}

func changePasswordHandler(request change_password.Request, passwordManager interactors.PasswordManager) (int, change_password.Response) {
	err := passwordManager.ChangePassword(mappers.ChangePasswordRequestToUser(request), request.NewPassword)
	if err != nil {
		code, err := handleBasicErrors(err)
		return code, change_password.NewErrorResponse(code, err)
	}

	return http.StatusOK, change_password.NewResponse(http.StatusOK, true)
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
	//rest.NotImplemented(c)
	c.JSON(http.StatusOK, reset_password.Response{})

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
	if isInvalidForgotPasswordRequest(c, &request) {
		c.JSON(http.StatusBadRequest, forgot_password.NewErrorResponse(http.StatusBadRequest,
			errors.New("email is required")))
	} else {
		c.JSON(forgotPasswordHandler(request.Email, passwordManager))
	}
}

func isInvalidForgotPasswordRequest(c *gin.Context, request *forgot_password.Request) bool {
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
