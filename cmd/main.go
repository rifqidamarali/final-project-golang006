package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rifqidamarali/final-project-golang006/internal/handler"
	"github.com/rifqidamarali/final-project-golang006/internal/infrastructure"
	"github.com/rifqidamarali/final-project-golang006/internal/middleware"
	"github.com/rifqidamarali/final-project-golang006/internal/repository"
	"github.com/rifqidamarali/final-project-golang006/internal/router"
	"github.com/rifqidamarali/final-project-golang006/internal/service"

	// _ "github.com/Calmantara/go-kominfo-2024/go-middleware/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	g := gin.Default()
	g.Use(gin.Recovery())

	// /public => generate JWT public
	usersGroup := g.Group("/users")
	gorm := infrastructure.NewGormPostgres()
	userRepository := repository.NewUserQuery(gorm)
	// userRepoMongo := repository.NewUserQueryMongo()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	auth := middleware.NewAuthorization(userService)
	userRouter := router.NewUserRouter(usersGroup, userHandler, auth)

	// mount
	userRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}



