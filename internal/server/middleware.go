package server

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/store"
	"github.com/USA-RedDragon/kosync/internal/utils"
	"github.com/gin-gonic/gin"
)

func applyMiddleware(r *gin.Engine, config *config.Config) {
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.TrustedPlatform = "X-Real-IP"

	// TODO: Set trusted proxies
	// err := r.SetTrustedProxies(config.TrustedProxies)
	// if err != nil {
	// 	slog.Error("Failed to set trusted proxies", "error", err.Error())
	// }
}

func requireLogin(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Request.Header.Get("X-Auth-User")
		password := c.Request.Header.Get("X-Auth-Key")
		if user == "" || password == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
			return
		}

		negativeResp := gin.H{"error": "Authentication failed"}

		db, ok := c.MustGet("store").(store.Store)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, negativeResp)
			return
		}

		// Check if the user exists and the password is correct
		dbUser, err := db.GetUserByUsername(user)
		if err != nil {
			if errors.Is(err, store.ErrUserNotFound) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, negativeResp)
				return
			}
			slog.Error("Failed to get user by username", "error", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, negativeResp)
			return
		}

		ok, err = utils.VerifyPassword(password, dbUser.Password, config.Auth.Salt)
		if err != nil {
			slog.Error("Failed to verify password", "error", err.Error())
			c.AbortWithStatusJSON(http.StatusInternalServerError, negativeResp)
			return
		}
		if ok {
			// Password matches
			c.Set("user", dbUser)
			c.Next()
			return
		}

		// Password does not match
		c.AbortWithStatusJSON(http.StatusUnauthorized, negativeResp)
	}
}
