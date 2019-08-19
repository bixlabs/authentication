package main

import (
	"fmt"
	"github.com/bixlabs/authentication/api/authentication"
	_ "github.com/bixlabs/authentication/api/docs" // docs is generated by Swag CLI, you have to import it.
	"github.com/bixlabs/authentication/authenticator/interactors/implementation"
	"github.com/bixlabs/authentication/authenticator/interactors/implementation/util"
	"github.com/bixlabs/authentication/database/user/memory"
	"github.com/bixlabs/authentication/database/user/sqlite"
	"github.com/bixlabs/authentication/tools"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"os"
	"time"
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
	authOperations := implementation.NewAuthenticator(userRepo, memory.DummySender{})
	passwordManager := implementation.NewPasswordManager(userRepo, memory.DummySender{})
	authentication.NewAuthenticatorRESTConfigurator(authOperations, passwordManager, router)
	runGinRouter(router, getRestConfiguration().Port)
}

func NewGinRouter() *gin.Engine {
	result := gin.New()

	result.Use(addRequestId())
	result.Use(logger())
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
		tools.Log().WithError(err).Panic("running the router for the rest configuration")

		panic(err)
	}
}

func getRestConfiguration() RestConfiguration {
	result := RestConfiguration{}
	err := env.Parse(&result)
	if err != nil {
		tools.Log().WithError(err).Panic("parsing the env variables for the rest configuration failed")
	}
	return result
}

func addRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := util.GenerateUniqueId()
		c.Set("rid", requestId)
		c.Header("X-Request-Id", requestId)
		c.Next()
	}
}

func logger() gin.HandlerFunc {
	if os.Getenv("AUTH_SERVER_APP_ENV") == "dev" {
		return gin.Logger()
	}

	logFile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logFile)
	return gin.LoggerWithFormatter(customFormatter)
}

func customFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - %s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.Keys["rid"],
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}

type RestConfiguration struct {
	Port string `env:"AUTH_SERVER_PORT" envDefault:"8080"`
}
