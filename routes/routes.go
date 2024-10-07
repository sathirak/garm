package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/controllers"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/healthz", controllers.Healthz)

		user := v1.Group("/user")
		user.POST("/", controllers.CreateUser)
	}
}
