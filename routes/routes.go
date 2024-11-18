package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/controllers"
)

func SetupRoutes(r *gin.Engine) {
	authV1 := r.Group("/v1")

	{
		// declare: public routes
		authV1.GET("/healthz", controllers.Healthz)
		authV1.GET("/verify", controllers.AuthenticateUser)

		authV1.POST("/user/password-check", controllers.CheckPasswordUser)
		authV1.POST("/user/sign-up", controllers.SignUpUser)
		authV1.POST("/user/sign-in", controllers.SignInUser)
		authV1.POST("/user/reset", controllers.ResetPasswordUser)
	}
}
