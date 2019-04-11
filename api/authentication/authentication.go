package authentication

import (
	"fmt"
	"github.com/bixlabs/authentication/todo/interactors"
	"github.com/bixlabs/authentication/todo/structures"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"time"
)

type authenticatorRESTConfigurator struct {
	handler interactors.TodoOperations
	Port    string `env:"WEB_SERVER_PORT" envDefault:"3000"`
}

func NewAuthenticatorRESTConfigurator(handler interactors.TodoOperations) {
	restConfig := getRestConfiguration(handler)
	router := configureGinRouter(restConfig)
	runGinRouter(router, restConfig.Port)
}

func getRestConfiguration(handler interactors.TodoOperations) authenticatorRESTConfigurator {
	restConfig := authenticatorRESTConfigurator{handler, ""}
	parseEnvVariables(&restConfig)
	return restConfig
}

func parseEnvVariables(restConfig *authenticatorRESTConfigurator) {
	err := env.Parse(restConfig)
	if err != nil {
		tools.Log().Panic("parsing the env variables for the rest configuration threw an error", err)
	}
}

func configureGinRouter(restConfig authenticatorRESTConfigurator) *gin.Engine {
	router := gin.Default()
	router.POST("/login", restConfig.login)
	router.GET("/signup", restConfig.signup)
	router.PUT("/change-password", restConfig.changePassword)
	router.DELETE("/todo/:id", restConfig.resetPassword)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func runGinRouter(router *gin.Engine, port string) {
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

// @Summary Create Todo
// @Description Creates a todo given the correct JSON representation of it.
// @Accept  json
// @Produce  json
// @Param todo body structures.Todo true "Todo structure"
// @Success 200 {object} structures.Todo
// @Header 200 {string} Token "qwerty"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todo [post]
func (config authenticatorRESTConfigurator) login(c *gin.Context) {
	var request Request
	var todo *structures.Todo

	if err := c.ShouldBind(&request); err == nil {
		tools.Log().WithFields(logrus.Fields{"Request": request}).Info("A request object was received")
		todo = config.handler.Create(RequestToTodo(request))
		c.String(http.StatusOK, fmt.Sprintf("Create was successful for TODO with name: %s", todo.Name))
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (config authenticatorRESTConfigurator) signup(c *gin.Context) {
	id := c.Param("id")
	todo := config.handler.Read(id)
	c.String(http.StatusOK, fmt.Sprintf("Read was successful for TODO with ID: %s", todo.ID))
}

type Request struct {
	ID          string    `json:"i_am"`
	Name        string    `json:"title"`
	Description string    `json:"the_rest"`
	DueDate     time.Time `json:"when_finish"`
}

func RequestToTodo(request Request) structures.Todo {
	return structures.Todo{ID: request.ID, Name: request.Name, Description: request.Description, DueDate: request.DueDate}
}

func (config authenticatorRESTConfigurator) changePassword(c *gin.Context) {
	var request Request
	var todo *structures.Todo

	if c.ShouldBind(&request) == nil {
		todo = config.handler.Update(RequestToTodo(request))
	} else {
		// handle validation case
		tools.Log().Info("Validation case")
	}

	c.String(http.StatusOK, fmt.Sprintf("Update was successful for TODO with name: %s", todo.Name))
}

func (config authenticatorRESTConfigurator) resetPassword(c *gin.Context) {
	id := c.Param("id")
	success := config.handler.Delete(id)
	c.String(http.StatusOK, fmt.Sprintf("Delete was successful %t", success))
}
