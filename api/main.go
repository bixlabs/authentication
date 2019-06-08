package main

import (
	"fmt"
	"github.com/bixlabs/authentication/api/authentication"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/database/user/in_memory"
	"github.com/bixlabs/authentication/database/user/sqlite"
	_ "github.com/bixlabs/authentication/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Go-Authenticator
// @version 1.0
// @description Leverage of authentication functionality

// @contact.name API Support
// @contact.url https://bixlabs.com/
// @contact.email jarrieta@bixlabs.com
// @name Authorization

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @BasePath /

func main() {
	tools.InitializeLogger()
	router := NewGinRouter()
	userRepo, closeDB := sqlite.NewSqliteStorage()
	defer closeDB()
	authOperations := implementation.NewAuthenticator(userRepo, in_memory.DummySender{})
	authentication.NewAuthenticatorRESTConfigurator(authOperations, router)
	runGinRouter(router, getRestConfiguration().Port)
}

func NewGinRouter() *gin.Engine {
	result := gin.Default()
	result.Use(gin.Logger())
	result.Use(gin.Recovery())
	configureSwagger(result)
	return result
}

func configureSwagger(result *gin.Engine) gin.IRoutes {
	return result.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func runGinRouter(router *gin.Engine, port string) {
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}
}

func getRestConfiguration() RestConfiguration {
	result := RestConfiguration{}
	err := env.Parse(&result)
	if err != nil {
		tools.Log().Panic("parsing the env variables for the rest configuration failed", err)
	}
	return result
}

type RestConfiguration struct {
	Port string `env:"AUTH_SERVER_PORT" envDefault:"8080"`
}
