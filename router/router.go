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
		c.JSON(404, gin.H{"code": 404, "message": "Page not found"})
	})

	router := service.Group("/api")
	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authenticationController.Register)
	authenticationRouter.POST("/login", authenticationController.Login)
	authenticationRouter.POST("/forgot-password", authenticationController.ForgotPassword)
	authenticationRouter.PATCH("/reset-password", authenticationController.ResetPassword)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.JwtMiddleware(userRepository), usersController.GetUsers)

	tagsRouter := router.Group("/tags")
	tagsRouter.GET("", middleware.JwtMiddleware(userRepository), tagsController.FindAll)
	tagsRouter.GET("/:tagId", middleware.JwtMiddleware(userRepository), tagsController.FindById)
	tagsRouter.POST("", middleware.JwtMiddleware(userRepository), tagsController.Create)
	tagsRouter.PATCH("/:tagId", middleware.JwtMiddleware(userRepository), tagsController.Update)
	tagsRouter.DELETE("/:tagId", middleware.JwtMiddleware(userRepository), tagsController.Delete)

	return service
}
