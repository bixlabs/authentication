package user_management

import (
	"github.com/bixlabs/authentication/api/authentication/structures/login"
	"github.com/bixlabs/authentication/api/user_management/structures/create"
	"github.com/bixlabs/authentication/api/user_management/structures/mappers"
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/authenticator/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

const tokenHeaderLength = 2

type userManagerRESTConfigurator struct {
	userManager interactors.UserManager
}

func NewUserManagerRESTConfigurator(userManager interactors.UserManager, r *gin.Engine) {
	configureUserManagerRoutes(userManagerRESTConfigurator{userManager}, r)
}

func configureUserManagerRoutes(restConfig userManagerRESTConfigurator, r *gin.Engine) *gin.Engine {
	router := r.Group("/v1/user-management")
	router.POST("/", restConfig.create)
	return r
}

// @Tags User
// @Summary Login functionality
// @Description Attempts to authenticate the user with the given credentials.
// @Accept  json
// @Produce  json
// @Param login body login.Request true "Login Request"
// @Success 200 {object} login.RestResponse
// @Failure 400 {object} rest.ResponseWrapper
// @Failure 401 {object} rest.ResponseWrapper
// @Failure 500 {object} rest.ResponseWrapper
// @Router /user/login [post]
func (config userManagerRESTConfigurator) create(c *gin.Context) {
	var request create.Request
	if c.ShouldBind(&request) != nil {

	}
	c.JSON(createHandler(request.User, config.userManager))

	//var request login.Request
	//if isInvalidLoginRequest(c, &request) {
	//	c.JSON(http.StatusBadRequest, login.NewErrorResponse(http.StatusBadRequest,
	//		errors.New("email or password missing")))
	//} else {
	//	c.JSON(loginHandler(request.Email, request.Password, config.authenticator))
	//}
}

func isInvalidLoginRequest(c *gin.Context, request *login.Request) bool {
	return c.ShouldBindJSON(request) != nil || request.Email == "" || request.Password == ""
}

func createHandler(user structures.User, userManager interactors.UserManager) (int, create.RestResponse) {
	response, err := userManager.Create(user)
	if err != nil {
		return handleLoginError(err)
	}
	return http.StatusOK, create.NewResponse(http.StatusOK, mappers.CreateResponseToResult(response))
}

func handleLoginError(err error) (int, create.RestResponse) {
	var code int
	switch err.(type) {
	case util.InvalidEmailError:
		code = http.StatusBadRequest
	}
	return code, create.NewErrorResponse(code, err)
}
