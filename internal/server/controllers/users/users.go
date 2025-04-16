package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "OK"})
}

func Auth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"state": "OK"})
}
