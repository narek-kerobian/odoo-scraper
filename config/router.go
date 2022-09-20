package config

import (
	"github.com/gin-gonic/gin"
	"github.com/narek-kerobian/odoo-scraper/controller"
	"github.com/narek-kerobian/odoo-scraper/middleware"
	"github.com/narek-kerobian/odoo-scraper/service"
)

// Defines application routes
func InitRoutes(r *gin.Engine, dbPath string) {
    db := service.InitDb(dbPath)

    // Set static path
    r.Static("/static", "./static")

    // Heartbeat controller
    r.GET("/heartbeat", controller.GetHeartbeat)

    // Load middlewares
    r.Use(middleware.ExposeGinEngine(r))

    // Page routes
    r.GET("",  controller.ListPages(db)) 
    r.GET("/:id", controller.EditPage(db)) 
    r.POST("/:id", controller.EditPage(db)) 
}

