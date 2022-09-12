package config

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/bracketnco/odoo-scraper/controller"
	"gitlab.com/bracketnco/odoo-scraper/middleware"
	"gitlab.com/bracketnco/odoo-scraper/service"
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

