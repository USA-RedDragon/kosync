package server

import (
	"net/http"

	"github.com/USA-RedDragon/kosync/internal/server/controllers/syncs"
	"github.com/USA-RedDragon/kosync/internal/server/controllers/users"
	"github.com/gin-gonic/gin"
)

func applyRoutes(r *gin.Engine) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"state": "OK"})
	})

	r.POST("/users/create", users.Create)
	r.GET("/users/auth", users.Auth)
	r.PUT("/syncs/progress", requireLogin(), syncs.UpdateProgress)
	r.GET("/syncs/progress/:document", requireLogin(), syncs.GetProgress)
}

func requireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}
