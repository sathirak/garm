package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/controllers"
	"github.com/sathirak/garm/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		// Public routes
		auth.GET("/healthz", controllers.Healthz)
		auth.POST("/email-password/password-check", controllers.PasswordCheck)

		auth.POST("/email-password/sign-up", controllers.SignUpEmailPassword)
		auth.POST("/email-password/sign-in", controllers.SignInEmailPassword)

		auth.POST("/email-password/reset/:userID", controllers.ResetEmailPassword)

		// Private routes
		service := auth.Group("/service")
		service.Use(middlewares.ApiKeyAuth())
		service.GET("/", controllers.Authenticate)
	}
}
