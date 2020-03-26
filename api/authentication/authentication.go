package authentication

import (
	"errors"
	"github.com/bixlabs/authentication/api/authentication/structures/changepass"
	"github.com/bixlabs/authentication/api/authentication/structures/forgotpass"
	"github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/api/authentication/structures/mappers"
	"github.com/bixlabs/authentication/api/authentication/structures/resetpass"
	"github.com/bixlabs/authentication/api/authentication/structures/signup"
	"github.com/bixlabs/authentication/api/authentication/structures/token"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const tokenHeaderLength = 2

type authenticatorRESTConfigurator struct {
	authenticator   interactors.Authenticator
	passwordManager interactors.PasswordManager
}

func NewAuthenticatorRESTConfigurator(auth interactors.Authenticator, pm interactors.PasswordManager, r *gin.Engine) {
	configureAuthRoutes(authenticatorRESTConfigurator{auth, pm}, r)
}

func configureAuthRoutes(restConfig authenticatorRESTConfigurator, r *gin.Engine) *gin.Engine {
	router := r.Group("/user")
	router.POST("/login", restConfig.login)
	router.POST("/signup", restConfig.signup)
	router.PUT("/change-password", restConfig.changePassword)
	router.PUT("/reset-password", restConfig.resetPassword)
	router.PUT("/reset-password-request", restConfig.forgotPassword)
	router.GET("token/validate", restConfig.verifyJWT)
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
// @Failure 500 {object} rest.ResponseWrapper
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
		return handleLoginError(err)
	}
	return http.StatusOK, login.NewResponse(http.StatusOK, mappers.LoginResponseToResult(*response))
}

func handleLoginError(err error) (int, login.Response) {
	var code int
	switch err.(type) {
	case util.InvalidEmailError:
		code = http.StatusBadRequest
	case util.PasswordLengthError:
		code = http.StatusBadRequest
	case util.WrongCredentialsError:
		code = http.StatusUnauthorized
	default:
		code = http.StatusInternalServerError
	}
	return code, login.NewErrorResponse(code, err)
}

