package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/controllers"
	"github.com/sathirak/garm/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		auth := v1.Group("/auth")

		// Public routes
		auth.GET("/healthz", controllers.Healthz)

		partner := auth.Group("/partner")
		partner.POST("/password-check", controllers.CheckPasswordPartner)
		partner.POST("/sign-up", controllers.SignUpPartner)
		partner.POST("/sign-in", controllers.SignInPartner)

		partner.POST("/reset", controllers.ResetPasswordPartner)

		// Private routes
		service := auth.Group("/service")
		service.Use(middlewares.ApiKeyAuth())
		service.GET("/", controllers.AuthenticatePartner)
	}
}
