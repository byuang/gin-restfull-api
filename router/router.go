package router

import (
	"gin-restfull-api/controller"
	"gin-restfull-api/middleware"
	"gin-restfull-api/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UsersRepository, authenticationController *controller.AuthenticationController, usersController *controller.UserController, tagsController *controller.TagsController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/authentication")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(userRepository), usersController.GetUsers)

	tagsRouter := router.Group("/tags")
	tagsRouter.GET("", middleware.DeserializeUser(userRepository), tagsController.FindAll)
	tagsRouter.GET("/:tagId", middleware.DeserializeUser(userRepository), tagsController.FindById)
	tagsRouter.POST("", middleware.DeserializeUser(userRepository), tagsController.Create)
	tagsRouter.PATCH("/:tagId", middleware.DeserializeUser(userRepository), tagsController.Update)
	tagsRouter.DELETE("/:tagId", middleware.DeserializeUser(userRepository), tagsController.Delete)

	return service
}
