package server

import (
	"net/http"

	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/server/controllers/syncs"
	"github.com/USA-RedDragon/kosync/internal/server/controllers/users"
	"github.com/gin-gonic/gin"
)

func applyRoutes(r *gin.Engine, config *config.Config) {
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"state": "OK"})
	})

	r.POST("/users/create", users.Create)
	r.GET("/users/auth", requireLogin(config), users.Auth)
	r.PUT("/syncs/progress", requireLogin(config), syncs.UpdateProgress)
	r.GET("/syncs/progress/:document", requireLogin(config), syncs.GetProgress)
}
