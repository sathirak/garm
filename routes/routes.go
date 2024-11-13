package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/controllers"
	"github.com/hotelbear/garm/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		auth := v1.Group("/auth")

		// Public routes
		auth.GET("/healthz", controllers.Healthz)

		user := auth.Group("/user")
		user.POST("/password-check", controllers.CheckPasswordUser)
		user.POST("/sign-up", controllers.SignUpUser)
		user.POST("/sign-in", controllers.SignInUser)

		user.POST("/reset", controllers.ResetPasswordUser)

		// Private routes
		service := auth.Group("/service")
		service.Use(middlewares.ApiKeyAuth())
		service.GET("/", controllers.AuthenticateUser)
	}
}
