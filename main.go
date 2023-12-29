package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/user_management_system/helper"
)

func main() {

	// logged when the server started.
	log.Info().Msg("Server is running...")

	routes := gin.Default()
	// it will respond a message with an access to rooth path : "/"
	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome!")
	})

	// creates a HTTP server on port : 8008 and uses defined routes
	server := &http.Server{
		Addr:    ":8008",
		Handler: routes,
	}

	// starts the server and listens for coming requests. If there is an error it stores it
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
