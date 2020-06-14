package healthcheck

import (
	"errors"
	"net/http"

	"github.com/bixlabs/authentication/api/healthcheck/structures/check"
	"github.com/bixlabs/authentication/authenticator/database/user"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/provider/email"
	"github.com/gin-gonic/gin"
)

type checker struct {
	passwordManager interactors.PasswordManager
	authenticator   interactors.Authenticator
	sender          email.Sender
	repository      user.Repository
}

func NewHealthCheckRESTConfigurator(repo user.Repository, emailSender email.Sender, auth interactors.Authenticator,
	passManager interactors.PasswordManager, r *gin.Engine) {
	services := checker{
		repository:      repo,
		sender:          emailSender,
		authenticator:   auth,
		passwordManager: passManager,
	}
	configureHealthCheckRoute(services, r)
}

// @Tags Healthcheck
// @Summary Healthcheck functionality
// @Description Verifies connection to all services
// @Accept  json
// @Produce  json
// @Success 200 {object} check.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 401 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /healthcheck [get]
func configureHealthCheckRoute(healthCheck checker, r *gin.Engine) *gin.Engine {
	group := r.Group("/v1/healthcheck")
	group.GET("/", healthCheck.healthCheck)
	return r
}

func (hc checker) healthCheck(c *gin.Context) {
	c.JSON(healthCheckHandler(hc.repository, hc.sender, hc.authenticator, hc.passwordManager))
}

func healthCheckHandler(repo user.Repository, emailSender email.Sender, auth interactors.Authenticator,
	passManager interactors.PasswordManager) (int, check.Response) {
	if auth == nil || passManager == nil ||
		repo == nil || emailSender == nil {
		return http.StatusInternalServerError, check.NewErrorResponse(http.StatusInternalServerError,
			errors.New("one or more services are down"))
	}

	return http.StatusOK, check.NewResponse(http.StatusOK)
}
