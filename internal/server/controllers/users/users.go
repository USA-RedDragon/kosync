package users

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/server/apimodels"
	"github.com/USA-RedDragon/kosync/internal/store"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	config, ok := c.MustGet("config").(*config.Config)
	if !ok {
		slog.Error("failed to get config from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}
	if !config.Auth.AllowRegistration {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "registration is disabled"})
		return
	}

	var user apimodels.UserCreateRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, ok := c.MustGet("store").(store.Store)
	if !ok {
		slog.Error("failed to get store from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}

	// Check if the user already exists
	_, err := db.GetUserByUsername(user.Username)
	if err != nil {
		if !errors.Is(err, store.ErrUserNotFound) {
			slog.Error("failed to get user from store", "error", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
			return
		} else {
			// User does not exist, proceed to create
			err = db.CreateUser(user.Username, user.Password)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "user created"})
			return
		}
	}

	// User already exists
	c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": "user already exists"})
}

func Auth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "authenticated"})
}
