package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/user_management_system/controller"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

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
