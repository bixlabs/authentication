package usermanager

import (
	"errors"
	"net/http"

	findOne "github.com/bixlabs/authentication/api/usermanager/structures/findone"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/gin-gonic/gin"
)

type userManagerRESTConfigurator struct {
	userManager interactors.UserManager
}

func NewUserManagerRESTConfigurator(userManager interactors.UserManager, engine *gin.Engine) {
	configureUserManagerRoutes(userManagerRESTConfigurator{userManager}, engine)
}

func configureUserManagerRoutes(restConfig userManagerRESTConfigurator, r *gin.Engine) *gin.Engine {
	router := r.Group("/v1/users")

	router.POST("/", restConfig.findOne)

	return r
}

// @Tags User
// @Summary Find one User functionality
// @Description Retreive one user by email.
// @Accept  json
// @Produce  json
// @Param findone body findone.Request true "Find User Request"
// @Success 201 {object} findone.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [post]
func (config userManagerRESTConfigurator) findOne(c *gin.Context) {
	var request findOne.Request

	if isInvalidFindOneRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New("email missing")))
	} else {
		c.JSON(findOneHandler(request.Email, config.userManager))
	}
}

func isInvalidFindOneRequest(c *gin.Context, request *findOne.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func findOneHandler(email string, handler interactors.UserManager) (int, findOne.Response) {
	user, err := handler.Find(email)

	if err != nil {
		return http.StatusNotFound, findOne.NewErrorResponse(http.StatusNotFound, err)
	}

	return http.StatusOK, findOne.NewResponse(http.StatusOK, user)
}