// @Summary Signup functionality
// @Description Attempts to create a user provided the correct information.
// @Accept  json
// @Produce  json
// @Param signup body signup.Request true "Signup Request"
// @Success 201 {object} signup.SwaggerResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /user/signup [post]
func (config authenticatorRESTConfigurator) signup(c *gin.Context) {
	var request signup.Request
	if isInvalidSignupRequest(c, &request) {
		c.JSON(http.StatusBadRequest, signup.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or password missing")))
	} else {
		c.JSON(signupHandler(request, config.authenticator))
	}
}

func isInvalidSignupRequest(c *gin.Context, request *signup.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func signupHandler(request signup.Request, handler interactors.Authenticator) (int, signup.Response) {
	_, err := handler.Signup(mappers.SignupRequestToUser(request))
	if err != nil {
		return handleSignUpError(err)
	}
	return http.StatusCreated, signup.NewResponse(http.StatusCreated, &signup.Result{Success: true})
}

func handleSignUpError(err error) (int, signup.Response) {
	if isInvalidEmail(err) || isPasswordLength(err) || isDuplicatedEmail(err) {
		return http.StatusBadRequest, signup.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, signup.NewErrorResponse(http.StatusInternalServerError, err)
}

func isInvalidEmail(err error) bool {
	_, ok := err.(util.InvalidEmailError)
	return ok
}

func isPasswordLength(err error) bool {
	_, ok := err.(util.PasswordLengthError)
	return ok
}

func isDuplicatedEmail(err error) bool {
	_, ok := err.(util.DuplicatedEmailError)
	return ok
}

// @Summary Change password functionality
// @Description It changes the password provided the old one and a new password.
// @Accept  json
// @Produce  json
// @Param changePassword body changepass.Request true "Change password Request"
// @Success 200 {object} changepass.SwaggerResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 401 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /user/change-password [put]
func (config authenticatorRESTConfigurator) changePassword(c *gin.Context) {
	var request changepass.Request
	if isInvalidChangePasswordRequest(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email, old password or new password missing")))
	} else {
		c.JSON(changePasswordHandler(request, config.passwordManager))
	}
}

func isInvalidChangePasswordRequest(c *gin.Context, request *changepass.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.OldPassword == "" ||
		request.NewPassword == ""
}

func changePasswordHandler(request changepass.Request, pm interactors.PasswordManager) (int, changepass.Response) {
	if err := pm.ChangePassword(mappers.ChangePasswordRequestToUser(request), request.NewPassword); err != nil {
		return handleChangePasswordError(err)
	}
	return http.StatusOK, changepass.NewResponse(http.StatusOK, true)
}

func handleChangePasswordError(err error) (int, changepass.Response) {
	if isInvalidEmail(err) || isPasswordLength(err) || isSamePasswordChange(err) {
		return http.StatusBadRequest, changepass.NewErrorResponse(http.StatusBadRequest, err)
	}
	if _, ok := err.(util.WrongCredentialsError); ok {
		return http.StatusUnauthorized, changepass.NewErrorResponse(http.StatusUnauthorized, err)
	}
	return http.StatusInternalServerError, changepass.NewErrorResponse(http.StatusInternalServerError, err)
}

func isSamePasswordChange(err error) bool {
	_, ok := err.(util.SamePasswordChangeError)
	return ok
}

// @Summary Reset password functionality
// @Description It resets your password given the correct code and new password.
// @Accept  json
// @Produce  json
// @Param resetPassword body resetpass.Request true "Reset password Request"
// @Success 200 {object} resetpass.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 404 {object} rest.ResponseWrapper
// @Failure 408 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Failure 504 {object} rest.ResponseWrapper
// @Router /user/reset-password [put]
func (config authenticatorRESTConfigurator) resetPassword(c *gin.Context) {
	var request passreset.Request
	if isInvalidResetPassword(c, &request) {
		c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
			errors.New("email or code missing")))
	} else {
		config.handleNoContentOrErrorResponse(request, c)
	}
}

func isInvalidResetPassword(c *gin.Context, request *passreset.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Code == ""
}

func (config authenticatorRESTConfigurator) handleNoContentOrErrorResponse(request passreset.Request, c *gin.Context) {
	if code, response := resetPasswordHandler(request.Email, request.Code, request.NewPassword, config.passwordManager); code == http.StatusNoContent { //nolint
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(code, response)
	}
}

func resetPasswordHandler(email string, code string, newPassword string, handler interactors.PasswordManager) (int, passreset.Response) { //nolint
	if err := handler.ResetPassword(email, code, newPassword); err != nil {
		return handleResetPasswordError(err)
	}
	return http.StatusNoContent, passreset.Response{}
}

func handleResetPasswordError(err error) (int, passreset.Response) {
	if isInvalidEmail(err) || isPasswordLength(err) || isInvalidCode(err) || isSamePasswordChange(err) {
		return http.StatusBadRequest, passreset.Response{}
	}
	return http.StatusInternalServerError, passreset.Response{}
}

func isInvalidCode(err error) bool {
	_, ok := err.(util.InvalidResetPasswordCode)
	return ok
}

// @Summary Forgot password request functionality
// @Description It enters into the flow of reset password sending an email with instructions
// @Accept  json
// @Produce  json
// @Param resetPassword body forgotpass.Request true "Forgot password request"
// @Success 202 {object} forgotpass.SwaggerResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /user/reset-password-request [put]
func (config authenticatorRESTConfigurator) forgotPassword(c *gin.Context) {
	var request forgotpass.Request
	if isInvalidForgotPasswordRequest(c, &request) {
		c.JSON(http.StatusBadRequest, forgotpass.NewErrorResponse(http.StatusBadRequest,
			errors.New("email is required")))
	} else {
		c.JSON(forgotPasswordHandler(request.Email, config.passwordManager))
	}
}

func isInvalidForgotPasswordRequest(c *gin.Context, request *forgotpass.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func forgotPasswordHandler(email string, handler interactors.PasswordManager) (int, forgotpass.Response) {
	_, err := handler.ForgotPassword(email)
	if err != nil {
		return handleForgotPasswordError(err)
	}
	return http.StatusAccepted, forgotpass.NewResponse(http.StatusAccepted, &forgotpass.Result{Success: true})
}

func handleForgotPasswordError(err error) (int, forgotpass.Response) {
	if isInvalidEmail(err) {
		return http.StatusBadRequest, forgotpass.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, forgotpass.NewErrorResponse(http.StatusInternalServerError, err)
}

// @Summary Validates a JWT and returns the claims for it.
// @Description If the JWT is valid this endpoint returns the user inside of the token.
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Authorization: Bearer <jwtToken>"
// @Success 200 {object} token.SwaggerResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 401 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /user/token/validate [get]
func (config authenticatorRESTConfigurator) verifyJWT(c *gin.Context) {
	t, err := getTokenFromHeader(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, token.NewErrorResponse(http.StatusBadRequest, err))
	} else {
		c.JSON(verifyJWTHandler(t, config.authenticator))
	}
}

func getTokenFromHeader(c *gin.Context) (string, error) {
	// TODO: Use ShouldBindHeader when gin framework releases the feature, it's in master but not release.
	t := c.Request.Header.Get("Authorization")

	if t == "" || !strings.Contains(t, "Bearer") {
		return "", errors.New("token missing or malformed")
	}
	headerSeparated := strings.Split(t, " ")
	if len(headerSeparated) != tokenHeaderLength {
		return "", errors.New("token missing or malformed")
	}
	return headerSeparated[1], nil
}

func verifyJWTHandler(t string, handler interactors.Authenticator) (int, token.Response) {
	user, err := handler.VerifyJWT(t)
	if err != nil {
		if isInvalidToken(err) {
			return http.StatusUnauthorized, token.NewErrorResponse(http.StatusUnauthorized, err)
		}
		return http.StatusInternalServerError, token.NewErrorResponse(http.StatusInternalServerError, err)
	}
	return http.StatusOK, token.NewResponse(http.StatusOK, &token.Result{User: user})
}

func isInvalidToken(err error) bool {
	_, ok := err.(util.InvalidJWTToken)
	return ok
}
