package controllers

import "github.com/gin-gonic/gin"

// Authorize godoc
//
//	@Summary		Policies for all services
//	@Description	Returns policies of garm
//	@Tags			policy
//	@Produce		json
//	@Router			/policies [get]
//
//	@Security		ApiKeyAuth
func Authorize(c *gin.Context) {
	c.File("./policies.json")
}
