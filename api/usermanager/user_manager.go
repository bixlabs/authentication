package usermanager

import (
	"github.com/bixlabs/authentication/authenticator/interactors"
	"github.com/gin-gonic/gin"
)
type userManagerRESTConfigurator struct {
	userManager interactors.UserManager
}

func NewUserManagerRESTConfigurator(userManager interactors.UserManager, engine *gin.Engine){
	configureUserManagerRoutes(userManagerRESTConfigurator{userManager}, engine)
}

func configureUserManagerRoutes(restConfig userManagerRESTConfigurator, r *gin.Engine) *gin.Engine{
	router := r.Group("/v1/user")

	router.GET("/", restConfig.find)

	return r
}

func (config userManagerRESTConfigurator) findOne(c *gin.Context) {
}