package main

import (
	"RestApi-Gin/controller"
	"RestApi-Gin/database"
	"RestApi-Gin/repository"
	"RestApi-Gin/router"
	"RestApi-Gin/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	db := database.ConnDb()
	validate := validator.New()

	//repository
	repositoryImpl := repository.NewMahasiswaRepositoryImpl(db)
	//service
	serviceImpl := service.NewMahasiswaServiceImpl(repositoryImpl, validate)
	//controller
	controllerImpl := controller.NewMahasiswaControllerImpl(serviceImpl)
	//router
	newRouter := router.NewRouter(controllerImpl)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: newRouter,
	}
	server.ListenAndServe()

}
