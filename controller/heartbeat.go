package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// An endpoint for checking if the service is alive
func GetHeartbeat(c *gin.Context) {
    c.JSON(http.StatusNoContent, c)
}

