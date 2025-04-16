package server

import (
	"github.com/USA-RedDragon/kosync/internal/config"
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
