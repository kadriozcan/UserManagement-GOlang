package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/user_management_system/controller"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	// Enable CORS for all routes
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	usersRouter := router.Group("/users")
	usersRouter.GET("", userController.GetAll)
	usersRouter.GET("/:id", userController.GetById)
	usersRouter.POST("", userController.Create)
	usersRouter.PATCH("/:id", userController.Update)
	usersRouter.DELETE("/:id", userController.Delete)

	return router
}
