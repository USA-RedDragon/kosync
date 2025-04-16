package syncs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "OK"})
}

func GetProgress(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "OK"})
}
