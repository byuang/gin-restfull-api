package main

import (
	"gin-restfull-api/config"
	"gin-restfull-api/controller"
	"gin-restfull-api/helper"
	"gin-restfull-api/model"
	"gin-restfull-api/repository"
	"gin-restfull-api/router"
	"gin-restfull-api/service"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})
	db.Table("tags").AutoMigrate(&model.Tags{})

	//Init Repository
	userRepository := repository.NewUsersRepositoryImpl(db)
	tagsRepository := repository.NewTagsREpositoryImpl(db)

	//Init Service
	authenticationService := service.NewAuthenticationServiceImpl(userRepository, validate)
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	//Init controller
	authenticationController := controller.NewAuthenticationController(authenticationService)
	usersController := controller.NewUsersController(userRepository)
	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(userRepository, authenticationController, usersController, tagsController)

	

	server := &http.Server{
		Addr:           ":" + loadConfig.ServerPort,
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
