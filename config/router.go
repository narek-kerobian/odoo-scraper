package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/controller"
)

// Defines application routes
func InitRoutes(r *gin.Engine) {
    // Heartbeat controller
    r.GET("/heartbeat", controller.GetHeartbeat)
}
