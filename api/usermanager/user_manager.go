package usermanager

import (
	"errors"
	"net/http"

	"github.com/bixlabs/authentication/api/usermanager/structures/create"
	"github.com/bixlabs/authentication/api/usermanager/structures/delete"
	findOne "github.com/bixlabs/authentication/api/usermanager/structures/findone"
	"github.com/bixlabs/authentication/api/usermanager/structures/mappers"
	"github.com/bixlabs/authentication/api/usermanager/structures/update"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/structures"
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

	router.POST("/search", restConfig.findOne)
	router.DELETE("/", restConfig.delete)
	router.PUT("/", restConfig.update)
	router.POST("/", restConfig.create)

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
// @Router /users/search [post]

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

// @Tags User
// @Summary Create one User functionality
// @Description Retreive user created.
// @Accept  json
// @Produce  json
// @Param create body create.Request true "Create User Request"
// @Success 201 {object} create.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [post]

func (config userManagerRESTConfigurator) create(c *gin.Context) {
	var request create.Request

	if isInvalidCreateRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New("email missing")))
	} else {
		c.JSON(createHandler(mappers.CreateRequestToUser(request), config.userManager))
	}
}

func isInvalidCreateRequest(c *gin.Context, request *create.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func createHandler(user structures.User, handler interactors.UserManager) (int, create.Response) {
	user, err := handler.Create(user)

	if err != nil {
		return http.StatusNotFound, create.NewErrorResponse(http.StatusNotFound, err)
	}

	return http.StatusOK, create.NewResponse(http.StatusOK, user)
}

// @Tags User
// @Summary Delete one User functionality
// @Description Delete one user by email.
// @Accept  json
// @Produce  json
// @Param delete body delete.Request true "Delete User Request"
// @Success 201 {object} delete.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [delete]
func (config userManagerRESTConfigurator) delete(c *gin.Context) {
	var request delete.Request

	if isInvalidDeleteRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New("email missing")))
	} else {
		c.JSON(deleteHandler(request.Email, config.userManager))
	}
}

func isInvalidDeleteRequest(c *gin.Context, request *delete.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func deleteHandler(email string, handler interactors.UserManager) (int, delete.Response) {
	err := handler.Delete(email)

	if err != nil {
		return http.StatusNotFound, delete.NewErrorResponse(http.StatusNotFound, err)
	}

	return http.StatusOK, delete.NewResponse(http.StatusOK)
}

// @Tags User
// @Summary Update User functionality
// @Description Update one user.
// @Accept  json
// @Produce  json
// @Param update body update.Request true "Update User Request"
// @Success 201 {object} update.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [put]
func (config userManagerRESTConfigurator) update(c *gin.Context) {
	var request update.Request

	if isInvalidUpdateRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New("email missing")))
	} else {
		c.JSON(updateHandler(request.Email, mappers.UpdateRequestToUpdateUser(request), config.userManager))
	}
}

func isInvalidUpdateRequest(c *gin.Context, request *update.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == ""
}

func updateHandler(email string, updateAttrs structures.UpdateUser, handler interactors.UserManager) (int, update.Response) {
	user, err := handler.Update(email, updateAttrs)

	if err != nil {
		return http.StatusNotFound, update.NewErrorResponse(http.StatusNotFound, err)
	}

	return http.StatusOK, update.NewResponse(http.StatusOK, user)
}
