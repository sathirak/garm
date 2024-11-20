package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/controllers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	{
		v1.GET("/healthz", controllers.Healthz)
		v1.GET("/verify", controllers.AuthenticateUser)
	}

	user := v1.Group("/user")
	{
		user.POST("/password-check", controllers.CheckPasswordUser)
		user.POST("/sign-up", controllers.SignUpUser)
		user.POST("/sign-in", controllers.SignInUser)
		user.POST("/reset", controllers.ResetPasswordUser)
	}
}
