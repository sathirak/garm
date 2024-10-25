package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/controllers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")

		auth.GET("/healthz", controllers.Healthz)

		// Authenticate all users
		auth.GET("/", controllers.Authenticate)

		// Email Password recipe
		auth.POST("/email-password/sign-up", controllers.SignUpEmailPassword)
		auth.POST("/email-password/sign-in", controllers.SignInEmailPassword)

    // Reset Email password credentials
    auth.POST("/email-password/reset/:userID", controllers.ResetEmailPassword)
	}
}
