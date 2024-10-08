package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/controllers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")

		auth.GET("/", controllers.Validate)
		auth.GET("/healthz", controllers.Healthz)
		auth.POST("/email-password/sign-up", controllers.SignUpEmailPassword)
		auth.POST("/email-password/sign-in", controllers.SignInEmailPassword)
	}
}
