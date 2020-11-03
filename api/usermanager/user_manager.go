package usermanager

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bixlabs/authentication/api/usermanager/structures/create"
	"github.com/bixlabs/authentication/api/usermanager/structures/delete"
	findOne "github.com/bixlabs/authentication/api/usermanager/structures/findone"
	"github.com/bixlabs/authentication/api/usermanager/structures/mappers"
	"github.com/bixlabs/authentication/api/usermanager/structures/update"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
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

const emailMissingMessage = "email missing"

// @Tags User
// @Summary Find one User functionality
// @Description Retrieve one user by email.
// @Accept  json
// @Produce  json
// @Param findone body findone.Request true "Find User Request"
// @Success 200 {object} findone.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users/search [post]
func (config userManagerRESTConfigurator) findOne(c *gin.Context) {
	var request findOne.Request

	if isInvalidFindOneRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New(emailMissingMessage)))
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
// @Description Retrieve user created.
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
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New(emailMissingMessage)))
	} else {
		c.JSON(createHandler(request, config.userManager))
	}
}

func isInvalidCreateRequest(c *gin.Context, request *create.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func createHandler(request create.Request, handler interactors.UserManager) (int, create.Response) {
	user, err := handler.Create(mappers.CreateRequestToUser(request))

	if err != nil {
		return handleCreateError(err)
	}

	return http.StatusOK, create.NewResponse(http.StatusOK, user)
}

func handleCreateError(err error) (int, create.Response) {
	if isInvalidEmail(err) || isPasswordLength(err) || isDuplicatedEmail(err) {
		return http.StatusBadRequest, create.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, create.NewErrorResponse(http.StatusInternalServerError, err)
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

// @Tags User
// @Summary Update User functionality
// @Description Update one user.
// @Accept  json
// @Produce  json
// @Param update body update.Request true "Update User Request"
// @Success 200 {object} update.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [put]
func (config userManagerRESTConfigurator) update(c *gin.Context) {
	var request update.Request

	if isInvalidUpdateRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest,
			errors.New("json body failed on parsing")))
	} else {
		c.JSON(updateHandler(request.ID, request, config.userManager))
	}
}

func isInvalidUpdateRequest(c *gin.Context, request *update.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.ID == ""
}

func updateHandler(email string, request update.Request, handler interactors.UserManager) (int, update.Response) {
	fmt.Println(request)
	user, err := handler.Update(email, mappers.UpdateRequestToUpdateUser(request))

	if err != nil {
		return handleUpdateError(err)
	}

	return http.StatusOK, update.NewResponse(http.StatusOK, user)
}

func handleUpdateError(err error) (int, update.Response) {
	if isInvalidEmail(err) || isPasswordLength(err) || isDuplicatedEmail(err) {
		return http.StatusBadRequest, update.NewErrorResponse(http.StatusBadRequest, err)
	}
	return http.StatusInternalServerError, update.NewErrorResponse(http.StatusInternalServerError, err)
}

// @Tags User
// @Summary Delete one User functionality
// @Description Delete one user by email.
// @Accept  json
// @Produce  json
// @Param delete body delete.Request true "Delete User Request"
// @Success 200 {object} delete.Response
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /users [delete]
func (config userManagerRESTConfigurator) delete(c *gin.Context) {
	var request delete.Request

	if isInvalidDeleteRequest(c, &request) {
		c.JSON(http.StatusBadRequest, findOne.NewErrorResponse(http.StatusBadRequest, errors.New(emailMissingMessage)))
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
