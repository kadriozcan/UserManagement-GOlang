package main

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/user_management_system/controller"
	"github.com/user_management_system/helper"
	"github.com/user_management_system/model"
	"github.com/user_management_system/repository"
	"github.com/user_management_system/router"
	"github.com/user_management_system/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// logged when the server started.
	log.Info().Msg("Server is running...")

	// Database connection
	dbSource := "./DB/user_management.db"
	db, errDB := gorm.Open(sqlite.Open(dbSource), &gorm.Config{})
	if errDB != nil {
		log.Fatal().Err(errDB).Msg("Failed to connect to the database")
	} else {
		log.Info().Msg("Connected to the database")
	}

	db.Table("users").AutoMigrate(&model.User{})

	// Repository
	userRepository := repository.NewUserRepositoryImpl(db)

	// Service
	userService := service.NewUserService(userRepository)

	// Controller
	userController := controller.NewUserController(userService)

	// Router
	routes := router.NewRouter(userController)

	// creates a HTTP server on port : 8008 and uses defined routes
	server := &http.Server{
		Addr:    ":8008",
		Handler: routes,
	}

	// starts the server and listens for coming requests. If there is an error it stores it
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
