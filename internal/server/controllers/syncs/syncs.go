package syncs

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/USA-RedDragon/kosync/internal/server/apimodels"
	"github.com/USA-RedDragon/kosync/internal/store"
	storeErrs "github.com/USA-RedDragon/kosync/internal/store/errors"
	"github.com/USA-RedDragon/kosync/internal/store/models"
	"github.com/gin-gonic/gin"
)

func UpdateProgress(c *gin.Context) {
	var progress apimodels.SyncProgressRequest
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, ok := c.MustGet("store").(store.Store)
	if !ok {
		slog.Error("failed to get store from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}

	user, ok := c.MustGet("user").(models.User)
	if !ok {
		slog.Error("failed to get user from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}

	err := db.UpdateProgress(models.Progress{
		User:       user.Username,
		Document:   progress.Document,
		Percentage: progress.Percentage,
		Progress:   progress.Progress,
		Device:     progress.Device,
	})
	if err != nil {
		slog.Error("failed to update progress", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to update progress"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func GetProgress(c *gin.Context) {
	document := c.Param("document")
	if document == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "document is required"})
		return
	}
	db, ok := c.MustGet("store").(store.Store)
	if !ok {
		slog.Error("failed to get store from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}

	user, ok := c.MustGet("user").(models.User)
	if !ok {
		slog.Error("failed to get user from context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "try again later"})
		return
	}

	progress, err := db.GetProgress(user.Username, document)
	if err != nil {
		if errors.Is(err, storeErrs.ErrProgressNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "progress not found"})
			return
		}
		slog.Error("failed to get progress", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to get progress"})
		return
	}

	c.JSON(http.StatusOK, apimodels.ProgressResponse{
		Document:   progress.Document,
		Percentage: progress.Percentage,
		Progress:   progress.Progress,
		Device:     progress.Device,
	})
}
